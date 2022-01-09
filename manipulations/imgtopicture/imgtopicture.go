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
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gauntface/go-html-asset-manager/assets/genimgs"
	"github.com/gauntface/go-html-asset-manager/manipulations"
	"github.com/gauntface/go-html-asset-manager/utils/config"
	"github.com/gauntface/go-html-asset-manager/utils/html/htmlparsing"
	"golang.org/x/net/html"
)

var (
	errRelPath = errors.New("unable to get relative path")

	genimgsOpen        = genimgs.Open
	genimgsLookupSizes = genimgs.LookupSizes
)

func Manipulator(runtime manipulations.Runtime, doc *html.Node) error {
	if !shouldRun(runtime.Config) {
		return nil
	}

	for _, i := range runtime.Config.ImgToPicture {
		err := manipulateWithConfig(runtime.S3, runtime.Debug, runtime.Config, i, doc)
		if err != nil {
			return err
		}
	}
	return nil
}

func shouldRun(conf *config.Config) bool {
	if conf == nil {
		return false
	}

	if conf.Assets == nil || conf.Assets.StaticDir == "" || conf.Assets.GeneratedDir == "" {
		return false
	}

	if len(conf.ImgToPicture) == 0 {
		return false
	}

	return true
}

func manipulateWithConfig(s3Client *s3.Client, debug bool, conf *config.Config, imgtopic *config.ImgToPicConfig, doc *html.Node) error {
	rawElements := htmlparsing.FindNodesByTag(imgtopic.ID, doc)
	rawElements = append(rawElements, htmlparsing.FindNodesByClassname(imgtopic.ID, doc)...)

	if debug {
		fmt.Printf("Found %v raw elements for %q\n", len(rawElements), imgtopic.ID)
	}

	var imgs []*html.Node
	for _, e := range rawElements {
		imgs = append(imgs, htmlparsing.FindNodesByTag("img", e)...)
	}

	if debug {
		fmt.Printf("Found %v img elements for %q\n", len(imgs), imgtopic.ID)
	}

	for _, ie := range imgs {
		err := manipulateImg(s3Client, debug, conf, imgtopic, ie)
		if err != nil {
			return err
		}
	}
	return nil
}

func manipulateImg(s3Client *s3.Client, debug bool, conf *config.Config, imgtopic *config.ImgToPicConfig, ie *html.Node) error {
	attributes := htmlparsing.Attributes(ie)

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
	i, err := genimgsOpen(conf, srcAttr.Val)
	if err != nil {
		return nil
	}

	// Get width and height from the image
	origWidth, origHeight := i.Bounds().Size().X, i.Bounds().Size().Y

	sizes, err := genimgsLookupSizes(s3Client, conf, srcAttr.Val)
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

	pe := pictureElement(imgtopic, ie, sizes, origWidth, origHeight)

	p.InsertBefore(pe, s)
	return nil
}

func pictureElement(imgtopic *config.ImgToPicConfig, imgElement *html.Node, sizes []genimgs.GenImg, origWidth, origHeight int) *html.Node {
	sourceSetByType := genimgs.GroupByType(sizes)
	sourceSetsArray := orderedSourceSets(sourceSetByType)

	picture := &html.Node{
		Type: html.ElementNode,
		Data: "picture",
		Attr: []html.Attribute{},
	}

	for _, imgs := range sourceSetsArray {
		picture.AppendChild(createSourceElement(imgtopic, imgs))
	}

	if imgtopic.Class != "" {
		picture.Attr = append(picture.Attr, html.Attribute{
			Key: "class",
			Val: imgtopic.Class,
		})
	}

	// Replace the img src="..." attribute to point to the largest generated asset
	if len(sourceSetByType[""]) > 0 {
		ratio := float64(origHeight) / float64(origWidth)
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
	}

	picture.AppendChild(imgElement)

	return picture
}

func createSourceElement(imgtopic *config.ImgToPicConfig, imgs []genimgs.GenImg) *html.Node {
	source := &html.Node{
		Type: html.ElementNode,
		Data: "source",
		Attr: []html.Attribute{},
	}

	if len(imgs) == 0 {
		return source
	}

	// Add type attribute if appropriate
	if imgs[0].Type != "" {
		source.Attr = append(source.Attr, html.Attribute{
			Key: "type",
			Val: imgs[0].Type,
		})
	}

	// Add sizes attribute
	source.Attr = append(source.Attr, html.Attribute{
		Key: "sizes",
		Val: strings.Join(imgtopic.SourceSizes, ","),
	})

	// Sort and add srcset attribute
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

	return source
}

func orderedSourceSets(sourceSetByType map[string][]genimgs.GenImg) [][]genimgs.GenImg {
	// Order of src-set is important and we prefer webp over other formats
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

	other := [][]genimgs.GenImg{}
	for _, i := range sourceSetByType {
		other = append(other, i)
	}

	// Sort the other values to ensure tests are reliable
	sort.Slice(other, func(i, j int) bool {
		return other[i][0].Type < other[j][0].Type
	})

	sourceSets = append(sourceSets, other...)

	return sourceSets
}

// H/T to Jack Archibald for a simple and concise explainer of picture
// https://jakearchibald.com/2015/anatomy-of-responsive-images/

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
