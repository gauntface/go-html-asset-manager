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
	"testing"

	"github.com/gauntface/go-html-asset-manager/v5/assets"
	"github.com/google/go-cmp/cmp"
)

func TestTypeFromSyncSet(t *testing.T) {
	tests := []struct {
		description string
		filename    string
		media       string
		inline      assets.Type
		sync        assets.Type
		async       assets.Type
		preload     assets.Type
		want        assets.Type
	}{
		{
			description: "return inline if no prefix matches",
			inline:      assets.InlineCSS,
			want:        assets.InlineCSS,
		},
		{
			description: "return inline if filename has the inline prefix",
			filename:    "example-inline",
			inline:      assets.InlineCSS,
			want:        assets.InlineCSS,
		},
		{
			description: "return sync if filename has the sync prefix",
			filename:    "example-sync",
			sync:        assets.SyncCSS,
			want:        assets.SyncCSS,
		},
		{
			description: "return async if filename has the async prefix",
			filename:    "example-async",
			async:       assets.AsyncCSS,
			want:        assets.AsyncCSS,
		},
		{
			description: "return preload if filename has the preload prefix",
			filename:    "example-preload",
			preload:     assets.PreloadCSS,
			want:        assets.PreloadCSS,
		},
		{
			description: "return last prefix only",
			filename:    "example-inline-sync-async",
			async:       assets.AsyncCSS,
			want:        assets.AsyncCSS,
		},
		{
			description: "return sync if filename has sync prefix before dot suffix",
			filename:    "example-sync.braille",
			sync:        assets.SyncCSS,
			want:        assets.SyncCSS,
		},
		{
			description: "return async if filename has async prefix before dot suffix",
			filename:    "example-async.screen",
			async:       assets.AsyncCSS,
			want:        assets.AsyncCSS,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := typeFromSyncSet(tt.filename, tt.media, tt.inline, tt.sync, tt.async, tt.preload)
			if got != tt.want {
				t.Errorf("Unexpected result; got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilename(t *testing.T) {
	tests := []struct {
		description  string
		path         string
		wantFilename string
		wantMedia    string
		wantExt      string
	}{
		{
			description:  "return filename and ext for full path",
			path:         "/example/nested/example.css",
			wantFilename: "example",
			wantMedia:    "",
			wantExt:      ".css",
		},
		{
			description:  "return filename and ext for file only",
			path:         "example.css",
			wantFilename: "example",
			wantMedia:    "",
			wantExt:      ".css",
		},
		{
			description:  "return filename and with empty ext",
			path:         "example",
			wantFilename: "example",
			wantMedia:    "",
			wantExt:      "",
		},
		{
			description:  "return filename, media and ext for simple file",
			path:         "example.print.css",
			wantFilename: "example",
			wantMedia:    "print",
			wantExt:      ".css",
		},
		{
			description:  "return filename, media and ext for sync file",
			path:         "example-sync.print.css",
			wantFilename: "example-sync",
			wantMedia:    "print",
			wantExt:      ".css",
		},
		{
			description:  "return filename, media and ext for complex file",
			path:         "example.screen and (max-width: 600px).css",
			wantFilename: "example",
			wantMedia:    "screen and (max-width: 600px)",
			wantExt:      ".css",
		},
		{
			description:  "return filename, media and ext for file with dots in name",
			path:         "example.test.dots.print.css",
			wantFilename: "example.test.dots",
			wantMedia:    "print",
			wantExt:      ".css",
		},
		{
			description:  "return filename and ext for file with dots in name but no media",
			path:         "example.test.print.dots.css",
			wantFilename: "example.test.print.dots",
			wantMedia:    "",
			wantExt:      ".css",
		},
		{
			description:  "return filename and ext for simple js file",
			path:         "example.print.js",
			wantFilename: "example.print",
			wantMedia:    "",
			wantExt:      ".js",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got, gotMedia, gotExt := filename(tt.path)
			if got != tt.wantFilename {
				t.Errorf("Unexpected filename; got %v, want %v", got, tt.wantFilename)
			}

			if gotMedia != tt.wantMedia {
				t.Errorf("Unexpected media; got %v, want %v", gotMedia, tt.wantMedia)
			}

			if gotExt != tt.wantExt {
				t.Errorf("Unexpected extension; got %v, want %v", gotExt, tt.wantExt)
			}
		})
	}
}

func TestAssetType(t *testing.T) {
	tests := []struct {
		description string
		filename    string
		media       string
		ext         string
		want        assets.Type
		wantError   error
	}{
		{
			description: "return error for unknown",
			filename:    "example",
			ext:         ".unknown",
			want:        assets.Unknown,
			wantError:   ErrUnknownType,
		},
		{
			description: "return inline css without prefix",
			filename:    "example",
			ext:         ".css",
			want:        assets.InlineCSS,
		},
		{
			description: "return sync css without prefix and media",
			filename:    "example",
			media:       "print",
			ext:         ".css",
			want:        assets.SyncCSS,
		},
		{
			description: "return inline css with prefix",
			filename:    "example-inline",
			ext:         ".css",
			want:        assets.InlineCSS,
		},
		{
			description: "return inline css with prefix and media",
			filename:    "example-inline",
			media:       "print",
			ext:         ".css",
			want:        assets.InlineCSS,
		},
		{
			description: "return sync css with prefix",
			filename:    "example-sync",
			ext:         ".css",
			want:        assets.SyncCSS,
		},
		{
			description: "return async css with prefix",
			filename:    "example-async",
			ext:         ".css",
			want:        assets.AsyncCSS,
		},
		{
			description: "return inline js without prefix",
			filename:    "example",
			ext:         ".js",
			want:        assets.InlineJS,
		},
		{
			description: "return inline js with prefix",
			filename:    "example-inline",
			ext:         ".js",
			want:        assets.InlineJS,
		},
		{
			description: "return sync js with prefix",
			filename:    "example-sync",
			ext:         ".js",
			want:        assets.SyncJS,
		},
		{
			description: "return async js with prefix",
			filename:    "example-async",
			ext:         ".js",
			want:        assets.AsyncJS,
		},
		{
			description: "return json",
			filename:    "example",
			ext:         ".json",
			want:        assets.JSON,
		},
		{
			description: "return html",
			filename:    "example",
			ext:         ".html",
			want:        assets.HTML,
		},
		{
			description: "return png",
			filename:    "example",
			ext:         ".png",
			want:        assets.PNG,
		},
		{
			description: "return jpeg for jpg",
			filename:    "example",
			ext:         ".jpg",
			want:        assets.JPEG,
		},
		{
			description: "return jpeg for jpeg",
			filename:    "example",
			ext:         ".jpeg",
			want:        assets.JPEG,
		},
		{
			description: "return webp for webp",
			filename:    "example",
			ext:         ".webp",
			want:        assets.WEBP,
		},
		{
			description: "return avif for avif",
			filename:    "example",
			ext:         ".avif",
			want:        assets.AVIF,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got, err := assetType(tt.filename, tt.media, tt.ext)
			if tt.wantError != nil {
				if !errors.Is(err, tt.wantError) {
					t.Fatalf("Different error returned; got %v, want %v", err, tt.wantError)
				}
			}
			if err != nil {
				return
			}

			if !cmp.Equal(got, tt.want) {
				t.Errorf("Unexpected result; Diff %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestIdentifyType(t *testing.T) {
	tests := []struct {
		description string
		path        string
		wantType    assets.Type
		wantMedia   string
		wantError   error
	}{
		{
			description: "return error for unknown",
			path:        "example.unknown",
			wantType:    assets.Unknown,
			wantMedia:   "",
			wantError:   ErrUnknownType,
		},
		{
			description: "return type for simple valid path",
			path:        "example.css",
			wantMedia:   "",
			wantType:    assets.InlineCSS,
		},
		{
			description: "return type for sync path with media",
			path:        "example-sync.print.css",
			wantMedia:   "print",
			wantType:    assets.SyncCSS,
		},
		{
			description: "return sync js for js with sync prefix and dot suffix",
			path:        "example-sync.braille.js",
			wantMedia:   "",
			wantType:    assets.SyncJS,
		},
		{
			description: "return async js for js with async prefix and dot suffix",
			path:        "example-async.screen (min-width: 200px).js",
			wantMedia:   "",
			wantType:    assets.AsyncJS,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			gotType, gotMedia, err := IdentifyType(tt.path)
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Different error returned; got %v, want %v", err, tt.wantError)
			}

			if !cmp.Equal(gotType, tt.wantType) {
				t.Errorf("Unexpected type; Diff %v", cmp.Diff(gotType, tt.wantType))
			}

			if !cmp.Equal(gotMedia, tt.wantMedia) {
				t.Errorf("Unexpected media; Diff %v", cmp.Diff(gotMedia, tt.wantMedia))
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	tests := []struct {
		description string
		path        string
		want        string
	}{
		{
			description: "return filename by default",
			path:        "example.unknown",
			want:        "example",
		},
		{
			description: "return filename if no prefix for css",
			path:        "example.css",
			want:        "example",
		},
		{
			description: "return filename if no prefix for js",
			path:        "example.js",
			want:        "example",
		},
		{
			description: "return filename without inline prefix for css",
			path:        "example-inline.css",
			want:        "example",
		},
		{
			description: "return filename without inline prefix for js",
			path:        "example-inline.js",
			want:        "example",
		},
		{
			description: "return filename without sync prefix for css",
			path:        "example-sync.css",
			want:        "example",
		},
		{
			description: "return filename without sync prefix for js",
			path:        "example-sync.js",
			want:        "example",
		},
		{
			description: "return filename without async prefix for css",
			path:        "example-async.css",
			want:        "example",
		},
		{
			description: "return filename without async prefix for js",
			path:        "example-async.js",
			want:        "example",
		},
		{
			description: "return filename without preload prefix for css",
			path:        "example-preload.css",
			want:        "example",
		},
		{
			description: "return filename without preload prefix for js",
			path:        "example-preload.js",
			want:        "example",
		},
		{
			description: "return filename without last prefix only",
			path:        "example-inline-sync-async.js",
			want:        "example-inline-sync",
		},
		{
			description: "return filename without sync prefix for js with dot suffix",
			path:        "example-sync.braille.js",
			want:        "example.braille",
		},
		{
			description: "return filename without async prefix for js with dot suffix",
			path:        "example-async.screen (min-width: 200px).js",
			want:        "example.screen (min-width: 200px)",
		},
		{
			description: "return filename with dot suffix for js without prefix",
			path:        "example.print.js",
			want:        "example.print",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := Generate(tt.path)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Unexpected result; Diff %v", cmp.Diff(got, tt.want))
			}
		})
	}
}
