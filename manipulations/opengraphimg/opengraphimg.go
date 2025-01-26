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
	"log"
	"sort"
	"strings"
	"sync"

	"github.com/gauntface/go-html-asset-manager/v5/assets/genimgs"
	"github.com/gauntface/go-html-asset-manager/v5/manipulations"
	"github.com/gauntface/go-html-asset-manager/v5/utils/html/htmlparsing"
	"golang.org/x/net/html"
)

const (
	RECOMMENDED_OG_IMG_WIDTH = 1200
)

var (
	genimgsLookupSizes = genimgs.LookupSizes
	imgCache           = map[string]genimgs.GenImg{}
	imgCacheLock       = sync.RWMutex{}
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

		img, err := getSuitableImg(runtime, c.Val)
		if err != nil {
			log.Printf("Warning: Unable to find suitable image for %q: %v", c.Val, err)
			continue
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

func getSuitableImg(runtime manipulations.Runtime, imgPath string) (*genimgs.GenImg, error) {
	imgCacheLock.Lock()
	defer imgCacheLock.Unlock()

	if img := lookupImgCache(imgPath); img != nil {
		return img, nil
	}

	imgs, err := genimgsLookupSizes(runtime.S3, runtime.Config, imgPath)
	if err != nil {
		return nil, err
	}

	imgsByTypes := genimgs.GroupByType(imgs)
	originalImages := imgsByTypes[""]
	if len(originalImages) == 0 {
		return nil, nil
	}

	sort.Slice(originalImages, func(i, j int) bool {
		return originalImages[i].Size > originalImages[j].Size
	})
	for _, i := range originalImages {
		if i.Size <= RECOMMENDED_OG_IMG_WIDTH {
			addImgCache(imgPath, i)
			return &i, nil
		}
	}

	return nil, nil
}

func lookupImgCache(imgPath string) *genimgs.GenImg {
	if img, ok := imgCache[imgPath]; ok {
		return &img
	}
	return nil
}

func addImgCache(imgPath string, img genimgs.GenImg) {
	imgCache[imgPath] = img
}
