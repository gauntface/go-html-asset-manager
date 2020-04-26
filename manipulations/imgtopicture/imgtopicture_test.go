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

	"github.com/disintegration/imaging"
	"github.com/gauntface/go-html-asset-manager/manipulations"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/html"
)

var errInjected = errors.New("injected error")

var reset func()

func TestMain(m *testing.M) {
	origImagingOpen := imagingOpen
	origFileHash := filesHash
	origReadDir := ioutilReadDir

	reset = func() {
		imagingOpen = origImagingOpen
		filesHash = origFileHash
		ioutilReadDir = origReadDir
	}

	os.Exit(m.Run())
}

func Test_Manipulator(t *testing.T) {
	tests := []struct {
		description string
		runtime     manipulations.Runtime
		doc         *html.Node
		imagingOpen func(filename string, opts ...imaging.DecodeOption) (image.Image, error)
		fileHash    func(filepath string) (string, error)
		readDir     func(dirname string) ([]os.FileInfo, error)
		want        string
		wantError   error
	}{
		{
			description: "do nothing if no generated directory is defined",
			runtime:     manipulations.Runtime{},
			doc:         MustGetNode(t, `<img/>`),
			want:        `<html><head></head><body><img/></body></html>`,
		},
		{
			description: "do nothing if no static directory is defined",
			runtime: manipulations.Runtime{
				GeneratedDir: "/example",
				StaticDir:    "",
			},
			doc:  MustGetNode(t, `<img/>`),
			want: `<html><head></head><body><img/></body></html>`,
		},
		{
			description: "return error if getting relative path fails",
			runtime: manipulations.Runtime{
				GeneratedDir: "..",
				StaticDir:    "/example",
			},
			doc:       MustGetNode(t, `<img/>`),
			want:      `<html><head></head><body><img/></body></html>`,
			wantError: errRelPath,
		},
		{
			description: "do nothing if img has no src attribute",
			runtime: manipulations.Runtime{
				GeneratedDir: "/generated",
				StaticDir:    "/",
			},
			doc:  MustGetNode(t, `<img/>`),
			want: `<html><head></head><body><img/></body></html>`,
		},
		{
			description: "do nothing if img has empty src attribute",
			runtime: manipulations.Runtime{
				GeneratedDir: "/generated",
				StaticDir:    "/",
			},
			doc:  MustGetNode(t, `<img src=""/>`),
			want: `<html><head></head><body><img src=""/></body></html>`,
		},
		{
			description: "do nothing if img has http src attribute",
			runtime: manipulations.Runtime{
				GeneratedDir: "/generated",
				StaticDir:    "/",
			},
			doc:  MustGetNode(t, `<img src="http://example/example.jpg"/>`),
			want: `<html><head></head><body><img src="http://example/example.jpg"/></body></html>`,
		},
		{
			description: "do nothing if img has https src attribute",
			runtime: manipulations.Runtime{
				GeneratedDir: "/generated",
				StaticDir:    "/",
			},
			doc:  MustGetNode(t, `<img src="https://example/example.jpg"/>`),
			want: `<html><head></head><body><img src="https://example/example.jpg"/></body></html>`,
		},
		{
			description: "do nothing if img has // src attribute",
			runtime: manipulations.Runtime{
				GeneratedDir: "/generated",
				StaticDir:    "/",
			},
			doc:  MustGetNode(t, `<img src="//example/example.jpg"/>`),
			want: `<html><head></head><body><img src="//example/example.jpg"/></body></html>`,
		},
		{
			description: "do nothing if opening an img fails",
			runtime: manipulations.Runtime{
				GeneratedDir: "/static/generated",
				StaticDir:    "/static",
			},
			doc: MustGetNode(t, `<img src="/example.jpeg"/>`),
			imagingOpen: func(filename string, opts ...imaging.DecodeOption) (image.Image, error) {
				wantFilename := "/static/example.jpeg"
				if filename != wantFilename {
					t.Fatalf("Unexpected filename; got %v, want %v", filename, wantFilename)
				}
				return nil, errInjected
			},
			want: `<html><head></head><body><img src="/example.jpeg"/></body></html>`,
		},
		{
			description: "return error if file hash fails",
			runtime: manipulations.Runtime{
				GeneratedDir: "/static/generated",
				StaticDir:    "/static",
			},
			doc: MustGetNode(t, `<img src="/example.jpeg"/>`),
			imagingOpen: func(filename string, opts ...imaging.DecodeOption) (image.Image, error) {
				return &ImageStub{
					BoundsReturn: image.Rectangle{
						Min: image.Point{
							X: 0,
							Y: 0,
						},
						Max: image.Point{
							X: 10,
							Y: 10,
						},
					},
				}, nil
			},
			fileHash: func(filepath string) (string, error) {
				return "", errInjected
			},
			want:      `<html><head></head><body><img src="/example.jpeg"/></body></html>`,
			wantError: errFileHash,
		},
		{
			description: "return error if reading generate directory fails",
			runtime: manipulations.Runtime{
				GeneratedDir: "/static/generated",
				StaticDir:    "/static",
			},
			doc: MustGetNode(t, `<img src="/example.jpeg"/>`),
			imagingOpen: func(filename string, opts ...imaging.DecodeOption) (image.Image, error) {
				return &ImageStub{
					BoundsReturn: image.Rectangle{
						Min: image.Point{
							X: 0,
							Y: 0,
						},
						Max: image.Point{
							X: 10,
							Y: 10,
						},
					},
				}, nil
			},
			fileHash: func(filepath string) (string, error) {
				return "abcd123", nil
			},
			readDir: func(dirname string) ([]os.FileInfo, error) {
				wantDirname := "/static/generated/example.abcd123"
				if dirname != wantDirname {
					t.Fatalf("unexpected dirname; got %v, want %v", dirname, wantDirname)
				}
				return nil, errInjected
			},
			want:      `<html><head></head><body><img src="/example.jpeg"/></body></html>`,
			wantError: errInjected,
		},
		{
			description: "do nothing if the generated directory does not exist",
			runtime: manipulations.Runtime{
				GeneratedDir: "/static/generated",
				StaticDir:    "/static",
			},
			doc: MustGetNode(t, `<img src="/example.jpeg"/>`),
			imagingOpen: func(filename string, opts ...imaging.DecodeOption) (image.Image, error) {
				return &ImageStub{
					BoundsReturn: image.Rectangle{
						Min: image.Point{
							X: 0,
							Y: 0,
						},
						Max: image.Point{
							X: 10,
							Y: 10,
						},
					},
				}, nil
			},
			fileHash: func(filepath string) (string, error) {
				return "abcd123", nil
			},
			readDir: func(dirname string) ([]os.FileInfo, error) {
				return nil, os.ErrNotExist
			},
			want: `<html><head></head><body><img src="/example.jpeg"/></body></html>`,
		},
		{
			description: "do nothing if the generated directory contains images that are not named correctly",
			runtime: manipulations.Runtime{
				GeneratedDir: "/static/generated",
				StaticDir:    "/static",
			},
			doc: MustGetNode(t, `<img src="/example.jpeg"/>`),
			imagingOpen: func(filename string, opts ...imaging.DecodeOption) (image.Image, error) {
				return &ImageStub{
					BoundsReturn: image.Rectangle{
						Min: image.Point{
							X: 0,
							Y: 0,
						},
						Max: image.Point{
							X: 10,
							Y: 10,
						},
					},
				}, nil
			},
			fileHash: func(filepath string) (string, error) {
				return "abcd123", nil
			},
			readDir: func(dirname string) ([]os.FileInfo, error) {
				return []os.FileInfo{
					&fileInfoStub{
						IsDirReturn: true,
					},
					&fileInfoStub{
						NameReturn: "not-a-number.jpeg",
					},
				}, nil
			},
			want: `<html><head></head><body><img src="/example.jpeg"/></body></html>`,
		},
		{
			description: "replace image with picture with webp and png sources",
			runtime: manipulations.Runtime{
				GeneratedDir: "/static/generated",
				StaticDir:    "/static",
			},
			doc: MustGetNode(t, `<img src="/example.jpeg"/>`),
			imagingOpen: func(filename string, opts ...imaging.DecodeOption) (image.Image, error) {
				return &ImageStub{
					BoundsReturn: image.Rectangle{
						Min: image.Point{
							X: 0,
							Y: 0,
						},
						Max: image.Point{
							X: 10,
							Y: 10,
						},
					},
				}, nil
			},
			fileHash: func(filepath string) (string, error) {
				return "abcd123", nil
			},
			readDir: func(dirname string) ([]os.FileInfo, error) {
				return []os.FileInfo{
					&fileInfoStub{
						NameReturn: "100.webp",
					},
					&fileInfoStub{
						NameReturn: "200.webp",
					},
					&fileInfoStub{
						NameReturn: "400.png",
					},
					&fileInfoStub{
						NameReturn: "300.png",
					},
				}, nil
			},
			want: `<html><head></head><body><picture width="10" height="10"><source type="image/webp" sizes="(min-width: 800px) 800px,100vw" srcset="/generated/example.abcd123/100.webp 100w,/generated/example.abcd123/200.webp 200w"/><source sizes="(min-width: 800px) 800px,100vw" srcset="/generated/example.abcd123/300.png 300w,/generated/example.abcd123/400.png 400w"/><img src="/generated/example.abcd123/400.png"/></picture></body></html>`,
		},
		{
			description: "replace image with picture excluding max size source",
			runtime: manipulations.Runtime{
				GeneratedDir: "/static/generated",
				StaticDir:    "/static",
			},
			doc: MustGetNode(t, `<img example="other-attribute" src="/example.jpeg"/>`),
			imagingOpen: func(filename string, opts ...imaging.DecodeOption) (image.Image, error) {
				return &ImageStub{
					BoundsReturn: image.Rectangle{
						Min: image.Point{
							X: 0,
							Y: 0,
						},
						Max: image.Point{
							X: 10,
							Y: 10,
						},
					},
				}, nil
			},
			fileHash: func(filepath string) (string, error) {
				return "abcd123", nil
			},
			readDir: func(dirname string) ([]os.FileInfo, error) {
				return []os.FileInfo{
					&fileInfoStub{
						NameReturn: "100.jpg",
					},
					&fileInfoStub{
						NameReturn: "100000.jpg",
					},
				}, nil
			},
			want: `<html><head></head><body><picture width="10" height="10"><source sizes="(min-width: 800px) 800px,100vw" srcset="/generated/example.abcd123/100.jpg 100w"/><img example="other-attribute" src="/generated/example.abcd123/100.jpg"/></picture></body></html>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			defer reset()

			imagingOpen = tt.imagingOpen
			filesHash = tt.fileHash
			ioutilReadDir = tt.readDir

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
