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

package asyncsrc

import (
	"github.com/gauntface/go-html-asset-manager/v5/manipulations"
	"github.com/gauntface/go-html-asset-manager/v5/utils/html/htmlparsing"
	"golang.org/x/net/html"
)

func Manipulator(runtime manipulations.Runtime, doc *html.Node) error {
	allElements := htmlparsing.FindNodesByTag("iframe", doc)
	for _, ele := range allElements {
		for i, a := range ele.Attr {
			if a.Key != "src" {
				continue
			}
			if a.Val == "" {
				continue
			}
			ele.Attr = append(ele.Attr, html.Attribute{
				Key: "data-src",
				Val: a.Val,
			})
			ele.Attr = append(ele.Attr[:i], ele.Attr[i+1:]...)
		}
	}
	return nil
}
