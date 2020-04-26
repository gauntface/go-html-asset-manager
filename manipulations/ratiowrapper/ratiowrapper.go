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
	"fmt"
	"strconv"

	"github.com/gauntface/go-html-asset-manager/html/htmlparsing"
	"github.com/gauntface/go-html-asset-manager/html/ratiocontainer"
	"github.com/gauntface/go-html-asset-manager/manipulations"
	"golang.org/x/net/html"
)

const (
	defaultWidth  int64 = 4
	defaultHeight int64 = 3
)

func Manipulator(runtime manipulations.Runtime, doc *html.Node) error {
	allElements := getElementsToWrap(doc)
	for _, ele := range allElements {
		// Create a map of the iframes attributes
		attributes := map[string]html.Attribute{}
		for _, a := range ele.Attr {
			attributes[a.Key] = a
		}

		// Check that the iframe has a width and height attribute
		width, height, err := widthAndHeight(attributes)
		if err != nil {
			continue
		}

		// Remove the soon to be redundant attributes
		delete(attributes, "width")
		delete(attributes, "height")
		ele.Attr = []html.Attribute{}
		for _, a := range attributes {
			ele.Attr = append(ele.Attr, a)
		}

		// Remove element from it's parent so it can be wrapped
		p := ele.Parent
		s := ele.NextSibling
		p.RemoveChild(ele)

		// Wrap the element and place before it's sibling
		var wrappedElement *html.Node
		switch ele.Data {
		case "picture", "img":
			wrappedElement = ratiocontainer.WrapWithMax(ele, width, height)
		default:
			wrappedElement = ratiocontainer.Wrap(ele, width, height)
		}

		p.InsertBefore(wrappedElement, s)
	}
	return nil
}

func widthAndHeight(attributes map[string]html.Attribute) (int64, int64, error) {
	widthAttribute, ok := attributes["width"]
	if !ok {
		return 0, 0, fmt.Errorf("no width attribute")
	}
	heightAttribute, ok := attributes["height"]
	if !ok {
		return 0, 0, fmt.Errorf("no height attribute")
	}

	// Parse the width and height attributes
	width, err := strconv.ParseInt(widthAttribute.Val, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse width attribute")
	}
	height, err := strconv.ParseInt(heightAttribute.Val, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse height attribute")
	}

	return width, height, nil
}

func getElementsToWrap(doc *html.Node) []*html.Node {
	tags := []string{
		"iframe",
		"picture",
		"img",
	}
	all := []*html.Node{}
	for _, t := range tags {
		els := htmlparsing.FindNodes(t, doc)
		all = append(all, els...)
	}
	return all
}
