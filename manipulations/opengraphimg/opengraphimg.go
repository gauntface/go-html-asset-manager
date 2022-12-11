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
	"fmt"
	"strings"

	"github.com/gauntface/go-html-asset-manager/v4/assets/genimgs"
	"github.com/gauntface/go-html-asset-manager/v4/manipulations"
	"github.com/gauntface/go-html-asset-manager/v4/utils/html/htmlparsing"
	"golang.org/x/net/html"
)

const (
	RECOMMENDED_OG_IMG_WIDTH = 1200
)

var (
	genimgsLookupSizes = genimgs.LookupSizes
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

		imgs, err := genimgsLookupSizes(runtime.S3, runtime.Config, c.Val)
		if err != nil {
			return err
		}

		imgsByTypes := genimgs.GroupByType(imgs)
		imgsByType := imgsByTypes[""]
		if len(imgsByType) == 0 {
			continue
		}

		var img *genimgs.GenImg = nil
		for _, i := range imgs {
			if i.Size <= RECOMMENDED_OG_IMG_WIDTH {
				img = &i
				break
			}
		}

		if img == nil {
			continue
		}

		c.Val = fmt.Sprintf("%v%v", runtime.Config.BaseURL, img.URL)
		attributes["content"] = c
		ele.Attr = htmlparsing.AttributesList(attributes)
	}
	return nil
}
