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

package ratiocontainer

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/html"
)

func Test_Wrap(t *testing.T) {
	tests := []struct {
		description string
		ele         *html.Node
		width       int64
		height      int64
		want        string
	}{
		{
			description: "wrap element with existing attributes",
			ele: &html.Node{
				Type: html.ElementNode,
				Data: "example",
				Attr: []html.Attribute{
					{
						Key: "class",
						Val: "example-class",
					},
					{
						Key: "style",
						Val: "example: style",
					},
				},
			},
			width:  1,
			height: 1,
			want:   `<div class="n-hopin-u-ratio-container"><div class="n-hopin-u-ratio-container__wrapper" style="position:relative;padding-bottom:100%;"><example class="example-class" style="example: style;position:absolute;width:100%;height:100%;"></example></div></div>`,
		},
		{
			description: "wrap element with no attributes",
			ele: &html.Node{
				Type: html.ElementNode,
				Data: "example",
			},
			width:  1,
			height: 1,
			want:   `<div class="n-hopin-u-ratio-container"><div class="n-hopin-u-ratio-container__wrapper" style="position:relative;padding-bottom:100%;"><example style="position:absolute;width:100%;height:100%;"></example></div></div>`,
		},
		{
			description: "wrap picture with img that has no attributes",
			ele: &html.Node{
				Type: html.ElementNode,
				Data: "picture",
				FirstChild: &html.Node{
					Type: html.ElementNode,
					Data: "img",
				},
			},
			width:  1,
			height: 1,
			want:   `<div class="n-hopin-u-ratio-container"><div class="n-hopin-u-ratio-container__wrapper" style="position:relative;padding-bottom:100%;"><picture><img style="position:absolute;width:100%;height:100%;"/></picture></div></div>`,
		},
		{
			description: "wrap picture with img that has attributes",
			ele: &html.Node{
				Type: html.ElementNode,
				Data: "picture",
				FirstChild: &html.Node{
					Type: html.ElementNode,
					Data: "img",
					Attr: []html.Attribute{
						{
							Key: "style",
							Val: "example:123px",
						},
					},
				},
			},
			width:  1,
			height: 1,
			want:   `<div class="n-hopin-u-ratio-container"><div class="n-hopin-u-ratio-container__wrapper" style="position:relative;padding-bottom:100%;"><picture><img style="example:123px;position:absolute;width:100%;height:100%;"/></picture></div></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := Wrap(tt.ele, tt.width, tt.height)

			if diff := cmp.Diff(MustRenderNode(t, got), tt.want); diff != "" {
				t.Fatalf("Unexpected HTML files; diff %v", diff)
			}
		})
	}
}

func Test_WrapWithMax(t *testing.T) {
	tests := []struct {
		description string
		ele         *html.Node
		width       int64
		height      int64
		want        string
	}{
		{
			description: "wrap element with existing attributes",
			ele: &html.Node{
				Type: html.ElementNode,
				Data: "example",
				Attr: []html.Attribute{
					{
						Key: "class",
						Val: "example-class",
					},
					{
						Key: "style",
						Val: "example: style",
					},
				},
			},
			width:  1,
			height: 1,
			want:   `<div class="n-hopin-u-ratio-container" style="max-width: 1px;"><div class="n-hopin-u-ratio-container__wrapper" style="position:relative;padding-bottom:100%;"><example class="example-class" style="example: style;position:absolute;width:100%;height:100%;"></example></div></div>`,
		},
		{
			description: "wrap element with no attributes",
			ele: &html.Node{
				Type: html.ElementNode,
				Data: "example",
			},
			width:  1,
			height: 1,
			want:   `<div class="n-hopin-u-ratio-container" style="max-width: 1px;"><div class="n-hopin-u-ratio-container__wrapper" style="position:relative;padding-bottom:100%;"><example style="position:absolute;width:100%;height:100%;"></example></div></div>`,
		},
		{
			description: "wrap picture with img that has no attributes",
			ele: &html.Node{
				Type: html.ElementNode,
				Data: "picture",
				FirstChild: &html.Node{
					Type: html.ElementNode,
					Data: "img",
				},
			},
			width:  1,
			height: 1,
			want:   `<div class="n-hopin-u-ratio-container" style="max-width: 1px;"><div class="n-hopin-u-ratio-container__wrapper" style="position:relative;padding-bottom:100%;"><picture><img style="position:absolute;width:100%;height:100%;"/></picture></div></div>`,
		},
		{
			description: "wrap picture with img that has attributes",
			ele: &html.Node{
				Type: html.ElementNode,
				Data: "picture",
				FirstChild: &html.Node{
					Type: html.ElementNode,
					Data: "img",
					Attr: []html.Attribute{
						{
							Key: "style",
							Val: "example:123px",
						},
					},
				},
			},
			width:  1,
			height: 1,
			want:   `<div class="n-hopin-u-ratio-container" style="max-width: 1px;"><div class="n-hopin-u-ratio-container__wrapper" style="position:relative;padding-bottom:100%;"><picture><img style="example:123px;position:absolute;width:100%;height:100%;"/></picture></div></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := WrapWithMax(tt.ele, tt.width, tt.height)

			if diff := cmp.Diff(MustRenderNode(t, got), tt.want); diff != "" {
				t.Fatalf("Unexpected HTML files; diff %v", diff)
			}
		})
	}
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
