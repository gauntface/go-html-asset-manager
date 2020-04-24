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

package lazyload

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
		runtime     manipulations.Runtime
		doc         *html.Node
		want        string
		wantError   error
	}{
		{
			description: "do nothing if no iframes or images exist",
			doc:         MustGetNode(t, ""),
			want:        "<html><head></head><body></body></html>",
		},
		{
			description: "do nothing if loading attribute already exists",
			doc:         MustGetNode(t, `<img src="/example.jpg" loading="random"/>`),
			want:        `<html><head></head><body><img src="/example.jpg" loading="random"/></body></html>`,
		},
		{
			description: "add lazy loading to image",
			doc:         MustGetNode(t, `<img src="/example.jpg"/>`),
			want:        `<html><head></head><body><img src="/example.jpg" loading="lazy"/></body></html>`,
		},
		{
			description: "add lazy loading to iframe",
			doc:         MustGetNode(t, `<iframe src="/example.jpg"></iframe>`),
			want:        `<html><head></head><body><iframe src="/example.jpg" loading="lazy"></iframe></body></html>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
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
