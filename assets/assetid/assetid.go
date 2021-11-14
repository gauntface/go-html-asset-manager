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

package assetid

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gauntface/go-html-asset-manager/assets"
)

const (
	inlinePrefix  = "-inline"
	syncPrefix    = "-sync"
	asyncPrefix   = "-async"
	preloadPrefix = "-preload"
)

var (
	prefixesToTrim = []string{
		inlinePrefix,
		syncPrefix,
		asyncPrefix,
		preloadPrefix,
	}

	mediaValues = []string{
		"print",
		"screen",
		"aural",
		"braille",
	}

	ErrUnknownType = errors.New("unknown asset type")
)

func Generate(path string) string {
	// Initialise ID to just the filename
	id, _, ext := filename(path)

	switch ext {
	case ".css", ".js":
		for _, pr := range prefixesToTrim {
			if strings.HasSuffix(id, pr) {
				id = strings.TrimSuffix(id, pr)
				break
			}
		}
	}

	return id
}

func IdentifyType(path string) (assets.Type, string, error) {
	fn, media, ext := filename(path)
	t, err := assetType(fn, ext)
	return t, media, err
}

func assetType(fn, ext string) (assets.Type, error) {
	switch ext {
	case ".css":
		return typeFromSyncSet(fn, assets.InlineCSS, assets.SyncCSS, assets.AsyncCSS, assets.PreloadCSS), nil
	case ".js":
		return typeFromSyncSet(fn, assets.InlineJS, assets.SyncJS, assets.AsyncJS, assets.PreloadJS), nil
	case ".json":
		return assets.JSON, nil
	case ".html":
		return assets.HTML, nil
	case ".png":
		return assets.PNG, nil
	case ".jpg", ".jpeg":
		return assets.JPEG, nil
	case ".webp":
		return assets.WEBP, nil
	}
	return assets.Unknown, fmt.Errorf("%w: for file %q with extension %q", ErrUnknownType, fn, ext)
}

func typeFromSyncSet(fn string, inline, sync, async, preload assets.Type) assets.Type {
	t := inline
	prefixes := map[string]assets.Type{
		inlinePrefix:  inline,
		syncPrefix:    sync,
		asyncPrefix:   async,
		preloadPrefix: preload,
	}
	for pr, ty := range prefixes {
		if strings.HasSuffix(fn, pr) {
			t = ty
			break
		}
	}
	return t
}

func filename(path string) (filename string, media string, ext string) {
	ext = filepath.Ext(path)
	media = ""
	filename = strings.TrimSuffix(filepath.Base(path), ext)
	pieces := strings.Split(filename, ".")
	if len(pieces) == 1 || ext != ".css" || !isMedia(pieces[len(pieces)-1]) {
		return filename, media, ext
	}

	media = pieces[len(pieces)-1]
	pieces = pieces[:len(pieces)-1]
	return strings.Join(pieces, "."), media, ext
}

func isMedia(s string) bool {
	for _, m := range mediaValues {
		if strings.Contains(strings.ToLower(s), m) {
			return true
		}
	}
	return false
}
