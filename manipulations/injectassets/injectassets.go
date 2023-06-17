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

	"github.com/gauntface/go-html-asset-manager/v5/assets"
	"github.com/gauntface/go-html-asset-manager/v5/assets/assetmanager"
	"github.com/gauntface/go-html-asset-manager/v5/manipulations"
	"github.com/gauntface/go-html-asset-manager/v5/utils/html/htmlparsing"
	"github.com/gauntface/go-html-asset-manager/v5/utils/stringui"
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

	injectMap := map[assets.Type]addAssetFunc{
		assets.InlineCSS:  addInlineCSS,
		assets.SyncCSS:    addSyncCSS,
		assets.AsyncCSS:   addAsyncCSS,
		assets.PreloadCSS: addPreloadCSS,

		assets.InlineJS:  addInlineJS,
		assets.SyncJS:    addSyncJS,
		assets.AsyncJS:   addAsyncJS,
		assets.PreloadJS: addPreloadJS,
	}

	assetOrder := []assets.Type{
		assets.InlineCSS,
		assets.InlineJS,

		assets.PreloadCSS,
		assets.PreloadJS,

		assets.AsyncCSS,
		assets.AsyncJS,

		assets.SyncCSS,
		assets.SyncJS,
	}

	sortedKeys := keys.Sorted()

	for _, a := range assetOrder {
		for _, k := range sortedKeys {
			assetsByType := runtime.Assets.WithID(k)
			assets := toArray(assetsByType)
			for _, as := range assets {
				if as.Type() != a {
					continue
				}

				injector, ok := injectMap[as.Type()]
				if ok {
					err := injector(headNode, bodyNode, as)
					if err != nil {
						return err
					}
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

func addInlineCSS(headNode, bodyNode *html.Node, asset assetmanager.Asset) error {
	c, err := asset.Contents()
	if err != nil {
		return err
	}
	style := htmlparsing.FindNodeByTag("style", headNode)
	if style == nil {
		headNode.AppendChild(htmlparsing.InlineCSSTag(c))
	} else {
		style.FirstChild.Data += " " + c
	}

	return nil
}

func addSyncCSS(headNode, bodyNode *html.Node, asset assetmanager.Asset) error {
	u, err := asset.URL()
	if err != nil {
		return err
	}

	node := headNode
	if asset.Media() != "" {
		node = bodyNode
	}

	node.AppendChild(htmlparsing.SyncCSSTag(htmlparsing.CSSTagData{
		URL:        u,
		Attributes: asset.Attributes(),
		Media:      asset.Media(),
	}))

	return nil
}

func addAsyncCSS(headNode, bodyNode *html.Node, asset assetmanager.Asset) error {
	u, err := asset.URL()
	if err != nil {
		return err
	}

	bodyNode.AppendChild(
		htmlparsing.AsyncCSSTag(
			htmlparsing.CSSTagData{
				URL:        u,
				Attributes: asset.Attributes(),
				Media:      asset.Media(),
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

	bodyNode.AppendChild(htmlparsing.PreloadTag("style", u))
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
	bodyNode.AppendChild(htmlparsing.SyncJSTag(htmlparsing.JSTagData{
		URL:        u,
		Attributes: asset.Attributes(),
	}))
	return nil
}

func addAsyncJS(headNode, bodyNode *html.Node, asset assetmanager.Asset) error {
	u, err := asset.URL()
	if err != nil {
		return err
	}
	bodyNode.AppendChild(htmlparsing.AsyncJSTag(htmlparsing.JSTagData{
		URL:        u,
		Attributes: asset.Attributes(),
	}))
	return nil
}

func addPreloadJS(headNode, bodyNode *html.Node, asset assetmanager.Asset) error {
	u, err := asset.URL()
	if err != nil {
		return err
	}
	bodyNode.AppendChild(htmlparsing.PreloadTag("script", u))
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
