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

package ratiowrapper

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/gauntface/go-html-asset-injector/manipulations"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/html"
)

func Test_Manipulator(t *testing.T) {
	tests := []struct {
		description string
		doc         *html.Node
		want        string
		wantError   error
	}{
		{
			description: "do not wrap iframe with no width and height using default ratio",
			doc:         MustGetNode(t, `<div><iframe src="/example.html"></iframe></div>`),
			want:        `<html><head></head><body><div><iframe src="/example.html"></iframe></div></body></html>`,
		},
		{
			description: "do not wrap iframe with no width using default ratio",
			doc:         MustGetNode(t, `<div><iframe height="1" src="/example.html"></iframe></div>`),
			want:        `<html><head></head><body><div><iframe height="1" src="/example.html"></iframe></div></body></html>`,
		},
		{
			description: "do not wrap iframe with no height using default ratio",
			doc:         MustGetNode(t, `<div><iframe width="1" src="/example.html"></iframe></div>`),
			want:        `<html><head></head><body><div><iframe width="1" src="/example.html"></iframe></div></body></html>`,
		},
		{
			description: "wrap iframe with width and height",
			doc:         MustGetNode(t, `<div><iframe width="4" height="3" src="/example.html"></iframe></div>`),
			want:        `<html><head></head><body><div><div class="n-hopin-u-ratio-container"><div class="n-hopin-u-ratio-container__wrapper" style="position:relative;padding-bottom:75%;"><iframe src="/example.html" style="position:absolute;width:100%;height:100%;"></iframe></div></div></div></body></html>`,
		},
		{
			description: "wrap all iframes with width and height",
			doc:         MustGetNode(t, `<div><iframe width="4" height="3" src="/example.html"></iframe></div><div><iframe width="16" height="9" src="/example.html"></iframe></div>`),
			want:        `<html><head></head><body><div><div class="n-hopin-u-ratio-container"><div class="n-hopin-u-ratio-container__wrapper" style="position:relative;padding-bottom:75%;"><iframe src="/example.html" style="position:absolute;width:100%;height:100%;"></iframe></div></div></div><div><div class="n-hopin-u-ratio-container"><div class="n-hopin-u-ratio-container__wrapper" style="position:relative;padding-bottom:56.25%;"><iframe src="/example.html" style="position:absolute;width:100%;height:100%;"></iframe></div></div></div></body></html>`,
		},
		{
			description: "do not wrap if width cannot be parsed",
			doc:         MustGetNode(t, `<div><iframe width="abc" height="3" src="/example.html"></iframe></div>`),
			want:        `<html><head></head><body><div><iframe width="abc" height="3" src="/example.html"></iframe></div></body></html>`,
		},
		{
			description: "do not wrap iframe if height cannot be parsed",
			doc:         MustGetNode(t, `<div><iframe width="4" height="abc" src="/example.html"></iframe></div>`),
			want:        `<html><head></head><body><div><iframe width="4" height="abc" src="/example.html"></iframe></div></body></html>`,
		},
		{
			description: "wrap picture with max size applied",
			doc:         MustGetNode(t, `<div><picture width="2" height="1"><img/></picture></div>`),
			want:        `<html><head></head><body><div><div class="n-hopin-u-ratio-container" style="max-width: 2px;"><div class="n-hopin-u-ratio-container__wrapper" style="position:relative;padding-bottom:50%;"><picture><img style="position:absolute;width:100%;height:100%;"/></picture></div></div></div></body></html>`,
		},
		{
			description: "wrap img with max size applied",
			doc:         MustGetNode(t, `<div><img width="2" height="1"/></div>`),
			want:        `<html><head></head><body><div><div class="n-hopin-u-ratio-container" style="max-width: 2px;"><div class="n-hopin-u-ratio-container__wrapper" style="position:relative;padding-bottom:50%;"><img style="position:absolute;width:100%;height:100%;"/></div></div></div></body></html>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			r := manipulations.Runtime{}

			err := Manipulator(r, tt.doc)
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Unexpected error; got %v, want %v", err, tt.wantError)
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
