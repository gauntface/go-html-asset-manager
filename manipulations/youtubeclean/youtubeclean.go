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

	"github.com/gauntface/go-html-asset-manager/manipulations"
	"github.com/gauntface/go-html-asset-manager/utils/html/htmlparsing"
	"github.com/gauntface/go-html-asset-manager/utils/html/ratiocontainer"
	"golang.org/x/net/html"
)

/*
 Desired iframe:
 <iframe allow="picture-in-picture" allowfullscreen="" src="https://www.youtube.com/embed/DiIaoIcoKNY?modestbranding=1&rel=0"></iframe>
*/

const (
	defaultWidth  int64 = 4
	defaultHeight int64 = 3
)

var (
	embedRegex = regexp.MustCompile(`/embed/(.*).*`)

	errURLParse = errors.New("unable to parse URL")
)

func Manipulator(runtime manipulations.Runtime, doc *html.Node) error {
	els := htmlparsing.FindNodes("iframe", doc)
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

		ytElement := ytElement(matches[1])

		htmlparsing.SwapNodes(ele, ytElement)
	}
	return nil
}

// H/T to @paulirish for the idea via https://github.com/paulirish/lite-youtube-embed
func ytElement(videoID string) *html.Node {
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

	anchor := &html.Node{
		Type: html.ElementNode,
		Data: "a",
		Attr: []html.Attribute{
			{
				Key: "href",
				Val: fmt.Sprintf("https://www.youtube.com/watch?v=%v", videoID),
			},
			{
				Key: "class",
				Val: "n-hopin-lite-yt__link",
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
		Attr: []html.Attribute{
			{
				Key: "class",
				Val: "n-hopin-lite-yt",
			},
			{
				Key: "videoid",
				Val: videoID,
			},
		},
	}
	container.AppendChild(anchor)

	return ratiocontainer.Wrap(container, defaultWidth, defaultHeight)
}
