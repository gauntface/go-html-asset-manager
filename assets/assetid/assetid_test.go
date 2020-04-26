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

	"github.com/gauntface/go-html-asset-manager/assets"
	"github.com/google/go-cmp/cmp"
)

func TestTypeFromSyncSet(t *testing.T) {
	tests := []struct {
		description string
		filename    string
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
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := typeFromSyncSet(tt.filename, tt.inline, tt.sync, tt.async, tt.preload)
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
		wantExt      string
	}{
		{
			description:  "return filename and ext for full path",
			path:         "/example/nested/example.css",
			wantFilename: "example",
			wantExt:      ".css",
		},
		{
			description:  "return filename and ext for file only",
			path:         "example.css",
			wantFilename: "example",
			wantExt:      ".css",
		},
		{
			description:  "return filename and with empty ext",
			path:         "example",
			wantFilename: "example",
			wantExt:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got, gotExt := filename(tt.path)
			if got != tt.wantFilename {
				t.Errorf("Unexpected filename; got %v, want %v", got, tt.wantFilename)
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
			description: "return inline css with prefix",
			filename:    "example-inline",
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
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got, err := assetType(tt.filename, tt.ext)
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
		want        assets.Type
		wantError   error
	}{
		{
			description: "return error for unknown",
			path:        "example.unknown",
			want:        assets.Unknown,
			wantError:   ErrUnknownType,
		},
		{
			description: "return type for valid path",
			path:        "example.css",
			want:        assets.InlineCSS,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got, err := IdentifyType(tt.path)
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
