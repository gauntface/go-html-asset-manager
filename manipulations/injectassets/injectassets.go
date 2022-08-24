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

package injectassets

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gauntface/go-html-asset-manager/v3/assets"
	"github.com/gauntface/go-html-asset-manager/v3/assets/assetmanager"
	"github.com/gauntface/go-html-asset-manager/v3/manipulations"
	"github.com/gauntface/go-html-asset-manager/v3/utils/html/htmlparsing"
	"github.com/gauntface/go-html-asset-manager/v3/utils/sets"
	"github.com/gauntface/go-html-asset-manager/v3/utils/stringui"
	"golang.org/x/net/html"
)

var (
	htmlparsingFindNodeByTag = htmlparsing.FindNodeByTag

	errElementNotFound = errors.New("html element not found")
)

func Manipulator(runtime manipulations.Runtime, doc *html.Node) error {
	keys := htmlparsing.GetKeys(doc)

	prettyPrintKeys(runtime.Debug, keys.Slice())

	headNode := htmlparsingFindNodeByTag("head", doc)
	if headNode == nil {
		return fmt.Errorf("%w: failed to find head element", errElementNotFound)
	}
	bodyNode := htmlparsingFindNodeByTag("body", doc)
	if bodyNode == nil {
		return fmt.Errorf("%w: failed to find body element", errElementNotFound)
	}

	inlineCSS := getAssetsForType(keys, runtime.Assets, assets.InlineCSS)
	err := addInlineCSS(headNode, inlineCSS)
	if err != nil {
		return err
	}

	injectMap := map[assets.Type]addAssetFunc{
		assets.SyncCSS:    addSyncCSS,
		assets.AsyncCSS:   addAsyncCSS,
		assets.PreloadCSS: addPreloadCSS,

		assets.InlineJS:  addInlineJS,
		assets.SyncJS:    addSyncJS,
		assets.AsyncJS:   addAsyncJS,
		assets.PreloadJS: addPreloadJS,
	}

	for _, k := range keys.Sorted() {
		assetsByType := runtime.Assets.WithID(k)
		assets := toArray(assetsByType)
		for _, as := range assets {
			injector, ok := injectMap[as.Type()]
			if ok {
				err := injector(headNode, bodyNode, as)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func toArray(assetsByType map[assets.Type][]assetmanager.Asset) []assetmanager.Asset {
	ar := []assetmanager.Asset{}
	for _, a := range assetsByType {
		ar = append(ar, a...)
	}

	sortAssets(ar)

	return ar
}

func getAssetsForType(keys sets.StringSet, manager manipulations.AssetManager, aType assets.Type) []assetmanager.Asset {
	assets := []assetmanager.Asset{}
	for _, k := range keys.Sorted() {
		assetsByType := manager.WithID(k)
		for ty, as := range assetsByType {
			if ty == aType {
				assets = append(assets, as...)
			}
		}
	}

	sortAssets(assets)

	return assets
}

func addInlineCSS(headNode *html.Node, assets []assetmanager.Asset) error {
	if len(assets) == 0 {
		return nil
	}

	contents := []string{}
	for _, a := range assets {
		c, err := a.Contents()
		if err != nil {
			return err
		}
		contents = append(contents, c)
	}
	headNode.AppendChild(htmlparsing.InlineCSSTag(strings.Join(contents, "\n\n")))
	return nil
}

func addSyncCSS(headNode, bodyNode *html.Node, asset assetmanager.Asset) error {
	u, err := asset.URL()
	if err != nil {
		return err
	}

	headNode.AppendChild(htmlparsing.SyncCSSTag(htmlparsing.CSSMediaPair{
		URL:   u,
		Media: asset.Media(),
	}))
	return nil
}

func addAsyncCSS(headNode, bodyNode *html.Node, asset assetmanager.Asset) error {
	u, err := asset.URL()
	if err != nil {
		return err
	}

	headNode.AppendChild(
		htmlparsing.AsyncCSSTag(
			htmlparsing.CSSMediaPair{
				URL:   u,
				Media: asset.Media(),
			},
		),
	)
	return nil
}

func addPreloadCSS(headNode, bodyNode *html.Node, asset assetmanager.Asset) error {
	u, err := asset.URL()
	if err != nil {
		return err
	}

	headNode.AppendChild(htmlparsing.PreloadTag("style", u))
	return nil
}

func addInlineJS(headNode, bodyNode *html.Node, asset assetmanager.Asset) error {
	c, err := asset.Contents()
	if err != nil {
		return err
	}
	bodyNode.AppendChild(htmlparsing.InlineJSTag(c))
	return nil
}

func addSyncJS(headNode, bodyNode *html.Node, asset assetmanager.Asset) error {
	u, err := asset.URL()
	if err != nil {
		return err
	}
	bodyNode.AppendChild(htmlparsing.SyncJSTag(u))
	return nil
}

func addAsyncJS(headNode, bodyNode *html.Node, asset assetmanager.Asset) error {
	u, err := asset.URL()
	if err != nil {
		return err
	}
	bodyNode.AppendChild(htmlparsing.AsyncJSTag(u))
	return nil
}

func addPreloadJS(headNode, bodyNode *html.Node, asset assetmanager.Asset) error {
	u, err := asset.URL()
	if err != nil {
		return err
	}
	headNode.AppendChild(htmlparsing.PreloadTag("script", u))
	return nil
}

func prettyPrintKeys(debug bool, keys []string) {
	if !debug {
		return
	}

	headings := []string{
		"Key",
	}

	rows := [][]string{}
	for _, key := range keys {
		rows = append(rows, []string{
			key,
		})
	}

	sort.Slice(rows, func(i, j int) bool {
		return rows[i][0] < rows[j][0]
	})

	fmt.Printf("HTML file keys\n")
	fmt.Println(stringui.Table(headings, rows))
}

func sortAssets(assets []assetmanager.Asset) {
	sort.Slice(assets, func(i, j int) bool {
		ui, _ := assets[i].URL()
		uj, _ := assets[j].URL()
		return ui < uj
	})
}

type addAssetFunc func(headNode, bodyNode *html.Node, asset assetmanager.Asset) error
