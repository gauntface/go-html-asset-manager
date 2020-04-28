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

package imgtopicture

import (
	"errors"
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gauntface/go-html-asset-manager/assets/genimgs"
	"github.com/gauntface/go-html-asset-manager/manipulations"
	"github.com/gauntface/go-html-asset-manager/utils/config"
	"github.com/gauntface/go-html-asset-manager/utils/html/htmlparsing"
	"golang.org/x/net/html"
)

/*

<picture>
	<!--
		type:   Type of asset
		srcset: Img and the image width
		sizes:  Sizes the image will be displayed at
	-->
	<source
		type="image/webp"
		srcset="
			/generated/example.1234abc/100.webp 100w,
			/generated/example.1234abc/200.webp 200w,
			/generated/example.1234abc/300.webp 300w
		"
		sizes="
			(min-width: 800px) 800px,
			100vw">

	<source
		srcset="
			/generated/example.1234abc/100.jpg 100w,
			/generated/example.1234abc/200.jpg 200w,
			/generated/example.1234abc/300.jpg 300w,
		"
		sizes="
			(min-width: 800px) 800px,
			100vw">

	<!--
		src:    Only used for browsers that don't support srcset
		srcset:
	-->
	<img
		alt="Example images"
		src="/generated/example.1234abc/100.jpg">
</picture>

*/

var (
	errRelPath = errors.New("unable to get relative path")
)

func Manipulator(runtime manipulations.Runtime, doc *html.Node) error {
	if runtime.Config == nil {
		return nil
	}

	if runtime.Config.GenAssets == nil || runtime.Config.GenAssets.OutputDir == "" {
		return nil
	}

	if runtime.Config.Assets == nil || runtime.Config.Assets.BinaryDir == "" {
		return nil
	}

	if len(runtime.Config.ImgToPicture) == 0 {
		return nil
	}

	genDir := runtime.Config.GenAssets.OutputDir
	staticDir := runtime.Config.Assets.BinaryDir

	generatedDirURL, err := filepath.Rel(staticDir, genDir)
	if err != nil {
		return fmt.Errorf("%w from %q to %q: %v", errRelPath, staticDir, genDir, err)
	}

	for _, i := range runtime.Config.ImgToPicture {
		err := manipulateWithConfig(runtime.Debug, runtime.Config, runtime.Config.GenAssets, i, doc, staticDir, genDir, generatedDirURL)
		if err != nil {
			return err
		}
	}
	return nil
}

func manipulateWithConfig(debug bool, fullConf *config.Config, genConf *config.GeneratedImagesConfig, conf *config.ImgToPicConfig, doc *html.Node, staticDir, genDir, generatedDirURL string) error {
	rawElements := htmlparsing.FindNodesByTag(conf.ID, doc)
	if len(rawElements) == 0 {
		rawElements = htmlparsing.FindNodesByClassname(conf.ID, doc)
	}

	if debug {
		fmt.Printf("Found %v raw elements for %q\n", len(rawElements), conf.ID)
	}

	var imgs []*html.Node
	for _, e := range rawElements {
		imgs = append(imgs, htmlparsing.FindNodesByTag("img", e)...)
	}

	if debug {
		fmt.Printf("Found %v img elements for %q\n", len(imgs), conf.ID)
	}

	for _, ie := range imgs {
		err := manipulateImg(debug, fullConf, genConf, conf, staticDir, genDir, generatedDirURL, conf.Class, ie)
		if err != nil {
			return err
		}
	}
	return nil
}

