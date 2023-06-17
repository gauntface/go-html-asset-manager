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

package preprocessors

import (
	"github.com/gauntface/go-html-asset-manager/v5/assets"
	"github.com/gauntface/go-html-asset-manager/v5/assets/assetmanager"
)

type Preprocessor func(runtime Runtime) error

type Runtime struct {
	Assets AssetManager
}

type AssetManager interface {
	All() []assetmanager.Asset
	StaticDir() string
	WithType(t assets.Type) []assetmanager.Asset
	AddRemote(a *assetmanager.RemoteAsset)
	AddLocal(a *assetmanager.LocalAsset)
}
