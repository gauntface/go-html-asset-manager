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

package ratiocontainer

import (
	"fmt"
	"strings"

	"github.com/gauntface/go-html-asset-injector/css"
	"github.com/gauntface/go-html-asset-injector/html/htmlparsing"
	"golang.org/x/net/html"
)

const (
	elementStyles = "position:absolute;width:100%;height:100%;"
)

func WrapWithMax(ele *html.Node, width, height int64) *html.Node {
	return wrap(ele, width, height, []string{fmt.Sprintf("max-width: %vpx;", width)})
}

func Wrap(ele *html.Node, width, height int64) *html.Node {
	return wrap(ele, width, height, nil)
}

func wrap(ele *html.Node, width, height int64, containerStyles []string) *html.Node {
	eleToWrap := ele
	eleToSize := ele
	if ele.Type == html.ElementNode && ele.Data == "picture" {
		eleToSize = htmlparsing.FindNode("img", ele)
	}

	var ok bool
	for i, a := range eleToSize.Attr {
		if a.Key == "style" {
			eleToSize.Attr[i].Val = fmt.Sprintf("%v;%v", a.Val, elementStyles)
			ok = true
		}
	}
	if !ok {
		eleToSize.Attr = append(eleToSize.Attr, html.Attribute{
			Key: "style",
			Val: elementStyles,
		})
	}

	defaultWrapperStyle := []string{
		"position:relative;",
		fmt.Sprintf("padding-bottom:%v%%;", (float64(height)/float64(width))*100),
	}
	wrapper := &html.Node{
		Type: html.ElementNode,
		Data: "div",
		Attr: []html.Attribute{
			{
				Key: "class",
				Val: css.Format(css.HopinNamespace, css.UtilityType, "ratio-container", "wrapper", ""),
			},
			{
				Key: "style",
				Val: strings.Join(defaultWrapperStyle, ""),
			},
		},
	}
	wrapper.AppendChild(eleToWrap)

	container := &html.Node{
		Type: html.ElementNode,
		Data: "div",
		Attr: []html.Attribute{
			{
				Key: "class",
				Val: css.Format(css.HopinNamespace, css.UtilityType, "ratio-container", "", ""),
			},
		},
	}

	if len(containerStyles) > 0 {
		container.Attr = append(container.Attr, html.Attribute{
			Key: "style",
			Val: strings.Join(containerStyles, ""),
		})
	}

	container.AppendChild(wrapper)

	return container
}
