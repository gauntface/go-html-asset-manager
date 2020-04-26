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
	"github.com/gauntface/go-html-asset-manager/manipulations"
	"github.com/gauntface/go-html-asset-manager/utils/html/htmlparsing"
	"golang.org/x/net/html"
)

func Manipulator(runtime manipulations.Runtime, doc *html.Node) error {
	allElements := getElementsToLazyLoad(doc)
	for _, ele := range allElements {
		// Create a map of the element attributes
		attributes := map[string]html.Attribute{}
		for _, a := range ele.Attr {
			attributes[a.Key] = a
		}

		// Bail on elements that already have a loading attribute
		if _, ok := attributes["loading"]; ok {
			return nil
		}

		ele.Attr = append(ele.Attr, html.Attribute{
			Key: "loading",
			Val: "lazy",
		})
	}
	return nil
}

func getElementsToLazyLoad(doc *html.Node) []*html.Node {
	tags := []string{
		"iframe",
		"img",
	}
	all := []*html.Node{}
	for _, t := range tags {
		els := htmlparsing.FindNodes(t, doc)
		all = append(all, els...)
	}
	return all
}