func manipulateImg(debug bool, fullconf *config.Config, genConf *config.GeneratedImagesConfig, conf *config.ImgToPicConfig, staticDir, genDir, generatedDirURL, pictureClass string, ie *html.Node) error {
	// Create a map of the img attributes
	attributes := map[string]html.Attribute{}
	for _, a := range ie.Attr {
		attributes[a.Key] = a
	}
	srcAttr, ok := attributes["src"]
	if !ok || srcAttr.Val == "" {
		if debug {
			fmt.Printf("Skipping img without src\n")
		}
		return nil
	}

	if strings.HasPrefix(srcAttr.Val, "http") || strings.HasPrefix(srcAttr.Val, "//") {
		if debug {
			fmt.Printf("Skipping img with abs URL %q\n", srcAttr.Val)
		}
		return nil
	}

	// Get the src image
	i, err := genimgs.Open(fullconf, srcAttr.Val)
	if err != nil {
		return nil
	}

	// Get width and height from the image
	width, height := i.Bounds().Size().X, i.Bounds().Size().Y

	sizes, err := genimgs.Lookup(fullconf, srcAttr.Val)
	if err != nil {
		return err
	}

	if len(sizes) == 0 {
		if debug {
			fmt.Printf("No sizes found for %q\n", srcAttr.Val)
		}
		return nil
	}

	// Remove element from it's parent so it can be wrapped by picture
	p := ie.Parent
	s := ie.NextSibling
	p.RemoveChild(ie)

	pe := pictureElement(conf, ie, sizes, width, height, pictureClass)

	p.InsertBefore(pe, s)
	return nil
}

func pictureElement(conf *config.ImgToPicConfig, imgElement *html.Node, sizes []genimgs.GenImg, width, height int, class string) *html.Node {
	// H/T to Jack Archibald for a simple and concise explainer of picture
	// https://jakearchibald.com/2015/anatomy-of-responsive-images/

	picture := &html.Node{
		Type: html.ElementNode,
		Data: "picture",
		Attr: []html.Attribute{},
	}

	if class != "" {
		picture.Attr = append(picture.Attr, html.Attribute{
			Key: "class",
			Val: class,
		})
	}

	sourceSetByType := genimgs.GroupByType(sizes)

	desiredOrder := []string{
		"image/webp",
	}
	sourceSets := [][]genimgs.GenImg{}
	for _, dt := range desiredOrder {
		v, ok := sourceSetByType[dt]
		if !ok {
			continue
		}
		sourceSets = append(sourceSets, v)
		delete(sourceSetByType, dt)
	}

	for _, i := range sourceSetByType {
		sourceSets = append(sourceSets, i)
	}

	for _, imgs := range sourceSets {
		source := &html.Node{
			Type: html.ElementNode,
			Data: "source",
			Attr: []html.Attribute{},
		}
		if imgs[0].Type != "" {
			source.Attr = append(source.Attr, html.Attribute{
				Key: "type",
				Val: imgs[0].Type,
			})
		}

		source.Attr = append(source.Attr, html.Attribute{
			Key: "sizes",
			Val: strings.Join(conf.SourceSizes, ","),
		})

		sort.Slice(imgs, func(i, j int) bool {
			return imgs[i].Size < imgs[j].Size
		})

		srcsetValues := []string{}
		for _, s := range imgs {
			srcsetValues = append(srcsetValues, fmt.Sprintf("%v %vw", s.URL, s.Size))
		}

		source.Attr = append(source.Attr, html.Attribute{
			Key: "srcset",
			Val: strings.Join(srcsetValues, ","),
		})

		picture.AppendChild(source)
	}

	ratio := float64(height) / float64(width)
	for i, a := range imgElement.Attr {
		if a.Key != "src" {
			continue
		}
		// Change the src of the img to the largest, generated, default URL
		largest := sourceSetByType[""][len(sourceSetByType[""])-1]
		imgElement.Attr[i].Val = largest.URL

		picture.Attr = append(picture.Attr, []html.Attribute{
			{
				Key: "width",
				Val: fmt.Sprintf("%v", largest.Size),
			},
			{
				Key: "height",
				Val: fmt.Sprintf("%v", int64(ratio*float64(largest.Size))),
			},
		}...)
	}

	picture.AppendChild(imgElement)

	return picture
}
