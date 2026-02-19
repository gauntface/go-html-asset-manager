/**
 * Copyright 2020 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 **/

package imgtopicture

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"os"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gauntface/go-html-asset-manager/v5/assets/genimgs"
	"github.com/gauntface/go-html-asset-manager/v5/manipulations"
	"github.com/gauntface/go-html-asset-manager/v5/utils/config"
	"github.com/gauntface/go-html-asset-manager/v5/utils/html/htmlparsing"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/html"
)

var errInjected = errors.New("injected error")

var reset func()

func TestMain(m *testing.M) {
	origGenimgsOpen := genimgsOpen
	origGenimgsLookupSizes := genimgsLookupSizes

	reset = func() {
		genimgsOpen = origGenimgsOpen
		genimgsLookupSizes = origGenimgsLookupSizes
	}

	os.Exit(m.Run())
}

func Test_shouldRun(t *testing.T) {
	tests := []struct {
		description string
		conf        *config.Config
		want        bool
	}{
		{
			description: "return false for nil conf",
			conf:        nil,
			want:        false,
		},
		{
			description: "return false for config without Assets",
			conf: &config.Config{
				Assets: nil,
			},
			want: false,
		},
		{
			description: "return false for config without static dir",
			conf: &config.Config{
				Assets: &config.AssetsConfig{
					StaticDir: "",
				},
			},
			want: false,
		},
		{
			description: "return false for config without generated dir",
			conf: &config.Config{
				Assets: &config.AssetsConfig{
					StaticDir: "/",
				},
			},
			want: false,
		},
		{
			description: "return false for config without ImgToPicture",
			conf: &config.Config{
				Assets: &config.AssetsConfig{
					StaticDir: "/",
				},
				ImgToPicture: nil,
			},
			want: false,
		},
		{
			description: "return true for config required params",
			conf: &config.Config{
				Assets: &config.AssetsConfig{
					StaticDir: "/",
				},
				ImgToPicture: []*config.ImgToPicConfig{
					{
						ID:          ".example",
						MaxWidth:    800,
						SourceSizes: []string{"(min-width: 800px) 800px", "100vw"},
					},
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := shouldRun(tt.conf)
			if got != tt.want {
				t.Errorf("Unexpected result; got %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_orderedSourceSets(t *testing.T) {
	tests := []struct {
		description     string
		sourceSetByType map[string][]genimgs.GenImg
		want            [][]genimgs.GenImg
	}{
		{
			description:     "return empty sets for no sources",
			sourceSetByType: map[string][]genimgs.GenImg{},
			want:            [][]genimgs.GenImg{},
		},
		{
			description: "return sorted sets",
			sourceSetByType: map[string][]genimgs.GenImg{
				"": {
					{
						Type: "",
						URL:  "/image.jpg",
					},
					{
						Type: "",
						URL:  "/image.png",
					},
				},
				"image/avif": {
					{
						Type: "image/avif",
						URL:  "/image.avif",
					},
				},
				"image/webp": {
					{
						Type: "image/webp",
						URL:  "/image.webp",
					},
				},
				"image/newformat": {
					{
						Type: "image/newformat",
						URL:  "/image.newformat",
					},
				},
			},
			want: [][]genimgs.GenImg{
				{
					{
						Type: "image/avif",
						URL:  "/image.avif",
					},
				},
				{
					{
						Type: "image/webp",
						URL:  "/image.webp",
					},
				},
				{
					{
						Type: "",
						URL:  "/image.jpg",
					},
					{
						Type: "",
						URL:  "/image.png",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := orderedSourceSets(tt.sourceSetByType)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Unexpected result; diff %v", diff)
			}
		})
	}
}

func Test_createSourceElement(t *testing.T) {
	tests := []struct {
		description string
		imgtopic    *config.ImgToPicConfig
		imgs        []genimgs.GenImg
		want        string
	}{
		{
			description: "return empty source element for no images",
			imgtopic:    &config.ImgToPicConfig{},
			imgs:        []genimgs.GenImg{},
			want:        `<source/>`,
		},
		{
			description: "return empty source element for no images",
			imgtopic: &config.ImgToPicConfig{
				SourceSizes: []string{
					"min-width(800px) 800px",
					"100vw",
				},
			},
			imgs: []genimgs.GenImg{
				{
					Type: "type/example",
					URL:  "/example-1.type",
					Size: 1,
				},
				{
					Type: "type/example",
					URL:  "/example-2.type",
					Size: 2,
				},
			},
			want: `<source type="type/example" sizes="min-width(800px) 800px,100vw" srcset="/example-1.type 1w,/example-2.type 2w"/>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := createSourceElement(tt.imgtopic, tt.imgs)
			if diff := cmp.Diff(MustRenderNode(t, got), tt.want); diff != "" {
				t.Fatalf("Unexpected return; diff %v", diff)
			}
		})
	}
}

func Test_pictureElement(t *testing.T) {
	tests := []struct {
		description string
		imgtopic    *config.ImgToPicConfig
		imgElement  *html.Node
		sizes       []genimgs.GenImg
		origWidth   int
		origHeight  int
		want        string
	}{
		{
			description: "return picture element for img without src",
			imgtopic: &config.ImgToPicConfig{
				SourceSizes: []string{
					"min-width(800px) 800px",
					"100vw",
				},
			},
			imgElement: &html.Node{
				Type: html.ElementNode,
				Data: "img",
				Attr: []html.Attribute{},
			},
			sizes:      []genimgs.GenImg{},
			origWidth:  1,
			origHeight: 1,
			want:       `<picture><img/></picture>`,
		},
		{
			description: "return picture element for img with src and gen images",
			imgtopic: &config.ImgToPicConfig{
				SourceSizes: []string{
					"min-width(800px) 800px",
					"100vw",
				},
			},
			imgElement: &html.Node{
				Type: html.ElementNode,
				Data: "img",
				Attr: []html.Attribute{
					{
						Key: "src",
						Val: "/example.png",
					},
				},
			},
			sizes: []genimgs.GenImg{
				{
					Type: "",
					Size: 1,
					URL:  "/example-1.png",
				},
				{
					Type: "",
					Size: 2,
					URL:  "/example-2.png",
				},
			},
			origWidth:  3,
			origHeight: 3,
			want:       `<picture><source sizes="min-width(800px) 800px,100vw" srcset="/example-1.png 1w,/example-2.png 2w"/><img src="/example-2.png"/></picture>`,
		},
		{
			description: "return picture element for img with gen images without a type",
			imgtopic: &config.ImgToPicConfig{
				SourceSizes: []string{
					"min-width(800px) 800px",
					"100vw",
				},
				Class: "picture-class",
			},
			imgElement: &html.Node{
				Type: html.ElementNode,
				Data: "img",
				Attr: []html.Attribute{
					{
						Key: "src",
						Val: "/example.png",
					},
					{
						Key: "class",
						Val: "example",
					},
				},
			},
			sizes: []genimgs.GenImg{
				{
					Type: "",
					Size: 1,
					URL:  "/example-1.png",
				},
				{
					Type: "",
					Size: 2,
					URL:  "/example-2.png",
				},
			},
			origWidth:  3,
			origHeight: 3,
			want:       `<picture class="picture-class"><source sizes="min-width(800px) 800px,100vw" srcset="/example-1.png 1w,/example-2.png 2w"/><img src="/example-2.png" class="example"/></picture>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := pictureElement(tt.imgtopic, tt.imgElement, tt.sizes, tt.origWidth, tt.origHeight)
			if diff := cmp.Diff(MustRenderNode(t, got), tt.want); diff != "" {
				t.Fatalf("Unexpected return; diff %v", diff)
			}
		})
	}
}

func Test_manipulateImg(t *testing.T) {
	tests := []struct {
		description        string
		debug              bool
		conf               *config.Config
		imgtopic           *config.ImgToPicConfig
		doc                *html.Node
		s3                 *s3.Client
		genimgsOpen        func(conf *config.Config, imgPath string) (image.Image, error)
		genimgsLookupSizes func(s3 genimgs.ListObjectsV2APIClient, conf *config.Config, imgPath string) ([]genimgs.GenImg, error)
		want               string
		wantError          error
	}{
		{
			description: "do nothing for img without src",
			doc:         MustGetNode(t, `<img/>`),
			want:        `<html><head></head><body><img/></body></html>`,
		},
		{
			description: "do nothing for img with a http src",
			doc:         MustGetNode(t, `<img src="http://example.com/example.png"/>`),
			want:        `<html><head></head><body><img src="http://example.com/example.png"/></body></html>`,
		},
		{
			description: "do nothing if the img cannot be found",
			doc:         MustGetNode(t, `<img src="/example.png"/>`),
			genimgsOpen: func(conf *config.Config, imgPath string) (image.Image, error) {
				wantImg := "/example.png"
				if wantImg != imgPath {
					t.Fatalf("Unexpected img path passed to genimgs.Open; got %v, want %v", imgPath, wantImg)
				}
				return nil, errInjected
			},
			want: `<html><head></head><body><img src="/example.png"/></body></html>`,
		},
		{
			description: "return error if sizes lookup fails",
			doc:         MustGetNode(t, `<img src="/example.png"/>`),
			genimgsOpen: func(conf *config.Config, imgPath string) (image.Image, error) {
				return &image.RGBA{}, nil
			},
			genimgsLookupSizes: func(s3 genimgs.ListObjectsV2APIClient, conf *config.Config, imgPath string) ([]genimgs.GenImg, error) {
				wantImg := "/example.png"
				if wantImg != imgPath {
					t.Fatalf("Unexpected img path passed to genimgs.LookupSizes; got %v, want %v", imgPath, wantImg)
				}
				return nil, errInjected
			},
			wantError: errInjected,
			want:      `<html><head></head><body><img src="/example.png"/></body></html>`,
		},
		{
			description: "do nothing if no sizes are found",
			doc:         MustGetNode(t, `<img src="/example.png"/>`),
			genimgsOpen: func(conf *config.Config, imgPath string) (image.Image, error) {
				return &image.RGBA{}, nil
			},
			genimgsLookupSizes: func(s3 genimgs.ListObjectsV2APIClient, conf *config.Config, imgPath string) ([]genimgs.GenImg, error) {
				return nil, nil
			},
			want: `<html><head></head><body><img src="/example.png"/></body></html>`,
		},
		{
			description: "replace img with picture",
			imgtopic: &config.ImgToPicConfig{
				SourceSizes: []string{"100vw"},
			},
			doc: MustGetNode(t, `<img src="/example.png"/>`),
			genimgsOpen: func(conf *config.Config, imgPath string) (image.Image, error) {
				return &image.RGBA{}, nil
			},
			genimgsLookupSizes: func(s3 genimgs.ListObjectsV2APIClient, conf *config.Config, imgPath string) ([]genimgs.GenImg, error) {
				return []genimgs.GenImg{
					{
						Type: "",
						Size: 100,
						URL:  "/example-100.png",
					},
				}, nil
			},
			want: `<html><head></head><body><picture><source sizes="100vw" srcset="/example-100.png 100w"/><img src="/example-100.png"/></picture></body></html>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			defer reset()

			genimgsOpen = tt.genimgsOpen
			genimgsLookupSizes = tt.genimgsLookupSizes

			err := manipulateImg(tt.s3, tt.debug, tt.conf, tt.imgtopic, htmlparsing.FindNodeByTag("img", tt.doc))
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			if diff := cmp.Diff(MustRenderNode(t, tt.doc), tt.want); diff != "" {
				t.Fatalf("Unexpected return; diff %v", diff)
			}
		})
	}
}

func Test_manipulateWithConfig(t *testing.T) {
	tests := []struct {
		description        string
		debug              bool
		conf               *config.Config
		imgtopic           *config.ImgToPicConfig
		doc                *html.Node
		s3                 *s3.Client
		genimgsOpen        func(conf *config.Config, imgPath string) (image.Image, error)
		genimgsLookupSizes func(s3 genimgs.ListObjectsV2APIClient, conf *config.Config, imgPath string) ([]genimgs.GenImg, error)
		want               string
		wantError          error
	}{
		{
			description: "replace img with picture for img tag",
			imgtopic: &config.ImgToPicConfig{
				ID:          "img",
				SourceSizes: []string{"100vw"},
			},
			doc: MustGetNode(t, `<img src="/example.png"/>`),
			genimgsOpen: func(conf *config.Config, imgPath string) (image.Image, error) {
				return &image.RGBA{}, nil
			},
			genimgsLookupSizes: func(s3 genimgs.ListObjectsV2APIClient, conf *config.Config, imgPath string) ([]genimgs.GenImg, error) {
				return []genimgs.GenImg{
					{
						Type: "",
						Size: 100,
						URL:  "/example-100.png",
					},
				}, nil
			},
			want: `<html><head></head><body><picture><source sizes="100vw" srcset="/example-100.png 100w"/><img src="/example-100.png"/></picture></body></html>`,
		},
		{
			description: "replace img with picture for class container",
			imgtopic: &config.ImgToPicConfig{
				ID:          "container",
				SourceSizes: []string{"100vw"},
			},
			doc: MustGetNode(t, `<div class="container"><img src="/example.png"/></div>`),
			genimgsOpen: func(conf *config.Config, imgPath string) (image.Image, error) {
				return &image.RGBA{}, nil
			},
			genimgsLookupSizes: func(s3 genimgs.ListObjectsV2APIClient, conf *config.Config, imgPath string) ([]genimgs.GenImg, error) {
				return []genimgs.GenImg{
					{
						Type: "",
						Size: 100,
						URL:  "/example-100.png",
					},
				}, nil
			},
			want: `<html><head></head><body><div class="container"><picture><source sizes="100vw" srcset="/example-100.png 100w"/><img src="/example-100.png"/></picture></div></body></html>`,
		},
		{
			description: "return error if manipulating the elements fails",
			imgtopic: &config.ImgToPicConfig{
				ID:          "img",
				SourceSizes: []string{"100vw"},
			},
			doc: MustGetNode(t, `<img src="/example.png"/>`),
			genimgsOpen: func(conf *config.Config, imgPath string) (image.Image, error) {
				return &image.RGBA{}, nil
			},
			genimgsLookupSizes: func(s3 genimgs.ListObjectsV2APIClient, conf *config.Config, imgPath string) ([]genimgs.GenImg, error) {
				return nil, errInjected
			},
			wantError: errInjected,
			want:      `<html><head></head><body><img src="/example.png"/></body></html>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			defer reset()

			genimgsOpen = tt.genimgsOpen
			genimgsLookupSizes = tt.genimgsLookupSizes

			err := manipulateWithConfig(tt.s3, tt.debug, tt.conf, tt.imgtopic, tt.doc)
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			if diff := cmp.Diff(MustRenderNode(t, tt.doc), tt.want); diff != "" {
				t.Fatalf("Unexpected return; diff %v", diff)
			}
		})
	}
}

func Test_Manipulator(t *testing.T) {
	tests := []struct {
		description        string
		runtime            manipulations.Runtime
		doc                *html.Node
		genimgsOpen        func(conf *config.Config, imgPath string) (image.Image, error)
		genimgsLookupSizes func(s3 genimgs.ListObjectsV2APIClient, conf *config.Config, imgPath string) ([]genimgs.GenImg, error)
		want               string
		wantError          error
	}{
		{
			description: "do nothing if should not run",
			runtime:     manipulations.Runtime{},
			doc:         MustGetNode(t, `<img/>`),
			want:        `<html><head></head><body><img/></body></html>`,
		},
		{
			description: "return error if manipulation fails",
			runtime: manipulations.Runtime{
				Config: &config.Config{
					Assets: &config.AssetsConfig{
						StaticDir: "/",
					},
					ImgToPicture: []*config.ImgToPicConfig{
						{
							ID:          "img",
							SourceSizes: []string{"100vw"},
						},
					},
				},
			},
			doc: MustGetNode(t, `<img src="/example.png"/>`),
			genimgsOpen: func(conf *config.Config, imgPath string) (image.Image, error) {
				return &image.RGBA{}, nil
			},
			genimgsLookupSizes: func(s3 genimgs.ListObjectsV2APIClient, conf *config.Config, imgPath string) ([]genimgs.GenImg, error) {
				return nil, errInjected
			},
			wantError: errInjected,
			want:      `<html><head></head><body><img src="/example.png"/></body></html>`,
		},
		{
			description: "manipulate images",
			runtime: manipulations.Runtime{
				Config: &config.Config{
					Assets: &config.AssetsConfig{
						StaticDir: "/",
					},
					ImgToPicture: []*config.ImgToPicConfig{
						{
							ID:          "img",
							SourceSizes: []string{"100vw"},
						},
					},
				},
			},
			doc: MustGetNode(t, `<img src="/example.png"/>`),
			genimgsOpen: func(conf *config.Config, imgPath string) (image.Image, error) {
				return &image.RGBA{}, nil
			},
			genimgsLookupSizes: func(s3 genimgs.ListObjectsV2APIClient, conf *config.Config, imgPath string) ([]genimgs.GenImg, error) {
				return []genimgs.GenImg{
					{
						Type: "",
						Size: 100,
						URL:  "/example-100.png",
					},
				}, nil
			},
			want: `<html><head></head><body><picture><source sizes="100vw" srcset="/example-100.png 100w"/><img src="/example-100.png"/></picture></body></html>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			defer reset()

			genimgsOpen = tt.genimgsOpen
			genimgsLookupSizes = tt.genimgsLookupSizes

			err := Manipulator(tt.runtime, tt.doc)
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Different error returned; got %v, want %v", err, tt.wantError)
			}

			if diff := cmp.Diff(MustRenderNode(t, tt.doc), tt.want); diff != "" {
				t.Fatalf("Unexpected HTML files; diff %v", diff)
			}
		})
	}
}

func MustGetNode(t *testing.T, input string) *html.Node {
	t.Helper()

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Failed to parse HTML: %v", err)
	}
	return doc
}

func MustRenderNode(t *testing.T, n *html.Node) string {
	t.Helper()

	if n == nil {
		return ""
	}

	var buf bytes.Buffer
	err := html.Render(&buf, n)
	if err != nil {
		t.Fatalf("failed to render html node to string: %v", err)
	}

	return buf.String()
}

type fileInfoStub struct {
	os.FileInfo

	IsDirReturn bool
	NameReturn  string
}

func (f *fileInfoStub) IsDir() bool {
	return f.IsDirReturn
}

func (f *fileInfoStub) Name() string {
	return f.NameReturn
}

type ImageStub struct {
	BoundsReturn image.Rectangle
}

func (i *ImageStub) ColorModel() color.Model {
	panic("At() not implemented")
}

func (i *ImageStub) Bounds() image.Rectangle {
	return i.BoundsReturn
}

func (i *ImageStub) At(x, y int) color.Color {
	panic("At() not implemented")
}
