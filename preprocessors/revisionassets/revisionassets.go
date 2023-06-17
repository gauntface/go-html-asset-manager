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

package revisionassets

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/gauntface/go-html-asset-manager/v5/assets"
	"github.com/gauntface/go-html-asset-manager/v5/assets/assetmanager"
	"github.com/gauntface/go-html-asset-manager/v5/preprocessors"
	"github.com/gauntface/go-html-asset-manager/v5/utils/files"
)

var (
	errRenameFailed = errors.New("unable to rename file")

	assetsToRevision = []assets.Type{
		assets.SyncCSS,
		assets.AsyncCSS,
		assets.PreloadCSS,
		assets.SyncJS,
		assets.AsyncJS,
		assets.PreloadJS,
	}

	osRename = os.Rename
)

func Preprocessor(runtime preprocessors.Runtime) error {
	allAssets := runtime.Assets.All()
	for _, a := range allAssets {
		if !a.IsLocal() {
			continue
		}

		la := a.(*assetmanager.LocalAsset)
		if !shouldRevision(la) {
			continue
		}

		newPath, err := revisionFile(la.Path())
		if err != nil {
			return err
		}
		la.UpdatePath(newPath)
	}

	return nil
}

func shouldRevision(asset assetmanager.Asset) bool {
	for _, t := range assetsToRevision {
		if t == asset.Type() {
			return true
		}
	}
	return false
}

func revisionFile(filepath string) (string, error) {
	hash, err := files.Hash(filepath)
	if err != nil {
		return "", err
	}

	ext := path.Ext(filepath)
	newFilepath := fmt.Sprintf("%v.%v%v", filepath[0:len(filepath)-len(ext)], hash, ext)

	err = osRename(filepath, newFilepath)
	if err != nil {
		return "", fmt.Errorf("%w %q; %v", errRenameFailed, filepath, err)
	}
	return newFilepath, nil
}
