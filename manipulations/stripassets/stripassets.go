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

package stripassets

import (
	"fmt"
	"sort"

	"github.com/gauntface/go-html-asset-manager/manipulations"
	"github.com/gauntface/go-html-asset-manager/utils/html/htmlparsing"
	"github.com/gauntface/go-html-asset-manager/utils/stringui"
	"golang.org/x/net/html"
)

var (
	htmlparsingFindNodesByTag = htmlparsing.FindNodesByTag
)

func Manipulator(runtime manipulations.Runtime, doc *html.Node) error {
	keys := htmlparsing.GetKeys(doc)

	prettyPrintKeys(runtime.Debug, keys.Slice())

	linkNodes := htmlparsingFindNodesByTag("link", doc)
	for _, l := range linkNodes {
		l.Parent.RemoveChild(l)
	}

	scriptNodes := htmlparsingFindNodesByTag("script", doc)
	for _, s := range scriptNodes {
		s.Parent.RemoveChild(s)
	}

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
