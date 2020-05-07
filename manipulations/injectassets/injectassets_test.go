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

package injectassets

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/gauntface/go-html-asset-manager/assets"
	"github.com/gauntface/go-html-asset-manager/assets/assetmanager"
	"github.com/gauntface/go-html-asset-manager/assets/assetstubs"
	"github.com/gauntface/go-html-asset-manager/manipulations"
	"github.com/gauntface/go-html-asset-manager/utils/html/htmlparsing"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/html"
)

var errInjected = errors.New("injected error")

var reset func()

func TestMain(m *testing.M) {
	origHTMLparsingFindNodeByTag := htmlparsingFindNodeByTag

	reset = func() {
		htmlparsingFindNodeByTag = origHTMLparsingFindNodeByTag
	}

	os.Exit(m.Run())
}

func TestPrettyPrintKeys(t *testing.T) {
	tests := []struct {
		description string
		debug       bool
		keys        []string
	}{
		{
			description: "do nothing when debug is false",
			debug:       false,
		},
		{
			description: "log keys",
			debug:       true,
			keys: []string{
				"html",
				"body",
				"c-example",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			prettyPrintKeys(tt.debug, tt.keys)
		})
	}
}

func TestManipulator(t *testing.T) {
	tests := []struct {
		description string
		debug       bool
		assets      *assetstubs.Manager
		doc         *html.Node
		findNode    func(tag string, node *html.Node) *html.Node
		wantError   error
		wantHTML    string
	}{
		{
			description: "return error if adding async css asset fails",
			findNode:    htmlparsing.FindNodeByTag,
			doc:         MustGetNode(t, `<div class="example-1 example-2"></div>`),
			assets: &assetstubs.Manager{
				WithIDReturn: map[string]map[assets.Type][]assetmanager.Asset{
					"example-1": map[assets.Type][]assetmanager.Asset{
						assets.AsyncCSS: []assetmanager.Asset{
							&assetstubs.Asset{
								URLError: errInjected,
							},
						},
					},
				},
			},
			wantError: errInjected,
		},
		{
			description: "return error if getting head node fails",
			assets:      &assetstubs.Manager{},
			findNode: func(tag string, node *html.Node) *html.Node {
				if tag == "head" {
					return nil
				}
				return &html.Node{}
			},
			doc:       MustGetNode(t, ``),
			wantError: errElementNotFound,
			wantHTML:  "<html><head></head><body></body></html>",
		},
		{
			description: "return error if getting body node fails",
			assets:      &assetstubs.Manager{},
			findNode: func(tag string, node *html.Node) *html.Node {
				if tag == "body" {
					return nil
				}
				return &html.Node{}
			},
			doc:       MustGetNode(t, ``),
			wantError: errElementNotFound,
			wantHTML:  "<html><head></head><body></body></html>",
		},
		{
			description: "return error if adding inline css asset fails",
			findNode:    htmlparsing.FindNodeByTag,
			doc:         MustGetNode(t, `<div class="example-1 example-2"></div>`),
			assets: &assetstubs.Manager{
				WithIDReturn: map[string]map[assets.Type][]assetmanager.Asset{
					"example-1": {
						assets.InlineCSS: []assetmanager.Asset{
							&assetstubs.Asset{
								TypeReturn:    assets.InlineCSS,
								ContentsError: errInjected,
							},
						},
					},
				},
			},
			wantError: errInjected,
		},
		{
			description: "return error if adding inline js asset fails",
			findNode:    htmlparsing.FindNodeByTag,
			doc:         MustGetNode(t, `<div class="example-1 example-2"></div>`),
			assets: &assetstubs.Manager{
				WithIDReturn: map[string]map[assets.Type][]assetmanager.Asset{
					"example-1": {
						assets.InlineJS: []assetmanager.Asset{
							&assetstubs.Asset{
								TypeReturn:    assets.InlineJS,
								ContentsError: errInjected,
							},
						},
					},
				},
			},
			wantError: errInjected,
		},
		{
			description: "add assets for all keys",
			findNode:    htmlparsing.FindNodeByTag,
			doc:         MustGetNode(t, `<div class="example-1 example-2"></div>`),
			assets: &assetstubs.Manager{
				WithIDReturn: map[string]map[assets.Type][]assetmanager.Asset{
					"example-1": {
						assets.InlineCSS: []assetmanager.Asset{
							&assetstubs.Asset{
								TypeReturn:     assets.InlineCSS,
								ContentsReturn: "example-1 inline contents",
							},
						},
						assets.AsyncCSS: []assetmanager.Asset{
							&assetstubs.Asset{
								TypeReturn: assets.AsyncCSS,
								URLReturn:  "/example-1-async.css",
							},
						},
					},
					"example-2": map[assets.Type][]assetmanager.Asset{
						assets.AsyncCSS: []assetmanager.Asset{
							&assetstubs.Asset{
								URLReturn: "/example-2-async.css",
							},
						},
					},
				},
			},
			wantHTML: `<html><head><style>example-1 inline contents</style></head><body><div class="example-1 example-2"></div><script>var haCSS = ['/example-1-async.css','/example-2-async.css'];</script></body></html>`,
		},
		{
			description: "add preload assets",
			findNode:    htmlparsing.FindNodeByTag,
			doc:         MustGetNode(t, `<div></div>`),
			assets: &assetstubs.Manager{
				WithIDReturn: map[string]map[assets.Type][]assetmanager.Asset{
					"div": {
						assets.PreloadCSS: []assetmanager.Asset{
							&assetstubs.Asset{
								TypeReturn: assets.PreloadCSS,
								URLReturn:  "/div-preload.css",
							},
						},
						assets.PreloadJS: []assetmanager.Asset{
							&assetstubs.Asset{
								TypeReturn: assets.PreloadJS,
								URLReturn:  "/div-preload.js",
							},
						},
					},
				},
			},
			wantHTML: `<html><head><link rel="preload" as="style" href="/div-preload.css"/><link rel="preload" as="script" href="/div-preload.js"/></head><body><div></div></body></html>`,
		},
		{
			description: "log keys if html file matches debug key",
			debug:       true,
			findNode:    htmlparsing.FindNodeByTag,
			doc:         MustGetNode(t, `<div class="example-1 example-2"></div>`),
			assets: &assetstubs.Manager{
				WithIDReturn: map[string]map[assets.Type][]assetmanager.Asset{
					"example-1": {
						assets.InlineCSS: []assetmanager.Asset{
							&assetstubs.Asset{
								TypeReturn:     assets.InlineCSS,
								ContentsReturn: "example-1 inline CSS contents",
							},
						},
						assets.InlineJS: []assetmanager.Asset{
							&assetstubs.Asset{
								TypeReturn:     assets.InlineJS,
								ContentsReturn: "example-1 inline JS contents",
							},
						},
					},
				},
			},
			wantHTML: `<html><head><style>example-1 inline CSS contents</style></head><body><div class="example-1 example-2"></div><script>example-1 inline JS contents</script></body></html>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			defer reset()

			htmlparsingFindNodeByTag = tt.findNode

			r := manipulations.Runtime{
				Assets: tt.assets,
				Debug:  tt.debug,
			}

			err := Manipulator(r, tt.doc)
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			if err != nil {
				return
			}

			if diff := cmp.Diff(MustRenderNode(t, tt.doc), tt.wantHTML); diff != "" {
				t.Fatalf("Unexpected HTML files; diff %v", diff)
			}
		})
	}
}

func TestAddInlineCSS(t *testing.T) {
	tests := []struct {
		description string
		assets      []assetmanager.Asset
		want        string
		wantError   error
	}{
		{
			description: "return error if getting contents fails",
			assets: []assetmanager.Asset{
				&assetstubs.Asset{
					ContentsError: errInjected,
				},
			},
			wantError: errInjected,
		},
		{
			description: "add asset to head",
			assets: []assetmanager.Asset{
				&assetstubs.Asset{
					ContentsReturn: `Example Content`,
				},
			},
			want: `<html><head><style>Example Content</style></head><body></body></html>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			doc := MustGetNode(t, "")
			head := htmlparsing.FindNodeByTag("head", doc)

			err := addInlineCSS(head, tt.assets)
			if !errors.Is(err, tt.wantError) {
				t.Errorf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			if err != nil {
				return
			}

			got := MustRenderNode(t, doc)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Unexpected result; Diff %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestSyncCSS(t *testing.T) {
	tests := []struct {
		description string
		asset       assetmanager.Asset
		want        string
		wantError   error
	}{
		{
			description: "return error if getting contents fails",
			asset: &assetstubs.Asset{
				URLError: errInjected,
			},
			wantError: errInjected,
		},
		{
			description: "add asset to head",
			asset: &assetstubs.Asset{
				URLReturn: "http://example.com/url.css",
			},
			want: `<html><head><link href="http://example.com/url.css" rel="stylesheet"/></head><body></body></html>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			doc := MustGetNode(t, "")
			head := htmlparsing.FindNodeByTag("head", doc)
			body := htmlparsing.FindNodeByTag("body", doc)

			err := addSyncCSS(head, body, tt.asset)
			if !errors.Is(err, tt.wantError) {
				t.Errorf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			if err != nil {
				return
			}

			got := MustRenderNode(t, doc)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Unexpected result; Diff %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestAddInlineJS(t *testing.T) {
	tests := []struct {
		description string
		asset       assetmanager.Asset
		want        string
		wantError   error
	}{
		{
			description: "return error if getting contents fails",
			asset: &assetstubs.Asset{
				ContentsError: errInjected,
			},
			wantError: errInjected,
		},
		{
			description: "add asset to head",
			asset: &assetstubs.Asset{
				ContentsReturn: "Example Content",
			},
			want: `<html><head></head><body><script>Example Content</script></body></html>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			doc := MustGetNode(t, "")
			head := htmlparsing.FindNodeByTag("head", doc)
			body := htmlparsing.FindNodeByTag("body", doc)

			err := addInlineJS(head, body, tt.asset)
			if !errors.Is(err, tt.wantError) {
				t.Errorf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			if err != nil {
				return
			}

			got := MustRenderNode(t, doc)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Unexpected result; Diff %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestAddSyncJS(t *testing.T) {
	tests := []struct {
		description string
		asset       assetmanager.Asset
		want        string
		wantError   error
	}{
		{
			description: "return error if getting contents fails",
			asset: &assetstubs.Asset{
				URLError: errInjected,
			},
			wantError: errInjected,
		},
		{
			description: "add asset to head",
			asset: &assetstubs.Asset{
				URLReturn: "http://example.com/url.js",
			},
			want: `<html><head></head><body><script src="http://example.com/url.js"></script></body></html>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			doc := MustGetNode(t, "")
			head := htmlparsing.FindNodeByTag("head", doc)
			body := htmlparsing.FindNodeByTag("body", doc)

			err := addSyncJS(head, body, tt.asset)
			if !errors.Is(err, tt.wantError) {
				t.Errorf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			if err != nil {
				return
			}

			got := MustRenderNode(t, doc)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Unexpected result; Diff %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestAddAsyncJS(t *testing.T) {
	tests := []struct {
		description string
		asset       assetmanager.Asset
		want        string
		wantError   error
	}{
		{
			description: "return error if getting contents fails",
			asset: &assetstubs.Asset{
				URLError: errInjected,
			},
			wantError: errInjected,
		},
		{
			description: "add asset to head",
			asset: &assetstubs.Asset{
				URLReturn: "http://example.com/url.js",
			},
			want: `<html><head></head><body><script src="http://example.com/url.js" async="" defer=""></script></body></html>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			doc := MustGetNode(t, "")
			head := htmlparsing.FindNodeByTag("head", doc)
			body := htmlparsing.FindNodeByTag("body", doc)

			err := addAsyncJS(head, body, tt.asset)
			if !errors.Is(err, tt.wantError) {
				t.Errorf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			if err != nil {
				return
			}

			got := MustRenderNode(t, doc)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Unexpected result; Diff %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestAddAsyncCSS(t *testing.T) {
	tests := []struct {
		description string
		assets      []assetmanager.Asset
		want        string
		wantError   error
	}{
		{
			description: "return error if getting contents fails",
			assets: []assetmanager.Asset{
				&assetstubs.Asset{
					URLError: errInjected,
				},
			},
			wantError: errInjected,
		},
		{
			description: "add asset to body",
			assets: []assetmanager.Asset{
				&assetstubs.Asset{
					URLReturn: "http://example.com/url.js",
				},
			},
			want: `<html><head></head><body><script>var haCSS = ['http://example.com/url.js'];</script></body></html>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			doc := MustGetNode(t, "")
			body := htmlparsing.FindNodeByTag("body", doc)

			err := addAsyncCSS(body, tt.assets)
			if !errors.Is(err, tt.wantError) {
				t.Errorf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			if err != nil {
				return
			}

			got := MustRenderNode(t, doc)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Unexpected result; Diff %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestAddPreloadCSS(t *testing.T) {
	tests := []struct {
		description string
		asset       assetmanager.Asset
		want        string
		wantError   error
	}{
		{
			description: "return error if getting url fails",
			asset: &assetstubs.Asset{
				URLError: errInjected,
			},
			wantError: errInjected,
		},
		{
			description: "add asset to head",
			asset: &assetstubs.Asset{
				URLReturn: "http://example.com/url.css",
			},
			want: `<html><head><link rel="preload" as="style" href="http://example.com/url.css"/></head><body></body></html>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			doc := MustGetNode(t, "")
			head := htmlparsing.FindNodeByTag("head", doc)
			body := htmlparsing.FindNodeByTag("body", doc)

			err := addPreloadCSS(head, body, tt.asset)
			if !errors.Is(err, tt.wantError) {
				t.Errorf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			if err != nil {
				return
			}

			got := MustRenderNode(t, doc)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Unexpected result; Diff %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestAddPreloadJS(t *testing.T) {
	tests := []struct {
		description string
		asset       assetmanager.Asset
		want        string
		wantError   error
	}{
		{
			description: "return error if getting url fails",
			asset: &assetstubs.Asset{
				URLError: errInjected,
			},
			wantError: errInjected,
		},
		{
			description: "add asset to head",
			asset: &assetstubs.Asset{
				URLReturn: "http://example.com/url.js",
			},
			want: `<html><head><link rel="preload" as="script" href="http://example.com/url.js"/></head><body></body></html>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			doc := MustGetNode(t, "")
			head := htmlparsing.FindNodeByTag("head", doc)
			body := htmlparsing.FindNodeByTag("body", doc)

			err := addPreloadJS(head, body, tt.asset)
			if !errors.Is(err, tt.wantError) {
				t.Errorf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			if err != nil {
				return
			}

			got := MustRenderNode(t, doc)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Unexpected result; Diff %v", cmp.Diff(got, tt.want))
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
