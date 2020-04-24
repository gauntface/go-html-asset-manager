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
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/gauntface/go-html-asset-injector/files"
	"github.com/gauntface/go-html-asset-injector/html/htmlparsing"
	"github.com/gauntface/go-html-asset-injector/manipulations"
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
	maxSize     int64 = 2400 // (800 x 3) (<max width> * <max density>)
	sourceSizes       = []string{
		"(min-width: 800px) 800px",
		"100vw",
	}

	errRelPath  = errors.New("unable to get relative path")
	errFileHash = errors.New("failed to get file hash")

	imagingOpen   = imaging.Open
	filesHash     = files.Hash
	ioutilReadDir = ioutil.ReadDir
)

func Manipulator(runtime manipulations.Runtime, doc *html.Node) error {
	if runtime.GeneratedDir == "" {
		return nil
	}

	if runtime.StaticDir == "" {
		return nil
	}

	generatedDirURL, err := filepath.Rel(runtime.StaticDir, runtime.GeneratedDir)
	if err != nil {
		return fmt.Errorf("%w from %q to %q: %v", errRelPath, runtime.StaticDir, runtime.GeneratedDir, err)
	}

	imgElements := htmlparsing.FindNodes("img", doc)
	for _, ie := range imgElements {
		err := manipulateImg(runtime, generatedDirURL, ie)
		if err != nil {
			return err
		}
	}
	return nil
}

func manipulateImg(runtime manipulations.Runtime, generatedDirURL string, ie *html.Node) error {
	// Create a map of the img attributes
	attributes := map[string]html.Attribute{}
	for _, a := range ie.Attr {
		attributes[a.Key] = a
	}
	srcAttr, ok := attributes["src"]
	if !ok || srcAttr.Val == "" {
		return nil
	}

	if strings.HasPrefix(srcAttr.Val, "http") || strings.HasPrefix(srcAttr.Val, "//") {
		return nil
	}

	srcPath := filepath.Join(runtime.StaticDir, srcAttr.Val)

	// Get the src image
	i, err := imagingOpen(srcPath)
	if err != nil {
		return nil
	}

	// Get width and height from the image
	width, height := i.Bounds().Size().X, i.Bounds().Size().Y

	hash, err := filesHash(srcPath)
	if err != nil {
		return fmt.Errorf("%w for img %q", errFileHash, srcPath)
	}

	// Get available sizes of the image
	sizes, err := getImageSizes(runtime.GeneratedDir, srcPath, hash, generatedDirURL)
	if err != nil {
		return err
	}

	if len(sizes) == 0 {
		return nil
	}

	// Remove element from it's parent so it can be wrapped by picture
	p := ie.Parent
	s := ie.NextSibling
	p.RemoveChild(ie)

	pe := pictureElement(ie, sizes, width, height)

	p.InsertBefore(pe, s)
	return nil
}

func getImageSizes(generatedDir, src, hash, generatedDirURL string) ([]generatedImage, error) {
	filename := filepath.Base(src)
	filename = strings.TrimSuffix(filename, filepath.Ext(filename))
	genSubDir := fmt.Sprintf("%v.%v", filename, hash)
	dirPath := filepath.Join(generatedDir, genSubDir)

	contents, err := ioutilReadDir(dirPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}

	imgs := []generatedImage{}
	for _, c := range contents {
		if c.IsDir() {
			continue
		}

		ext := filepath.Ext(c.Name())
		filename := strings.TrimSuffix(c.Name(), ext)
		size, err := strconv.ParseInt(filename, 10, 64)
		if err != nil {
			continue
		}

		if size > maxSize {
			continue
		}

		var typ string
		switch ext {
		case ".webp":
			typ = "image/webp"
			break
		}

		imgs = append(imgs, generatedImage{
			URL:  filepath.Join("/", generatedDirURL, genSubDir, c.Name()),
			Type: typ,
			Size: size,
		})
	}

	return imgs, nil
}

func pictureElement(imgElement *html.Node, sizes []generatedImage, width, height int) *html.Node {
	// H/T to Jack Archibald for a simple and concise explainer of picture
	// https://jakearchibald.com/2015/anatomy-of-responsive-images/

	picture := &html.Node{
		Type: html.ElementNode,
		Data: "picture",
		Attr: []html.Attribute{
			{
				Key: "width",
				Val: fmt.Sprintf("%v", width),
			},
			{
				Key: "height",
				Val: fmt.Sprintf("%v", height),
			},
		},
	}

	sourceSetByType := map[string][]generatedImage{}
	for _, s := range sizes {
		_, ok := sourceSetByType[s.Type]
		if !ok {
			sourceSetByType[s.Type] = []generatedImage{}
		}

		sourceSetByType[s.Type] = append(sourceSetByType[s.Type], s)
	}

	sourceSets := sortSourceSets(sourceSetByType)

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
			Val: strings.Join(sourceSizes, ","),
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

	if len(sourceSetByType[""]) > 0 {
		for i, a := range imgElement.Attr {
			if a.Key != "src" {
				continue
			}
			// Change the src of the img to the largest, generated, default URL
			imgElement.Attr[i].Val = sourceSetByType[""][len(sourceSetByType[""])-1].URL
		}
	}

	picture.AppendChild(imgElement)

	return picture
}

func sortSourceSets(sourceSetByType map[string][]generatedImage) [][]generatedImage {
	desiredOrder := []string{
		"image/webp",
	}

	output := [][]generatedImage{}
	for _, dt := range desiredOrder {
		v, ok := sourceSetByType[dt]
		if !ok {
			continue
		}
		output = append(output, v)
		delete(sourceSetByType, dt)
	}

	for _, i := range sourceSetByType {
		output = append(output, i)
	}
	return output
}

type generatedImage struct {
	URL  string
	Type string
	Size int64
}
