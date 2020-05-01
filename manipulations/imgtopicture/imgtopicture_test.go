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

	"github.com/gauntface/go-html-asset-manager/assets/genimgs"
	"github.com/gauntface/go-html-asset-manager/manipulations"
	"github.com/gauntface/go-html-asset-manager/utils/config"
	"github.com/gauntface/go-html-asset-manager/utils/html/htmlparsing"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/html"
)

var errInjected = errors.New("injected error")

var reset func()

func TestMain(m *testing.M) {
	reset = func() {
	}

	os.Exit(m.Run())
}

func Test_Manipulator(t *testing.T) {
	tests := []struct {
		description string
		runtime     manipulations.Runtime
		doc         *html.Node
		want        string
		wantError   error
	}{
		{
			description: "do nothing if should not run",
			runtime:     manipulations.Runtime{},
			doc:         MustGetNode(t, `<img/>`),
			want:        `<html><head></head><body><img/></body></html>`,
		},

		{
			description: "do nothing if img has no src attribute",
			runtime: manipulations.Runtime{
				Config: &config.Config{
					GenAssets: &config.GeneratedImagesConfig{
						OutputDir: "/generated",
					},
					Assets: &config.AssetsConfig{
						StaticDir: "/",
					},
				},
			},
			doc:  MustGetNode(t, `<img/>`),
			want: `<html><head></head><body><img/></body></html>`,
		},
		{
			description: "do nothing if img has empty src attribute",
			runtime: manipulations.Runtime{
				Config: &config.Config{
					Assets: &config.AssetsConfig{
						StaticDir:    "/",
						GeneratedDir: "/generated",
					},
					ImgToPicture: []*config.ImgToPicConfig{
						{
							ID:          ".example",
							MaxWidth:    800,
							SourceSizes: []string{"(min-width: 800px) 800px", "100vw"},
						},
					},
				},
			},
			doc:  MustGetNode(t, `<img src=""/>`),
			want: `<html><head></head><body><img src=""/></body></html>`,
		},
		{
			description: "do nothing if img has http src attribute",
			runtime: manipulations.Runtime{
				Config: &config.Config{
					GenAssets: &config.GeneratedImagesConfig{
						OutputDir: "/generated",
					},
					Assets: &config.AssetsConfig{
						StaticDir: "/",
					},
				},
			},
			doc:  MustGetNode(t, `<img src="http://example/example.jpg"/>`),
			want: `<html><head></head><body><img src="http://example/example.jpg"/></body></html>`,
		},
		{
			description: "do nothing if img has https src attribute",
			runtime: manipulations.Runtime{
				Config: &config.Config{
					GenAssets: &config.GeneratedImagesConfig{
						OutputDir: "/generated",
					},
					Assets: &config.AssetsConfig{
						StaticDir: "/",
					},
				},
			},
			doc:  MustGetNode(t, `<img src="https://example/example.jpg"/>`),
			want: `<html><head></head><body><img src="https://example/example.jpg"/></body></html>`,
		},
		{
			description: "do nothing if img has // src attribute",
			runtime: manipulations.Runtime{
				Config: &config.Config{
					GenAssets: &config.GeneratedImagesConfig{
						OutputDir: "/generated",
					},
					Assets: &config.AssetsConfig{
						StaticDir: "/",
					},
				},
			},
			doc:  MustGetNode(t, `<img src="//example/example.jpg"/>`),
			want: `<html><head></head><body><img src="//example/example.jpg"/></body></html>`,
		},
		{
			description: "do nothing if opening an img fails",
			runtime: manipulations.Runtime{
				Config: &config.Config{
					GenAssets: &config.GeneratedImagesConfig{
						OutputDir: "/generated",
					},
					Assets: &config.AssetsConfig{
						StaticDir: "/",
					},
				},
			},
			doc:  MustGetNode(t, `<img src="/example.jpeg"/>`),
			want: `<html><head></head><body><img src="/example.jpeg"/></body></html>`,
		},
		/* {
			description: "return error if file hash fails",
			runtime: manipulations.Runtime{
				Config: &config.Config{
					GenAssets: &config.GeneratedImagesConfig{
						OutputDir: "/generated",
					},
					Assets: &config.AssetsConfig{
						StaticDir: "/",
					},
				},
			},
			doc: MustGetNode(t, `<img src="/example.jpeg"/>`),
			want:      `<html><head></head><body><img src="/example.jpeg"/></body></html>`,
			wantError: errFileHash,
		},*/
		/* {
			description: "return error if reading generate directory fails",
			runtime: manipulations.Runtime{
				Config: &config.Config{
					GenAssets: &config.GeneratedImagesConfig{
						OutputDir: "/static/generated",
					},
					Assets: &config.AssetsConfig{
						StaticDir: "/static",
					},
				},
			},
			doc: MustGetNode(t, `<img src="/example.jpeg"/>`),
			want:      `<html><head></head><body><img src="/example.jpeg"/></body></html>`,
			wantError: errInjected,
		},*/
		{
			description: "do nothing if the generated directory does not exist",
			runtime: manipulations.Runtime{
				Config: &config.Config{
					GenAssets: &config.GeneratedImagesConfig{
						OutputDir: "/static/generated",
					},
					Assets: &config.AssetsConfig{
						StaticDir: "/static",
					},
				},
			},
			doc:  MustGetNode(t, `<img src="/example.jpeg"/>`),
			want: `<html><head></head><body><img src="/example.jpeg"/></body></html>`,
		},
		{
			description: "do nothing if the generated directory contains images that are not named correctly",
			runtime: manipulations.Runtime{
				Config: &config.Config{
					GenAssets: &config.GeneratedImagesConfig{
						OutputDir: "/static/generated",
					},
					Assets: &config.AssetsConfig{
						StaticDir: "/static",
					},
				},
			},
			doc:  MustGetNode(t, `<img src="/example.jpeg"/>`),
			want: `<html><head></head><body><img src="/example.jpeg"/></body></html>`,
		},
		/*{
			description: "replace image with picture with webp and png sources",
			runtime: manipulations.Runtime{
				Debug: true,
				Config: &config.Config{
					GenAssets: &config.GeneratedImagesConfig{
						OutputDir:  "/static/generated",
						MaxDensity: 3,
					},
					Assets: &config.AssetsConfig{
						StaticDir: "/static",
					},
					ImgToPicture: []*config.ImgToPicConfig{
						{
							ID:       "img",
							MaxWidth: 800,
							SourceSizes: []string{
								"(min-width: 800px) 800px",
								"100vw",
							},
						},
					},
				},
			},
			doc: MustGetNode(t, `<img src="/example.jpeg"/>`),
			want: `<html><head></head><body><picture width="400" height="400"><source type="image/webp" sizes="(min-width: 800px) 800px,100vw" srcset="/generated/example.abcd123/100.webp 100w,/generated/example.abcd123/200.webp 200w"/><source sizes="(min-width: 800px) 800px,100vw" srcset="/generated/example.abcd123/300.png 300w,/generated/example.abcd123/400.png 400w"/><img src="/generated/example.abcd123/400.png"/></picture></body></html>`,
		},
		{
			description: "replace image with picture excluding max size source",
			runtime: manipulations.Runtime{
				Config: &config.Config{
					GenAssets: &config.GeneratedImagesConfig{
						OutputDir:  "/static/generated",
						MaxDensity: 3,
					},
					Assets: &config.AssetsConfig{
						StaticDir: "/static",
					},
					ImgToPicture: []*config.ImgToPicConfig{
						{
							ID:       "example-selector",
							MaxWidth: 800,
							SourceSizes: []string{
								"(min-width: 800px) 800px",
								"100vw",
							},
						},
					},
				},
			},
			doc: MustGetNode(t, `<img class="example-selector" example="other-attribute" src="/example.jpeg"/>`),
			want: `<html><head></head><body><picture width="100" height="100"><source sizes="(min-width: 800px) 800px,100vw" srcset="/generated/example.abcd123/100.jpg 100w"/><img class="example-selector" example="other-attribute" src="/generated/example.abcd123/100.jpg"/></picture></body></html>`,
		},*/
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			defer reset()

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
					StaticDir:    "/",
					GeneratedDir: "",
				},
			},
			want: false,
		},
		{
			description: "return false for config without ImgToPicture",
			conf: &config.Config{
				Assets: &config.AssetsConfig{
					StaticDir:    "/",
					GeneratedDir: "/generated",
				},
				ImgToPicture: nil,
			},
			want: false,
		},
		{
			description: "return true for config required params",
			conf: &config.Config{
				Assets: &config.AssetsConfig{
					StaticDir:    "/",
					GeneratedDir: "/generated",
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
				"image/jpg": {
					{
						Type: "image/jpg",
						URL:  "/image.jpg",
					},
				},
				"image/png": {
					{
						Type: "image/png",
						URL:  "/image.png",
					},
				},
				"image/webp": {
					{
						Type: "image/webp",
						URL:  "/image.webp",
					},
				},
			},
			want: [][]genimgs.GenImg{
				{
					{
						Type: "image/webp",
						URL:  "/image.webp",
					},
				},
				{
					{
						Type: "image/jpg",
						URL:  "/image.jpg",
					},
				},
				{
					{
						Type: "image/png",
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
					Type: "image/png",
					Size: 1,
					URL:  "/example-1.png",
				},
				{
					Type: "image/png",
					Size: 2,
					URL:  "/example-2.png",
				},
			},
			origWidth:  3,
			origHeight: 3,
			want:       `<picture><source type="image/png" sizes="min-width(800px) 800px,100vw" srcset="/example-1.png 1w,/example-2.png 2w"/><img src="/example.png"/></picture>`,
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
			want:       `<picture class="picture-class" width="2" height="2"><source sizes="min-width(800px) 800px,100vw" srcset="/example-1.png 1w,/example-2.png 2w"/><img src="/example-2.png" class="example"/></picture>`,
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
		description string
		debug       bool
		conf        *config.Config
		imgtopic    *config.ImgToPicConfig
		doc         *html.Node
		want        string
		wantError   error
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
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			err := manipulateImg(tt.debug, tt.conf, tt.imgtopic, htmlparsing.FindNodeByTag("img", tt.doc))
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			if diff := cmp.Diff(MustRenderNode(t, tt.doc), tt.want); diff != "" {
				t.Fatalf("Unexpected return; diff %v", diff)
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
