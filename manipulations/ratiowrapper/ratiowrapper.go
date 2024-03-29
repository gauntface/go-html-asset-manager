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

	"github.com/gauntface/go-html-asset-manager/v5/manipulations"
	"github.com/gauntface/go-html-asset-manager/v5/utils/html/htmlparsing"
	"github.com/gauntface/go-html-asset-manager/v5/utils/html/ratiostyles"
	"golang.org/x/net/html"
)

func Manipulator(runtime manipulations.Runtime, doc *html.Node) error {
	if runtime.Config == nil || len(runtime.Config.RatioWrapper) == 0 {
		return nil
	}

	allElements := getElementsToWrap(runtime.Config.RatioWrapper, doc)
	for _, ele := range allElements {
		// Create a map of the iframes attributes
		attributes := map[string]html.Attribute{}
		for _, a := range ele.Attr {
			attributes[a.Key] = a
		}

		// Check that the element has a width and height attribute
		width, height, err := widthAndHeight(attributes)
		if err != nil {
			continue
		}

		// Remove the soon to be redundant attributes
		delete(attributes, "width")
		delete(attributes, "height")
		ele.Attr = htmlparsing.AttributesList(attributes)

		ratiostyles.AddAspectRatio(ele, width, height)
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

func getElementsToWrap(queries []string, doc *html.Node) []*html.Node {
	var rawElements []*html.Node
	for _, q := range queries {
		rawElements = append(rawElements, htmlparsing.FindNodesByTag(q, doc)...)
		rawElements = append(rawElements, htmlparsing.FindNodesByClassname(q, doc)...)
	}

	tags := []string{
		"iframe",
		"img",
	}
	all := []*html.Node{}

	for _, e := range rawElements {
		for _, t := range tags {
			els := htmlparsing.FindNodesByTag(t, e)
			all = append(all, els...)
		}
	}
	return all
}
