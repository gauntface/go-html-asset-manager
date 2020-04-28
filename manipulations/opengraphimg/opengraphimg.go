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

package opengraphimg

import (
	"strings"

	"github.com/gauntface/go-html-asset-manager/assets/genimgs"
	"github.com/gauntface/go-html-asset-manager/manipulations"
	"github.com/gauntface/go-html-asset-manager/utils/html/htmlparsing"
	"golang.org/x/net/html"
)

func Manipulator(runtime manipulations.Runtime, doc *html.Node) error {
	els := htmlparsing.FindNodesByTag("meta", doc)
	for _, ele := range els {
		// Create a map of the element attributes
		attributes := map[string]html.Attribute{}
		for _, a := range ele.Attr {
			attributes[a.Key] = a
		}

		// Bail on elements that aren't property="og:image"
		p, ok := attributes["property"]
		if !ok {
			continue
		}
		if !strings.EqualFold(p.Val, "og:image") {
			continue
		}

		c, ok := attributes["content"]
		if !ok {
			continue
		}

		imgs, err := genimgs.Lookup(runtime.Config, c.Val)
		if err != nil {
			return err
		}

		imgsByTypes := genimgs.GroupByType(imgs)
		imgsByType := imgsByTypes[""]
		if len(imgs) == 0 {
			return nil
		}

		largestImg := imgsByType[len(imgsByType)-1]

		c.Val = largestImg.URL
		attributes["content"] = c
		ele.Attr = []html.Attribute{}
		for _, a := range attributes {
			ele.Attr = append(ele.Attr, a)
		}
	}
	return nil
}
