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
	"fmt"
	"strings"

	"github.com/gauntface/go-html-asset-manager/utils/sets"
	"golang.org/x/net/html"
)

const (
	AsyncCSSClassname = "n-hopin-u-async-css"
)

func GetKeys(node *html.Node) sets.StringSet {
	keys := sets.NewStringSet("always")
	if node.Type == html.ElementNode {
		// Data is the HTML element tag (i.e. body / <body>)
		keys.Add(node.Data)

		for _, a := range node.Attr {
			if a.Key == "class" {
				classes := strings.Split(a.Val, " ")
				for _, c := range classes {
					keys.Add(strings.TrimSpace(c))
				}
			} else {
				keys.Add(a.Key)
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		cks := GetKeys(child)
		keys.Merge(cks)
	}
	return keys
}

func InlineCSSTag(contents string) *html.Node {
	return &html.Node{
		Type: html.ElementNode,
		Data: "style",
		FirstChild: &html.Node{
			Type: html.TextNode,
			Data: contents,
		},
	}
}

func SyncCSSTag(url string) *html.Node {
	return &html.Node{
		Type: html.ElementNode,
		Data: "link",
		Attr: []html.Attribute{
			{Key: "href", Val: url},
			{Key: "rel", Val: "stylesheet"},
		},
	}
}

func AsyncCSSTag(urls []string) *html.Node {
	asyncCSS := []string{}
	for _, url := range urls {
		asyncCSS = append(asyncCSS, fmt.Sprintf("'%v'", url))
	}

	js := fmt.Sprintf(
		`var haCSS = [%v];`,
		strings.Join(asyncCSS, ","),
	)

	return &html.Node{
		Type: html.ElementNode,
		Data: "script",
		FirstChild: &html.Node{
			Type: html.TextNode,
			Data: js,
		},
	}
}

func InlineJSTag(contents string) *html.Node {
	return &html.Node{
		Type: html.ElementNode,
		Data: "script",
		FirstChild: &html.Node{
			Type: html.TextNode,
			Data: contents,
		},
	}
}

func SyncJSTag(url string) *html.Node {
	return &html.Node{
		Type: html.ElementNode,
		Data: "script",
		Attr: []html.Attribute{
			{Key: "src", Val: url},
		},
	}
}

func AsyncJSTag(url string) *html.Node {
	return &html.Node{
		Type: html.ElementNode,
		Data: "script",
		Attr: []html.Attribute{
			{Key: "src", Val: url},
			{Key: "async"},
			{Key: "defer"},
		},
	}
}

func PreloadTag(as, url string) *html.Node {
	return &html.Node{
		Type: html.ElementNode,
		Data: "link",
		Attr: []html.Attribute{
			{Key: "rel", Val: "preload"},
			{Key: "as", Val: as},
			{Key: "href", Val: url},
		},
	}
}

func FindNodeByTag(tag string, node *html.Node) *html.Node {
	if node.Type == html.ElementNode {
		if node.Data == tag {
			return node
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		e := FindNodeByTag(tag, child)
		if e != nil {
			return e
		}
	}
	return nil
}

func FindNodesByTag(tag string, node *html.Node) []*html.Node {
	elements := []*html.Node{}
	if node.Type == html.ElementNode {
		if node.Data == tag {
			elements = append(elements, node)
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		ce := FindNodesByTag(tag, child)
		elements = append(elements, ce...)
	}
	return elements
}

func FindNodesByClassname(classname string, node *html.Node) []*html.Node {
	elements := []*html.Node{}
	if node.Type == html.ElementNode {
		for _, a := range node.Attr {
			if a.Key == "class" {
				classes := strings.Split(a.Val, " ")
				for _, c := range classes {
					if c == classname {
						elements = append(elements, node)
					}
				}
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		ce := FindNodesByClassname(classname, child)
		elements = append(elements, ce...)
	}
	return elements
}

func SwapNodes(original, new *html.Node) {
	p := original.Parent
	s := original.NextSibling
	p.RemoveChild(original)

	p.InsertBefore(new, s)
}

func Attributes(e *html.Node) map[string]html.Attribute {
	attributes := map[string]html.Attribute{}
	for _, a := range e.Attr {
		attributes[a.Key] = a
	}
	return attributes
}