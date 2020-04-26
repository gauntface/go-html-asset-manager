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
	"fmt"
	"strings"

	"github.com/gauntface/go-html-asset-manager/utils/html/htmlentities"
	"golang.org/x/net/html"
)

var specialEncodings = []string{
	"&amp;", // & should be first to avoid encoding other encodings
	"&lt;",
	"&gt;",
	"&quot;",
	"&apos;",
	"&grave;",
}

// Full list of characters and HTML encoding can be found
// here: https://dev.w3.org/html5/html-author/charref

func Encode(html string) string {
	// Step through the special ascii characters that must be encoded
	for _, e := range specialEncodings {
		entity := htmlentities.List[e]
		html = strings.ReplaceAll(html, entity.Characters, e)
	}

	// This is needed since the render function can't handle
	// double quotes correctly.
	for encoding, entity := range htmlentities.List {
		if len(entity.Codepoints) == 1 && entity.Codepoints[0] <= 127 {
			continue
		}
		html = strings.ReplaceAll(html, entity.Characters, encoding)
	}
	return html
}

var elementsToSkip = []string{
	"iframe",
	"noembed",
	"noframes",
	"noscript",
	"plaintext",
	"script",
	"style",
	"xmp",
}

func EncodeNodes(node *html.Node) {
	if node == nil {
		return
	}

	switch node.Type {
	case html.ElementNode:
		for _, skip := range elementsToSkip {
			if node.Data == skip {
				return
			}
		}
		if node.Data == "meta" {
			attrStrings := []string{}
			for _, a := range node.Attr {
				if a.Key == "content" {
					a.Val = Encode(a.Val)
				}
				attrStrings = append(attrStrings, fmt.Sprintf(`%v="%v"`, a.Key, a.Val))
			}
			node.Data = fmt.Sprintf("<meta %v>", strings.Join(attrStrings, " "))
			node.Type = html.RawNode
		}
		break
	case html.TextNode:
		node.Data = Encode(node.Data)
		node.Type = html.RawNode
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		EncodeNodes(child)
	}
}
