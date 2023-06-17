/**
 * Copyright 2021 Google LLC
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

package stripassets

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/gauntface/go-html-asset-manager/v5/manipulations"
	"github.com/gauntface/go-html-asset-manager/v5/utils/html/htmlparsing"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/html"
)

var errInjected = errors.New("injected error")

var reset func()

func TestMain(m *testing.M) {
	origHTMLparsingFindNodesByTag := htmlparsingFindNodesByTag

	reset = func() {
		htmlparsingFindNodesByTag = origHTMLparsingFindNodesByTag
	}

	os.Exit(m.Run())
}

func TestManipulator(t *testing.T) {
	tests := []struct {
		description string
		doc         *html.Node
		findNodes   func(tag string, node *html.Node) []*html.Node
		wantError   error
		wantHTML    string
	}{
		{
			description: "remove style tag",
			findNodes:   htmlparsing.FindNodesByTag,
			doc:         MustGetNode(t, `<style></style>`),
			wantHTML:    `<html><head></head><body></body></html>`,
		},
		{
			description: "remove script tag",
			findNodes:   htmlparsing.FindNodesByTag,
			doc:         MustGetNode(t, `<script></script>`),
			wantHTML:    `<html><head></head><body></body></html>`,
		},
		{
			description: "remove stylesheet link tag",
			findNodes:   htmlparsing.FindNodesByTag,
			doc:         MustGetNode(t, `<link rel="stylesheet"></link>`),
			wantHTML:    `<html><head></head><body></body></html>`,
		},
		{
			description: "leave non-stylesheet link tag",
			findNodes:   htmlparsing.FindNodesByTag,
			doc:         MustGetNode(t, `<link></link>`),
			wantHTML:    `<html><head><link/></head><body></body></html>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			defer reset()

			htmlparsingFindNodesByTag = tt.findNodes

			r := manipulations.Runtime{}

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
