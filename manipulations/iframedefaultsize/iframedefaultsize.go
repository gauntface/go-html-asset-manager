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
	"fmt"
	"sort"
	"strconv"

	"github.com/gauntface/go-html-asset-injector/html/htmlparsing"
	"github.com/gauntface/go-html-asset-injector/manipulations"
	"golang.org/x/net/html"
)

const (
	defaultWidth  int64 = 4
	defaultHeight int64 = 3
)

func Manipulator(runtime manipulations.Runtime, doc *html.Node) error {
	allElements := htmlparsing.FindNodes("iframe", doc)
	for _, ele := range allElements {
		// Create a map of the iframes attributes
		attributes := map[string]html.Attribute{}
		for _, a := range ele.Attr {
			attributes[a.Key] = a
		}

		// Check that the iframe has a width and height attribute
		width, height := widthAndHeight(attributes)

		// Update / set the width and height attributes to the element
		attributes["width"] = html.Attribute{
			Key: "width",
			Val: fmt.Sprintf("%v", width),
		}
		attributes["height"] = html.Attribute{
			Key: "height",
			Val: fmt.Sprintf("%v", height),
		}
		ele.Attr = []html.Attribute{}
		for _, a := range attributes {
			ele.Attr = append(ele.Attr, a)
		}
		sort.Slice(ele.Attr, func(i, j int) bool {
			return ele.Attr[i].Key < ele.Attr[j].Key
		})
	}
	return nil
}

func widthAndHeight(attributes map[string]html.Attribute) (int64, int64) {
	widthAttribute, ok := attributes["width"]
	if !ok {
		return defaultWidth, defaultHeight
	}
	heightAttribute, ok := attributes["height"]
	if !ok {
		return defaultWidth, defaultHeight
	}

	// Parse the width and height attributes
	width, err := strconv.ParseInt(widthAttribute.Val, 10, 64)
	if err != nil {
		return defaultWidth, defaultHeight
	}
	height, err := strconv.ParseInt(heightAttribute.Val, 10, 64)
	if err != nil {
		return defaultWidth, defaultHeight
	}

	return width, height
}
