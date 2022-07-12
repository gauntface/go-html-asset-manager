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

package stripassets

import (
	"github.com/gauntface/go-html-asset-manager/v3/manipulations"
	"github.com/gauntface/go-html-asset-manager/v3/utils/html/htmlparsing"
	"golang.org/x/net/html"
)

var (
	htmlparsingFindNodesByTag = htmlparsing.FindNodesByTag
)

func Manipulator(runtime manipulations.Runtime, doc *html.Node) error {
	tags := []string{"style", "script"}
	for _, t := range tags {
		tagNodes := htmlparsingFindNodesByTag(t, doc)
		for _, t := range tagNodes {
			t.Parent.RemoveChild(t)
		}
	}

	linkNodes := htmlparsingFindNodesByTag("link", doc)
	for _, l := range linkNodes {
		a := htmlparsing.Attributes(l)
		if v, ok := a["rel"]; ok && v.Val == "stylesheet" {
			l.Parent.RemoveChild(l)
		}
	}

	return nil
}
