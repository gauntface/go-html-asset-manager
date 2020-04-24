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

package htmlencoding

import (
	"bytes"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/html"
)

// These HTML entities are used by Goldmark
// https://github.com/yuin/goldmark/blob/12fc98ebcd751822bbd43314119e483472e6f55f/extension/typographer.go
// This means that markdown -> HTML Entities -> *html.Node
// can result in these characters that then need
// to be converted back to HTML entities.

func Test_Encode(t *testing.T) {
	tests := []struct {
		description string
		input       string
		want        string
	}{
		{
			description: "Encode all special ascii entities",
			input:       "& < > \" ' and `",
			want:        `&amp; &lt; &gt; &quot; &apos; and &grave;`,
		},
		{
			description: "Encode all golmark entities",
			input: `
“
”
‘
’
–
—
…
«
»
`,
			want: `
&ldquo;
&rdquo;
&lsquo;
&rsquo;
&ndash;
&mdash;
&mldr;
&laquo;
&raquo;
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := Encode(tt.input)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Unexpected result; Diff %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func Test_EncodeNodes(t *testing.T) {
	tests := []struct {
		description string
		html        *html.Node
		want        string
	}{
		{
			description: "do nothing if the node is nil",
			html:        nil,
			want:        "",
		},
		{
			description: "encode text nodes",
			html:        MustGetNode(t, `<p>Example Text Node - & < > " " '</p>`),
			want:        "<html><head></head><body><p>Example Text Node - &amp; &lt; &gt; &quot; &quot; &apos;</p></body></html>",
		},
		{
			description: "ignore elements in skip list",
			html:        MustGetNode(t, `<style>Example Text Node - & < > " " '</style>`),
			want:        `<html><head><style>Example Text Node - & < > " " '</style></head><body></body></html>`,
		},
		{
			description: "ignore elements in skip list",
			html:        MustGetNode(t, `<meta content="- & < > '">`),
			want:        `<html><head><meta content="- &amp; &lt; &gt; &apos;"></head><body></body></html>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			EncodeNodes(tt.html)
			if diff := cmp.Diff(MustRenderNode(t, tt.html), tt.want); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
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
