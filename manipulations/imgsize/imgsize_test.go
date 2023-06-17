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

package imgsize

import (
	"bytes"
	"errors"
	"image"
	"os"
	"strings"
	"testing"

	"github.com/gauntface/go-html-asset-manager/v5/manipulations"
	"github.com/gauntface/go-html-asset-manager/v5/utils/config"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/html"
)

var errInjected = errors.New("injected error")

var reset func()

func TestMain(m *testing.M) {
	origGenimgsOpen := genimgsOpen

	reset = func() {
		genimgsOpen = origGenimgsOpen
	}

	os.Exit(m.Run())
}

func Test_Manipulator(t *testing.T) {
	tests := []struct {
		description string
		runtime     manipulations.Runtime
		doc         *html.Node
		open        func(conf *config.Config, imgPath string) (image.Image, error)
		want        string
		wantError   error
	}{
		{
			description: "do nothing if no images exist",
			doc:         MustGetNode(t, ""),
			runtime: manipulations.Runtime{
				Config: &config.Config{
					Assets: &config.AssetsConfig{
						StaticDir: "/static",
					},
				},
			},
			want: "<html><head></head><body></body></html>",
		},
		{
			description: "add width and height to image",
			doc:         MustGetNode(t, `<img src="/example.jpg"/>`),
			runtime: manipulations.Runtime{
				Config: &config.Config{
					Assets: &config.AssetsConfig{
						StaticDir: "/static",
					},
				},
			},
			open: func(conf *config.Config, imgPath string) (image.Image, error) {
				wantPath := "/example.jpg"
				if imgPath != wantPath {
					t.Errorf("unexpected img path; got %v, want %v", imgPath, wantPath)
				}
				return &image.RGBA{
					Rect: image.Rect(0, 0, 1, 2),
				}, nil
			},
			want: `<html><head></head><body><img height="2" src="/example.jpg" width="1"/></body></html>`,
		},
		{
			description: "replace width and height to image",
			doc:         MustGetNode(t, `<img src="/example.jpg" width="1" height="2"/>`),
			runtime: manipulations.Runtime{
				Config: &config.Config{
					Assets: &config.AssetsConfig{
						StaticDir: "/static",
					},
				},
			},
			open: func(conf *config.Config, imgPath string) (image.Image, error) {
				wantPath := "/example.jpg"
				if imgPath != wantPath {
					t.Errorf("unexpected img path; got %v, want %v", imgPath, wantPath)
				}
				return &image.RGBA{
					Rect: image.Rect(0, 0, 3, 4),
				}, nil
			},
			want: `<html><head></head><body><img height="4" src="/example.jpg" width="3"/></body></html>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			t.Cleanup(reset)

			genimgsOpen = tt.open

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
