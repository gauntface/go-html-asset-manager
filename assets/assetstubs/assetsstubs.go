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

package assetstubs

import (
	"testing"

	"github.com/gauntface/go-html-asset-manager/v2/assets"
	"github.com/gauntface/go-html-asset-manager/v2/assets/assetmanager"
)

type Manager struct {
	AllReturn []assetmanager.Asset

	WithIDReturn map[string](map[assets.Type][]assetmanager.Asset)

	WithTypeReturn map[assets.Type][]assetmanager.Asset
}

func (m *Manager) All() []assetmanager.Asset {
	return m.AllReturn
}

func (m *Manager) WithType(t assets.Type) []assetmanager.Asset {
	return m.WithTypeReturn[t]
}

func (m *Manager) WithID(id string) map[assets.Type][]assetmanager.Asset {
	return m.WithIDReturn[id]
}

func (m *Manager) AddRemote(a *assetmanager.RemoteAsset) {}

func (m *Manager) String() string {
	return ""
}

type Asset struct {
	IDReturn string

	TypeReturn  assets.Type
	MediaReturn string

	URLReturn string

	ContentsReturn string
	ContentsError  error

	IsLocalReturn bool

	StringReturn string

	DebugReturn bool

	PathReturn string
}

func (a *Asset) ID() string {
	return a.IDReturn
}

func (a *Asset) Type() assets.Type {
	return a.TypeReturn
}

func (a *Asset) Media() string {
	return a.MediaReturn
}

func (a *Asset) URL() string {
	return a.URLReturn
}

func (a *Asset) Contents() (string, error) {
	return a.ContentsReturn, a.ContentsError
}

func (a *Asset) IsLocal() bool {
	return a.IsLocalReturn
}

func (a *Asset) String() string {
	return a.StringReturn
}

func (a *Asset) Debug(d string) bool {
	return a.DebugReturn
}

func (a *Asset) Path() string {
	return a.PathReturn
}

func MustNewLocalAsset(t *testing.T, dir, f string) *assetmanager.LocalAsset {
	t.Helper()

	a, err := assetmanager.NewLocalAsset(dir, f)
	if err != nil {
		t.Fatalf("Failed to make new local asset: %v", err)
	}
	return a
}
