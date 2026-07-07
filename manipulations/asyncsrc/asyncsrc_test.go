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

package asyncsrc

import (
	"bytes"
	"strings"
	"testing"

	"github.com/gauntface/go-html-asset-manager/v5/manipulations"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/html"
)

func Test_Manipulator(t *testing.T) {
	tests := []struct {
		description string
		runtime     manipulations.Runtime
		doc         *html.Node
		want        string
	}{
		{
			description: "do nothing if no iframes exist",
			doc:         MustGetNode(t, ""),
			want:        "<html><head></head><body></body></html>",
		},
		{
			description: "do nothing if iframe has no src attribute",
			doc:         MustGetNode(t, `<iframe iframeborder="0"></iframe>`),
			want:        `<html><head></head><body><iframe iframeborder="0"></iframe></body></html>`,
		},
		{
			description: "leave an empty src attribute untouched",
			doc:         MustGetNode(t, `<iframe src=""></iframe>`),
			want:        `<html><head></head><body><iframe src=""></iframe></body></html>`,
		},
		{
			description: "swap src for data-src, keeping other attributes",
			doc:         MustGetNode(t, `<iframe iframeborder="0" src="https://example.com" other="test"></iframe>`),
			want:        `<html><head></head><body><iframe iframeborder="0" other="test" data-src="https://example.com"></iframe></body></html>`,
		},
		{
			description: "swaps src for data-src on every iframe on the page",
			doc:         MustGetNode(t, `<iframe src="https://example.com/1"></iframe><iframe src="https://example.com/2"></iframe>`),
			want:        `<html><head></head><body><iframe data-src="https://example.com/1"></iframe><iframe data-src="https://example.com/2"></iframe></body></html>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			err := Manipulator(tt.runtime, tt.doc)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
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
