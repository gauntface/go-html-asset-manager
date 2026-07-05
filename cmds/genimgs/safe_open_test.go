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

package main

import (
	"errors"
	"image"
	"strings"
	"testing"

	"github.com/disintegration/imaging"
)

func Test_safeOpenImage(t *testing.T) {
	origImagingOpen := imagingOpen
	defer func() { imagingOpen = origImagingOpen }()

	t.Run("returns the image on success", func(t *testing.T) {
		want := image.NewRGBA(image.Rect(0, 0, 1, 1))
		imagingOpen = func(path string, opts ...imaging.DecodeOption) (image.Image, error) {
			return want, nil
		}

		got, err := safeOpenImage("example.png")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if got != image.Image(want) {
			t.Fatalf("Unexpected image returned")
		}
	})

	t.Run("returns the underlying error without a panic", func(t *testing.T) {
		errInjected := errors.New("injected error")
		imagingOpen = func(path string, opts ...imaging.DecodeOption) (image.Image, error) {
			return nil, errInjected
		}

		_, err := safeOpenImage("example.png")
		if !errors.Is(err, errInjected) {
			t.Fatalf("Unexpected error; got %v, want %v", err, errInjected)
		}
	})

	t.Run("recovers a panic and returns it as an error instead", func(t *testing.T) {
		// This is the shape of failure disintegration/imaging has an
		// unpatched CVE for (CVE-2023-36308: panic on a crafted TIFF file),
		// so a single malformed image must not crash the whole batch run.
		imagingOpen = func(path string, opts ...imaging.DecodeOption) (image.Image, error) {
			panic("index out of range")
		}

		err := func() (err error) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("safeOpenImage should have recovered the panic itself, got: %v", r)
				}
			}()
			_, err = safeOpenImage("example.tiff")
			return err
		}()

		if err == nil {
			t.Fatalf("Expected an error, got nil")
		}
		if !strings.Contains(err.Error(), "example.tiff") {
			t.Fatalf("Expected error to mention the file path; got %v", err)
		}
	})
}
