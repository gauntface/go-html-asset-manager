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

package files

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	filepathWalk = filepath.Walk
	osOpen       = os.Open
	ioCopy       = io.Copy

	errOpenFailed = errors.New("unable to open file")
	errHashFailed = errors.New("unable to generate hash for file")
)

func Find(dir string, exts ...string) ([]string, error) {
	found := []string{}
	err := filepathWalk(dir, func(absPath string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		if info.IsDir() {
			return nil
		}

		fe := filepath.Ext(info.Name())
		for _, ext := range exts {
			if strings.EqualFold(fe, ext) {
				found = append(found, absPath)
				return nil
			}
		}

		return nil
	})
	return found, err
}

func Hash(filepath string) (string, error) {
	f, err := osOpen(filepath)
	if err != nil {
		return "", fmt.Errorf("%w %q; %v", errOpenFailed, filepath, err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := ioCopy(h, f); err != nil {
		return "", fmt.Errorf("%w %q; %v", errHashFailed, filepath, err)
	}

	return hex.EncodeToString(h.Sum(nil))[0:7], nil
}
