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

package iframedefaultsize

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/gauntface/go-html-asset-manager/v2/manipulations"
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
			description: "do nothing for iframe with valid width and height",
			doc:         MustGetNode(t, `<iframe width="1" height="1"></iframe>`),
			want:        `<html><head></head><body><iframe height="1" width="1"></iframe></body></html>`,
		},
		{
			description: "apply default width and height if iframe has no attributes",
			doc:         MustGetNode(t, `<iframe></iframe>`),
			want:        `<html><head></head><body><iframe height="3" width="4"></iframe></body></html>`,
		},
		{
			description: "apply default width and height if iframe has just width attribute",
			doc:         MustGetNode(t, `<iframe width="1"></iframe>`),
			want:        `<html><head></head><body><iframe height="3" width="4"></iframe></body></html>`,
		},
		{
			description: "apply default width and height if iframe has just height attribute",
			doc:         MustGetNode(t, `<iframe height="1"></iframe>`),
			want:        `<html><head></head><body><iframe height="3" width="4"></iframe></body></html>`,
		},
		{
			description: "apply default width and height if iframe has invalid width attribute",
			doc:         MustGetNode(t, `<iframe width="abc" height="1"></iframe>`),
			want:        `<html><head></head><body><iframe height="3" width="4"></iframe></body></html>`,
		},
		{
			description: "apply default width and height if iframe has invalid height attribute",
			doc:         MustGetNode(t, `<iframe width="1" height="abc"></iframe>`),
			want:        `<html><head></head><body><iframe height="3" width="4"></iframe></body></html>`,
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
