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

package youtubeclean

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/gauntface/go-html-asset-manager/v4/manipulations"
	"github.com/gauntface/go-html-asset-manager/v4/utils/css"
	"github.com/gauntface/go-html-asset-manager/v4/utils/html/htmlparsing"
	"github.com/gauntface/go-html-asset-manager/v4/utils/html/ratiostyles"
	"golang.org/x/net/html"
)

const (
	defaultWidth  int64 = 4
	defaultHeight int64 = 3
)

var (
	embedRegex = regexp.MustCompile(`/embed/(.*).*`)

	errURLParse = errors.New("unable to parse URL")

	supportedParams = map[string]bool{
		"list": true,
	}
)

func Manipulator(runtime manipulations.Runtime, doc *html.Node) error {
	els := htmlparsing.FindNodesByTag("iframe", doc)
	for _, ele := range els {
		// Create a map of the element attributes
		attributes := map[string]html.Attribute{}
		for _, a := range ele.Attr {
			attributes[a.Key] = a
		}

		if _, ok := attributes["src"]; !ok {
			return nil
		}

		src := attributes["src"].Val

		u, err := url.Parse(src)
		if err != nil {
			return fmt.Errorf("%w %q: %v", errURLParse, src, err)
		}

		if u.Host == "" {
			if !strings.Contains(u.Path, "www.youtube.com") {
				return nil
			}
		} else if u.Host != "www.youtube.com" {
			return nil
		}

		matches := embedRegex.FindStringSubmatch(u.Path)
		if len(matches) == 0 {
			return nil
		}

		ytElement := ytElement(matches[1], queryParams(u.Query()))
		htmlparsing.SwapNodes(ele, ytElement)
	}
	return nil
}

// H/T to @paulirish for the idea via https://github.com/paulirish/lite-youtube-embed
func ytElement(videoID string, params url.Values) *html.Node {
	posterImg := &html.Node{
		Type: html.ElementNode,
		Data: "img",
		Attr: []html.Attribute{
			{
				Key: "src",
				Val: fmt.Sprintf("https://i.ytimg.com/vi/%v/hqdefault.jpg", videoID),
			},
			{
				Key: "style",
				Val: "width: 100%; height: 100%; object-fit: contain;",
			},
		},
	}

	vURL := fmt.Sprintf("https://www.youtube.com/watch?v=%v", videoID)
	paramAttribs := []html.Attribute{}
	if len(params) > 0 {
		ps := paramsString(params)
		vURL = fmt.Sprintf("%v&%v", vURL, ps)
		paramAttribs = append(paramAttribs, html.Attribute{
			Key: "videoparams",
			Val: ps,
		})
	}

	anchor := &html.Node{
		Type: html.ElementNode,
		Data: "a",
		Attr: []html.Attribute{
			{
				Key: "href",
				Val: vURL,
			},
			{
				Key: "class",
				Val: css.Format(css.ComponentType, "lite-yt", css.WithElement("link")),
			},
			{
				Key: "target",
				Val: "_blank",
			},
		},
	}
	anchor.AppendChild(posterImg)

	container := &html.Node{
		Type: html.ElementNode,
		Data: "div",
		Attr: append(
			[]html.Attribute{
				{
					Key: "class",
					Val: css.Format(css.ComponentType, "lite-yt"),
				},
				{
					Key: "videoid",
					Val: videoID,
				},
			},
			paramAttribs...,
		),
	}
	container.AppendChild(anchor)

	ratiostyles.AddAspectRatio(container, defaultWidth, defaultHeight)

	return container
}

func queryParams(params url.Values) url.Values {
	n := url.Values{}
	for k, v := range params {
		if _, ok := supportedParams[k]; !ok {
			continue
		}
		// Use the last value
		n[k] = v
	}
	return n
}

func paramsString(params url.Values) string {
	pairs := []string{}
	for k, vs := range params {
		for _, v := range vs {
			pairs = append(pairs, fmt.Sprintf("%v=%v", k, v))
		}
	}
	return strings.Join(pairs, "&")
}
