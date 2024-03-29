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
	"sort"
	"strings"

	"github.com/gauntface/go-html-asset-manager/v5/utils/sets"
	"golang.org/x/net/html"
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

func SyncCSSTag(cm CSSTagData) *html.Node {
	attr := cm.Attributes
	attr = append(attr, []html.Attribute{
		{Key: "href", Val: cm.URL},
		{Key: "rel", Val: "stylesheet"},
	}...)

	if cm.Media != "" {
		attr = append(attr, html.Attribute{Key: "media", Val: cm.Media})
	}

	return &html.Node{
		Type: html.ElementNode,
		Data: "link",
		Attr: attr,
	}
}

func AsyncCSSTag(cm CSSTagData) *html.Node {
	attr := cm.Attributes
	attr = append(attr, []html.Attribute{
		{Key: "href", Val: cm.URL},
		{Key: "rel", Val: "stylesheet"},
		{Key: "media", Val: "print"},
	}...)

	if cm.Media != "print" {
		finalMedia := "all"
		if cm.Media != "" {
			finalMedia = cm.Media
		}
		attr = append(attr, html.Attribute{
			Key: "onload",
			Val: fmt.Sprintf(`this.media='%v'`, finalMedia),
		})
	}

	return &html.Node{
		Type: html.ElementNode,
		Data: "link",
		Attr: attr,
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

func SyncJSTag(jm JSTagData) *html.Node {
	attr := jm.Attributes
	attr = append(attr, []html.Attribute{
		{Key: "src", Val: jm.URL},
	}...)
	return &html.Node{
		Type: html.ElementNode,
		Data: "script",
		Attr: attr,
	}
}

func AsyncJSTag(jm JSTagData) *html.Node {
	attr := jm.Attributes
	attr = append(attr, []html.Attribute{
		{Key: "src", Val: jm.URL},
		{Key: "async"},
		{Key: "defer"},
	}...)
	return &html.Node{
		Type: html.ElementNode,
		Data: "script",
		Attr: attr,
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

func AttributesList(attrs map[string]html.Attribute) []html.Attribute {
	attributes := []html.Attribute{}
	for _, a := range attrs {
		attributes = append(attributes, a)
	}
	sort.Slice(attributes, func(i, j int) bool {
		return attributes[i].Key < attributes[j].Key
	})
	return attributes
}

type CSSTagData struct {
	URL        string
	Attributes []html.Attribute
	Media      string
}

type JSTagData struct {
	URL        string
	Attributes []html.Attribute
}
