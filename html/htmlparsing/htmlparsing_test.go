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

package htmlparsing

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/gauntface/go-html-asset-manager/sets"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/html"
)

var errInjected = errors.New("injected error")

func Test_GetKeys(t *testing.T) {
	tests := []struct {
		description string
		html        string
		want        sets.StringSet
	}{
		{
			description: "return keys for empty file",
			want: sets.NewStringSet(
				"always",
				"html",
				"head",
				"body",
			),
		},
		{
			description: "return html tags",
			html:        `<section><p>Example</p><p>Dupe</p><hr /></section>`,
			want: sets.NewStringSet(
				"always",
				"html",
				"head",
				"body",
				"section",
				"p",
				"hr",
			),
		},
		{
			description: "return css class names",
			html:        `<body class="example-1 example-2"><p class="example-3"></p></body>`,
			want: sets.NewStringSet(
				"always",
				"html",
				"head",
				"body",
				"p",
				"example-1",
				"example-2",
				"example-3",
			),
		},
		{
			description: "return html attributes",
			html:        `<body value="example"><p async></p></body>`,
			want: sets.NewStringSet(
				"always",
				"html",
				"head",
				"body",
				"p",
				"value",
				"async",
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := GetKeys(MustGetNode(t, tt.html))
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}

func Test_InlineCSSTag(t *testing.T) {
	tests := []struct {
		description string
		contents    string
		want        string
	}{
		{
			description: "return inline css",
			contents:    "Hello World",
			want:        `<style>Hello World</style>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := InlineCSSTag(tt.contents)
			if diff := cmp.Diff(MustRenderNode(t, got), tt.want); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}

func Test_SyncCSSTag(t *testing.T) {
	tests := []struct {
		description string
		url         string
		want        string
	}{
		{
			description: "return link tag",
			url:         "/example.css",
			want:        `<link href="/example.css" rel="stylesheet"/>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := SyncCSSTag(tt.url)
			if diff := cmp.Diff(MustRenderNode(t, got), tt.want); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}

func Test_AsyncCSSTag(t *testing.T) {
	tests := []struct {
		description string
		urls        []string
		want        string
	}{
		{
			description: "return link tag",
			urls: []string{
				"/example-1.css",
				"/example-2.css",
			},
			want: `<script>var haCSS = ['/example-1.css','/example-2.css'];</script>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := AsyncCSSTag(tt.urls)
			if diff := cmp.Diff(MustRenderNode(t, got), tt.want); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}

func Test_InlineJSTag(t *testing.T) {
	tests := []struct {
		description string
		contents    string
		want        string
	}{
		{
			description: "return inline JS",
			contents:    "Hello World",
			want:        `<script>Hello World</script>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := InlineJSTag(tt.contents)
			if diff := cmp.Diff(MustRenderNode(t, got), tt.want); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}

func Test_SyncJSTag(t *testing.T) {
	tests := []struct {
		description string
		url         string
		want        string
	}{
		{
			description: "return script tag",
			url:         "/example.js",
			want:        `<script src="/example.js"></script>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := SyncJSTag(tt.url)
			if diff := cmp.Diff(MustRenderNode(t, got), tt.want); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}

func Test_AsyncJSTag(t *testing.T) {
	tests := []struct {
		description string
		url         string
		want        string
	}{
		{
			description: "return script tag",
			url:         "/example.js",
			want:        `<script src="/example.js" async="" defer=""></script>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := AsyncJSTag(tt.url)
			if diff := cmp.Diff(MustRenderNode(t, got), tt.want); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}

func Test_PreloadTag(t *testing.T) {
	tests := []struct {
		description string
		as          string
		url         string
		want        string
	}{
		{
			description: "return link tag",
			as:          "example-as",
			url:         "/example.css",
			want:        `<link rel="preload" as="example-as" href="/example.css"/>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := PreloadTag(tt.as, tt.url)
			if diff := cmp.Diff(MustRenderNode(t, got), tt.want); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}

func Test_FindNode(t *testing.T) {
	tests := []struct {
		description string
		tag         string
		input       string
		want        string
	}{
		{
			description: "return body tag from top level",
			tag:         `body`,
			input:       `<body>Example</body>`,
			want:        `<body>Example</body>`,
		},
		{
			description: "return nested p tag",
			tag:         `p`,
			input:       `<body><p>Example</p></body>`,
			want:        `<p>Example</p>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := FindNode(tt.tag, MustGetNode(t, tt.input))
			if diff := cmp.Diff(MustRenderNode(t, got), tt.want); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}

func Test_FindNodes(t *testing.T) {
	tests := []struct {
		description string
		tag         string
		input       string
		want        []string
	}{
		{
			description: "return single tag from top level",
			tag:         `body`,
			input:       `<body>Example</body>`,
			want:        []string{`<body>Example</body>`},
		},
		{
			description: "return nested p tag",
			tag:         `p`,
			input:       `<body><p>Example</p><section><div><p>Nested</p></div></section></body>`,
			want:        []string{`<p>Example</p>`, `<p>Nested</p>`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			gotEle := FindNodes(tt.tag, MustGetNode(t, tt.input))
			got := []string{}
			for _, g := range gotEle {
				got = append(got, MustRenderNode(t, g))
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}

func Test_SwapNodes(t *testing.T) {
	doc := MustGetNode(t, `<div></div>`)

	div := FindNode("div", doc)

	SwapNodes(div, &html.Node{
		Type: html.ElementNode,
		Data: "a",
	})

	got := MustRenderNode(t, doc)
	want := `<html><head></head><body><a></a></body></html>`

	if diff := cmp.Diff(got, want); diff != "" {
		t.Fatalf("Unexpected result; diff %v", diff)
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
