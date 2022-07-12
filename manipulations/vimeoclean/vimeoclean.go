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

package vimeoclean

// H/T to @luwes for the idea via https://github.com/luwes/lite-vimeo-embed

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"sort"
	"strings"

	"github.com/gauntface/go-html-asset-manager/v2/manipulations"
	"github.com/gauntface/go-html-asset-manager/v2/utils/config"
	"github.com/gauntface/go-html-asset-manager/v2/utils/css"
	"github.com/gauntface/go-html-asset-manager/v2/utils/html/htmlparsing"
	"github.com/gauntface/go-html-asset-manager/v2/utils/html/ratiostyles"
	"github.com/gauntface/go-html-asset-manager/v2/utils/vimeoapi"
	"golang.org/x/net/html"
)

var (
	embedRegex = regexp.MustCompile(`/video/(.*).*`)

	errURLParse = errors.New("unable to parse URL")
)

func Manipulator(runtime manipulations.Runtime, doc *html.Node) error {
	if !runtime.HasVimeo {
		return nil
	}

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
			if !strings.Contains(u.Path, "player.vimeo.com") {
				return nil
			}
		} else if u.Host != "player.vimeo.com" {
			return nil
		}

		matches := embedRegex.FindStringSubmatch(u.Path)
		if len(matches) == 0 {
			return nil
		}

		sizes := getSizes(runtime.Config, ele)
		if len(sizes) == 0 {
			return nil
		}

		videoID := matches[1]
		video, err := runtime.Vimeo.Video(videoID)
		if err != nil {
			return err
		}
		if len(video.Pictures.Sizes) == 0 {
			return nil
		}

		viElement := vimeoElement(videoID, video, sizes)

		htmlparsing.SwapNodes(ele, viElement)
	}
	return nil
}

func getSizes(conf *config.Config, ele *html.Node) []string {
	var i *config.ImgToPicConfig
	for e := ele; e != nil; e = e.Parent {
		if e.Type != html.ElementNode {
			continue
		}

		for _, c := range conf.ImgToPicture {
			if c.ID == e.Data {
				i = c
				break
			}

			attrs := htmlparsing.Attributes(e)
			a, ok := attrs["class"]
			if !ok {
				continue
			}

			classes := strings.Split(a.Val, " ")
			for _, cl := range classes {
				if c.ID == cl {
					i = c
					break
				}
			}
		}
	}

	if i != nil {
		return i.SourceSizes
	}
	return nil
}

func vimeoElement(videoID string, video *vimeoapi.Video, sizes []string) *html.Node {
	imgElement := posterElement(video, sizes)

	anchor := &html.Node{
		Type: html.ElementNode,
		Data: "a",
		Attr: []html.Attribute{
			{
				Key: "href",
				Val: video.Link,
			},
			{
				Key: "class",
				Val: css.Format(css.ComponentType, "lite-vi", css.WithElement("link")),
			},
			{
				Key: "target",
				Val: "_blank",
			},
		},
	}
	anchor.AppendChild(imgElement)

	container := &html.Node{
		Type: html.ElementNode,
		Data: "div",
		Attr: []html.Attribute{
			{
				Key: "class",
				Val: css.Format(css.ComponentType, "lite-vi"),
			},
			{
				Key: "videoid",
				Val: videoID,
			},
		},
	}
	container.AppendChild(anchor)

	ratiostyles.AddAspectRatio(container, int64(video.Width), int64(video.Height))

	return container
}

func posterElement(video *vimeoapi.Video, sizes []string) *html.Node {
	pictures := video.Pictures.Sizes

	// Sort and add srcset attribute
	sort.Slice(pictures, func(i, j int) bool {
		return pictures[i].Width < pictures[j].Width
	})

	posterImg := &html.Node{
		Type: html.ElementNode,
		Data: "img",
		Attr: []html.Attribute{
			{
				Key: "style",
				Val: "width: 100%; height: 100%; object-fit: contain;",
			},
			{
				Key: "sizes",
				Val: strings.Join(sizes, ","),
			},
		},
	}

	srcsets := []string{}
	for _, p := range pictures {
		srcsets = append(srcsets, fmt.Sprintf("%v %vw", p.LinkWithPlayButton, p.Width))
	}

	posterImg.Attr = append(posterImg.Attr, html.Attribute{
		Key: "srcset",
		Val: strings.Join(srcsets, ","),
	})

	posterImg.Attr = append(posterImg.Attr, html.Attribute{
		Key: "src",
		Val: pictures[len(pictures)-1].LinkWithPlayButton,
	})

	altParts := []string{"Vimeo video"}
	if video.Name != "" {
		altParts = append(altParts, fmt.Sprintf("%q", video.Name))
	}
	if video.Description != "" {
		altParts = append(altParts, fmt.Sprintf("described as %q", video.Description))
	}

	posterImg.Attr = append(posterImg.Attr, html.Attribute{
		Key: "alt",
		Val: strings.Join(altParts, " "),
	})

	return posterImg
}
