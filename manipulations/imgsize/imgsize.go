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

package imgsize

import (
	"fmt"
	"strings"

	"github.com/gauntface/go-html-asset-manager/v3/assets/genimgs"
	"github.com/gauntface/go-html-asset-manager/v3/manipulations"
	"github.com/gauntface/go-html-asset-manager/v3/utils/config"
	"github.com/gauntface/go-html-asset-manager/v3/utils/html/htmlparsing"
	"golang.org/x/net/html"
)

var (
	genimgsOpen = genimgs.Open
)

func Manipulator(runtime manipulations.Runtime, doc *html.Node) error {
	if !shouldRun(runtime.Config) {
		return nil
	}

	imgs := htmlparsing.FindNodesByTag("img", doc)
	for _, ele := range imgs {
		// Create a map of the element attributes
		attributes := htmlparsing.Attributes(ele)

		srcAttr, ok := attributes["src"]
		if !ok || srcAttr.Val == "" {
			if runtime.Debug {
				fmt.Printf("Skipping img without src\n")
			}
			continue
		}

		if strings.HasPrefix(srcAttr.Val, "http") || strings.HasPrefix(srcAttr.Val, "//") {
			if runtime.Debug {
				fmt.Printf("Skipping img with abs URL %q\n", srcAttr.Val)
			}
			continue
		}

		// Get the src image
		i, err := genimgsOpen(runtime.Config, srcAttr.Val)
		if err != nil {
			fmt.Printf("Failed to open img %q\n", srcAttr.Val)
			continue
		}

		// Get width and height from the image
		origWidth, origHeight := i.Bounds().Size().X, i.Bounds().Size().Y

		attributes["width"] = html.Attribute{
			Key: "width",
			Val: fmt.Sprintf("%v", origWidth),
		}
		attributes["height"] = html.Attribute{
			Key: "height",
			Val: fmt.Sprintf("%v", origHeight),
		}

		ele.Attr = htmlparsing.AttributesList(attributes)

	}
	return nil
}

func shouldRun(conf *config.Config) bool {
	if conf == nil {
		return false
	}

	if conf.Assets == nil || conf.Assets.StaticDir == "" {
		return false
	}

	return true
}
