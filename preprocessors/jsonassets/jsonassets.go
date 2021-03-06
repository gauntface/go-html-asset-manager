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

package jsonassets

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gauntface/go-html-asset-manager/assets"
	"github.com/gauntface/go-html-asset-manager/assets/assetmanager"
	"github.com/gauntface/go-html-asset-manager/preprocessors"
)

var (
	errJSONParseFailed = errors.New("unable to parse JSON file")
)

func Preprocessor(runtime preprocessors.Runtime) error {
	jsonAssets := runtime.Assets.WithType(assets.JSON)
	for _, a := range jsonAssets {
		remoteURLs, err := addJSONFile(a)
		if err != nil {
			return err
		}

		types := map[assets.Type][]string{
			assets.SyncCSS:    remoteURLs.CSS.Sync,
			assets.AsyncCSS:   remoteURLs.CSS.Async,
			assets.PreloadCSS: remoteURLs.CSS.Preload,

			assets.SyncJS:    remoteURLs.JS.Sync,
			assets.AsyncJS:   remoteURLs.JS.Async,
			assets.PreloadJS: remoteURLs.JS.Preload,
		}
		for t, urls := range types {
			for _, u := range urls {
				runtime.Assets.AddRemote(assetmanager.NewRemoteAsset(a.ID(), u, t))
			}
		}
	}
	return nil
}

func addJSONFile(asset assetmanager.Asset) (*jsonAssets, error) {
	c, err := asset.Contents()
	if err != nil {
		return nil, err
	}

	var content jsonAssets
	if err := json.Unmarshal([]byte(c), &content); err != nil {
		return nil, fmt.Errorf("%w %q; %v", errJSONParseFailed, asset, err)
	}

	return &content, nil
}

type jsonAssets struct {
	CSS jsonAssetGroup `json:"css,omitempty"`
	JS  jsonAssetGroup `json:"js,omitempty"`
}

type jsonAssetGroup struct {
	Sync    []string
	Async   []string
	Preload []string
}
