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

// Package manipulations defines the interfaces all html doc manipulations must
// implement
package manipulations

import (
	"github.com/gauntface/go-html-asset-manager/assets"
	"github.com/gauntface/go-html-asset-manager/assets/assetmanager"
	"github.com/gauntface/go-html-asset-manager/utils/config"
	"github.com/gauntface/go-html-asset-manager/utils/vimeoapi"
	"golang.org/x/net/html"
)

type Manipulator func(runtime Runtime, doc *html.Node) error

type Runtime struct {
	Debug  bool
	Assets AssetManager
	Config *config.Config

	HasVimeo bool
	Vimeo    vimeoapiClient
}

type AssetManager interface {
	WithID(id string) map[assets.Type][]assetmanager.Asset
}

type vimeoapiClient interface {
	Video(videoID string) (*vimeoapi.Video, error)
}
