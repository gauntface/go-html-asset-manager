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

package ratiostyles

import (
	"fmt"

	"github.com/gauntface/go-html-asset-manager/v3/utils/html/htmlparsing"
	"golang.org/x/net/html"
)

func AddAspectRatio(ele *html.Node, width, height int64) {
	eleToSize := ele
	if ele.Type == html.ElementNode && ele.Data == "picture" {
		eleToSize = htmlparsing.FindNodeByTag("img", ele)
	}

	var ok bool
	styleToAdd := fmt.Sprintf("aspect-ratio: auto %v / %v", width, height)
	for i, a := range eleToSize.Attr {
		if a.Key == "style" {
			eleToSize.Attr[i].Val = fmt.Sprintf("%v;%v", a.Val, styleToAdd)
			ok = true
		}
	}
	if !ok {
		eleToSize.Attr = append(eleToSize.Attr, html.Attribute{
			Key: "style",
			Val: styleToAdd,
		})
	}
}
