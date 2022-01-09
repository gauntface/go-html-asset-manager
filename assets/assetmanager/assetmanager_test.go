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

package assetmanager

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/gauntface/go-html-asset-manager/v2/assets"
	"github.com/gauntface/go-html-asset-manager/v2/assets/assetid"
	"github.com/google/go-cmp/cmp"
)

var errInjected = errors.New("injected error")

var reset func()

func TestMain(m *testing.M) {
	origFilesFind := filesFind

	reset = func() {
		filesFind = origFilesFind
	}

	os.Exit(m.Run())
}

func TestNewManager(t *testing.T) {
	tests := []struct {
		description string
		htmlDir     string
		staticDir   string
		jsonDir     string
		filesFind   func(dir string, exts ...string) ([]string, error)
		want        *Manager
		wantError   error
	}{
		{
			description: "do nothing if dirs are not set",
			want: &Manager{
				remoteAssets: []*RemoteAsset{},
				localAssets:  []*LocalAsset{},
			},
		},
		{
			description: "return error if retrieving HTML fails",
			htmlDir:     "/html/",
			filesFind: func(dir string, exts ...string) ([]string, error) {
				wantDir := "/html/"
				if dir != wantDir {
					t.Fatalf("Unexpected dir for files.Find; got %v, want %v", dir, wantDir)
				}
				wantExts := []string{".html"}
				if diff := cmp.Diff(exts, wantExts); diff != "" {
					t.Fatalf("Unexpected exts for files.Find; diff %v", diff)
				}
				return nil, errInjected
			},
			wantError: errInjected,
		},
		{
			description: "return error if retrieving assets fails",
			staticDir:   "/assets/",
			filesFind: func(dir string, exts ...string) ([]string, error) {
				wantDir := "/assets/"
				if dir != wantDir {
					t.Fatalf("Unexpected dir for files.Find; got %v, want %v", dir, wantDir)
				}
				wantExts := []string{".css", ".js", ".png", ".jpg", ".jpeg", ".webp", ".avif"}
				if diff := cmp.Diff(exts, wantExts); diff != "" {
					t.Fatalf("Unexpected exts for files.Find; diff %v", diff)
				}
				return nil, errInjected
			},
			wantError: errInjected,
		},
		{
			description: "return error if retrieving json fails",
			staticDir:   "",
			jsonDir:     "/json/",
			filesFind: func(dir string, exts ...string) ([]string, error) {
				wantDir := "/json/"
				if dir != wantDir {
					t.Fatalf("Unexpected dir for files.Find; got %v, want %v", dir, wantDir)
				}
				wantExts := []string{".json"}
				if diff := cmp.Diff(exts, wantExts); diff != "" {
					t.Fatalf("Unexpected exts for files.Find; diff %v", diff)
				}
				return nil, errInjected
			},
			wantError: errInjected,
		},
		{
			description: "return manager with assets no media",
			staticDir:   "/assets/",
			jsonDir:     "/json/",
			filesFind: func(dir string, exts ...string) ([]string, error) {
				switch dir {
				case "/assets/":
					return []string{
						"/assets/example.css",
						"/assets/example-inline.css",
						"/assets/example-sync.css",
						"/assets/example-async.css",
						"/assets/example.js",
						"/assets/example-inline.js",
						"/assets/example-sync.js",
						"/assets/example-async.js",
					}, nil
				case "/json/":
					return []string{
						"/json/example.json",
						"/json/example-inline.json",
					}, nil
				}
				return nil, fmt.Errorf("unknown dir passed to files.Find(): %q", dir)
			},
			want: &Manager{
				remoteAssets: []*RemoteAsset{},
				localAssets: []*LocalAsset{
					// CSS
					{
						assetType:    assets.InlineCSS,
						id:           "example",
						originalPath: "/assets/example.css",
						path:         "/assets/example.css",
						relativeDir:  "/assets/",
						readFile:     ioutil.ReadFile,
					},
					{
						assetType:    assets.InlineCSS,
						id:           "example",
						originalPath: "/assets/example-inline.css",
						path:         "/assets/example-inline.css",
						relativeDir:  "/assets/",
						readFile:     ioutil.ReadFile,
					},
					{
						assetType:    assets.SyncCSS,
						id:           "example",
						originalPath: "/assets/example-sync.css",
						path:         "/assets/example-sync.css",
						relativeDir:  "/assets/",
						readFile:     ioutil.ReadFile,
					},
					{
						assetType:    assets.AsyncCSS,
						id:           "example",
						originalPath: "/assets/example-async.css",
						path:         "/assets/example-async.css",
						relativeDir:  "/assets/",
						readFile:     ioutil.ReadFile,
					},

					// JS
					{
						assetType:    assets.InlineJS,
						id:           "example",
						originalPath: "/assets/example.js",
						path:         "/assets/example.js",
						relativeDir:  "/assets/",
						readFile:     ioutil.ReadFile,
					},
					{
						assetType:    assets.InlineJS,
						id:           "example",
						originalPath: "/assets/example-inline.js",
						path:         "/assets/example-inline.js",
						relativeDir:  "/assets/",
						readFile:     ioutil.ReadFile,
					},
					{
						assetType:    assets.SyncJS,
						id:           "example",
						originalPath: "/assets/example-sync.js",
						path:         "/assets/example-sync.js",
						relativeDir:  "/assets/",
						readFile:     ioutil.ReadFile,
					},
					{
						assetType:    assets.AsyncJS,
						id:           "example",
						originalPath: "/assets/example-async.js",
						path:         "/assets/example-async.js",
						relativeDir:  "/assets/",
						readFile:     ioutil.ReadFile,
					},
					{
						assetType:    assets.JSON,
						id:           "example",
						originalPath: "/json/example.json",
						path:         "/json/example.json",
						relativeDir:  "/json/",
						readFile:     ioutil.ReadFile,
					},
					{
						assetType:    assets.JSON,
						id:           "example-inline",
						originalPath: "/json/example-inline.json",
						path:         "/json/example-inline.json",
						relativeDir:  "/json/",
						readFile:     ioutil.ReadFile,
					},
				},
			},
		},
		{
			description: "return manager with assets and media",
			staticDir:   "/assets/",
			jsonDir:     "/json/",
			filesFind: func(dir string, exts ...string) ([]string, error) {
				switch dir {
				case "/assets/":
					return []string{
						"/assets/example.print.css",
						"/assets/example-inline.screen.css",
						"/assets/example-sync.braille.css",
						"/assets/example-async.screen (min-width: 200px).css",
						"/assets/example.print.js",
						"/assets/example-inline.screen.js",
						"/assets/example-sync.braille.js",
						"/assets/example-async.screen (min-width: 200px).js",
					}, nil
				case "/json/":
					return []string{
						"/json/example.print.json",
						"/json/example-inline.screen (min-width: 200px).json",
					}, nil
				}
				return nil, fmt.Errorf("unknown dir passed to files.Find(): %q", dir)
			},
			want: &Manager{
				remoteAssets: []*RemoteAsset{},
				localAssets: []*LocalAsset{
					// CSS
					{
						assetType:    assets.InlineCSS,
						assetMedia:   "print",
						id:           "example",
						originalPath: "/assets/example.print.css",
						path:         "/assets/example.print.css",
						relativeDir:  "/assets/",
						readFile:     ioutil.ReadFile,
					},
					{
						assetType:    assets.InlineCSS,
						assetMedia:   "screen",
						id:           "example",
						originalPath: "/assets/example-inline.screen.css",
						path:         "/assets/example-inline.screen.css",
						relativeDir:  "/assets/",
						readFile:     ioutil.ReadFile,
					},
					{
						assetType:    assets.SyncCSS,
						assetMedia:   "braille",
						id:           "example",
						originalPath: "/assets/example-sync.braille.css",
						path:         "/assets/example-sync.braille.css",
						relativeDir:  "/assets/",
						readFile:     ioutil.ReadFile,
					},
					{
						assetType:    assets.AsyncCSS,
						assetMedia:   "screen (min-width: 200px)",
						id:           "example",
						originalPath: "/assets/example-async.screen (min-width: 200px).css",
						path:         "/assets/example-async.screen (min-width: 200px).css",
						relativeDir:  "/assets/",
						readFile:     ioutil.ReadFile,
					},

					// JS
					{
						assetType:    assets.InlineJS,
						id:           "example.print",
						originalPath: "/assets/example.print.js",
						path:         "/assets/example.print.js",
						relativeDir:  "/assets/",
						readFile:     ioutil.ReadFile,
					},
					{
						assetType:    assets.InlineJS,
						id:           "example-inline.screen",
						originalPath: "/assets/example-inline.screen.js",
						path:         "/assets/example-inline.screen.js",
						relativeDir:  "/assets/",
						readFile:     ioutil.ReadFile,
					},
					{
						assetType:    assets.InlineJS,
						id:           "example-sync.braille",
						originalPath: "/assets/example-sync.braille.js",
						path:         "/assets/example-sync.braille.js",
						relativeDir:  "/assets/",
						readFile:     ioutil.ReadFile,
					},
					{
						assetType:    assets.InlineJS,
						id:           "example-async.screen (min-width: 200px)",
						originalPath: "/assets/example-async.screen (min-width: 200px).js",
						path:         "/assets/example-async.screen (min-width: 200px).js",
						relativeDir:  "/assets/",
						readFile:     ioutil.ReadFile,
					},
					{
						assetType:    assets.JSON,
						id:           "example.print",
						originalPath: "/json/example.print.json",
						path:         "/json/example.print.json",
						relativeDir:  "/json/",
						readFile:     ioutil.ReadFile,
					},
					{
						assetType:    assets.JSON,
						id:           "example-inline.screen (min-width: 200px)",
						originalPath: "/json/example-inline.screen (min-width: 200px).json",
						path:         "/json/example-inline.screen (min-width: 200px).json",
						relativeDir:  "/json/",
						readFile:     ioutil.ReadFile,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			defer reset()

			filesFind = tt.filesFind

			got, err := NewManager(tt.htmlDir, tt.staticDir, tt.jsonDir)
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			if err != nil {
				return
			}

			opts := []cmp.Option{
				cmp.AllowUnexported(Manager{}, LocalAsset{}),
				cmp.Comparer(func(x, y func(string) ([]byte, error)) bool {
					return reflect.ValueOf(x).Pointer() == reflect.ValueOf(y).Pointer()
				}),
			}
			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Errorf("Unexpected result; Diff %v", diff)
			}
		})
	}
}

func Test_findLocalAssets(t *testing.T) {
	tests := []struct {
		description string
		dir         string
		exts        []string
		filesFind   func(dir string, exts ...string) ([]string, error)
		want        []*LocalAsset
		wantError   error
	}{
		{
			description: "do nothing if dir is empty",
		},
		{
			description: "return error if files.Find fails",
			dir:         "/example/",
			filesFind: func(dir string, exts ...string) ([]string, error) {
				return nil, errInjected
			},
			wantError: errInjected,
		},
		{
			description: "return error if local asset files",
			dir:         "/example/",
			filesFind: func(dir string, exts ...string) ([]string, error) {
				return []string{"/example/example.unknown-ext"}, nil
			},
			wantError: assetid.ErrUnknownType,
		},
		{
			description: "return assets",
			dir:         "/example/",
			filesFind: func(dir string, exts ...string) ([]string, error) {
				return []string{
					"/example/example.css",
					"/example/example-inline.css",
					"/example/example-sync.css",
					"/example/example-async.css",
				}, nil
			},
			want: []*LocalAsset{
				{
					assetType:    assets.InlineCSS,
					id:           "example",
					originalPath: "/example/example.css",
					path:         "/example/example.css",
					relativeDir:  "/example/",
					readFile:     ioutil.ReadFile,
				},
				{
					assetType:    assets.InlineCSS,
					id:           "example",
					originalPath: "/example/example-inline.css",
					path:         "/example/example-inline.css",
					relativeDir:  "/example/",
					readFile:     ioutil.ReadFile,
				},
				{
					assetType:    assets.SyncCSS,
					id:           "example",
					originalPath: "/example/example-sync.css",
					path:         "/example/example-sync.css",
					relativeDir:  "/example/",
					readFile:     ioutil.ReadFile,
				},
				{
					assetType:    assets.AsyncCSS,
					id:           "example",
					originalPath: "/example/example-async.css",
					path:         "/example/example-async.css",
					relativeDir:  "/example/",
					readFile:     ioutil.ReadFile,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			defer reset()

			filesFind = tt.filesFind

			got, err := findLocalAssets(tt.dir, tt.exts...)
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			if err != nil {
				return
			}

			opts := []cmp.Option{
				cmp.AllowUnexported(LocalAsset{}),
				cmp.Comparer(func(x, y func(string) ([]byte, error)) bool {
					return reflect.ValueOf(x).Pointer() == reflect.ValueOf(y).Pointer()
				}),
			}
			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Errorf("Unexpected result; Diff %v", diff)
			}
		})
	}
}

func TestNewLocalAsset(t *testing.T) {
	tests := []struct {
		description string
		relDir      string
		path        string
		want        *LocalAsset
		wantError   error
	}{
		{
			description: "return error if type cannot be identified",
			path:        "/example.unknown-ext",
			wantError:   assetid.ErrUnknownType,
		},
		{
			description: "return local asset",
			path:        "/example/example.css",
			relDir:      "/example/",
			want: &LocalAsset{
				assetType:    assets.InlineCSS,
				id:           "example",
				originalPath: "/example/example.css",
				path:         "/example/example.css",
				relativeDir:  "/example/",

				readFile: ioutil.ReadFile,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got, err := NewLocalAsset(tt.relDir, tt.path)
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			if err != nil {
				return
			}

			opts := []cmp.Option{
				cmp.AllowUnexported(LocalAsset{}),
				cmp.Comparer(func(x, y func(string) ([]byte, error)) bool {
					return reflect.ValueOf(x).Pointer() == reflect.ValueOf(y).Pointer()
				}),
			}
			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Errorf("Unexpected result; Diff %v", diff)
			}
		})
	}
}

func TestManager_All(t *testing.T) {
	tests := []struct {
		description string
		manager     *Manager
		want        []Asset
		wantError   error
	}{
		{
			description: "return all assets",
			manager: &Manager{
				localAssets: []*LocalAsset{
					{
						id: "1",
					},
					{
						id: "2",
					},
				},
				remoteAssets: []*RemoteAsset{
					{
						id: "3",
					},
					{
						id: "4",
					},
				},
			},
			want: []Asset{
				&LocalAsset{id: "1"},
				&LocalAsset{id: "2"},
				&RemoteAsset{id: "3"},
				&RemoteAsset{id: "4"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := tt.manager.All()
			opts := []cmp.Option{
				cmp.AllowUnexported(LocalAsset{}, RemoteAsset{}),
			}
			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Errorf("Unexpected result; Diff %v", diff)
			}
		})
	}
}

func TestManager_WithType(t *testing.T) {
	tests := []struct {
		description string
		assetType   assets.Type
		manager     *Manager
		want        []Asset
	}{
		{
			description: "return assets with a particular type",
			assetType:   assets.HTML,
			manager: &Manager{
				localAssets: []*LocalAsset{
					{
						id: "1",
					},
					{
						id:        "2",
						assetType: assets.HTML,
					},
				},
				remoteAssets: []*RemoteAsset{
					{
						id:        "3",
						assetType: assets.PNG,
					},
					{
						id:        "4",
						assetType: assets.HTML,
					},
				},
			},
			want: []Asset{
				&LocalAsset{id: "2", assetType: assets.HTML},
				&RemoteAsset{id: "4", assetType: assets.HTML},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := tt.manager.WithType(tt.assetType)
			opts := []cmp.Option{
				cmp.AllowUnexported(LocalAsset{}, RemoteAsset{}),
			}
			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Errorf("Unexpected result; Diff %v", diff)
			}
		})
	}
}

func TestManager_WithID(t *testing.T) {
	tests := []struct {
		description string
		ID          string
		manager     *Manager
		want        map[assets.Type][]Asset
	}{
		{
			description: "return assets with a particular type",
			ID:          "1",
			manager: &Manager{
				localAssets: []*LocalAsset{
					{
						id:        "1",
						assetType: assets.HTML,
					},
					{
						id: "2",
					},
				},
				remoteAssets: []*RemoteAsset{
					{
						id:        "1",
						assetType: assets.HTML,
					},
					{
						id:        "1",
						assetType: assets.PNG,
					},
					{
						id: "3",
					},
				},
			},
			want: map[assets.Type][]Asset{
				assets.HTML: {
					&LocalAsset{id: "1", assetType: assets.HTML},
					&RemoteAsset{id: "1", assetType: assets.HTML},
				},
				assets.PNG: {
					&RemoteAsset{id: "1", assetType: assets.PNG},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := tt.manager.WithID(tt.ID)
			opts := []cmp.Option{
				cmp.AllowUnexported(LocalAsset{}, RemoteAsset{}),
			}
			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Errorf("Unexpected result; Diff %v", diff)
			}
		})
	}
}

func TestManager_AddRemote(t *testing.T) {
	tests := []struct {
		description string
		asset       *RemoteAsset
		manager     *Manager
		want        []*RemoteAsset
	}{
		{
			description: "Add remote asset to manager",
			manager:     &Manager{},
			asset:       &RemoteAsset{id: "1"},
			want: []*RemoteAsset{
				{
					id: "1",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			tt.manager.AddRemote(tt.asset)
			opts := []cmp.Option{
				cmp.AllowUnexported(RemoteAsset{}),
			}
			if diff := cmp.Diff(tt.manager.remoteAssets, tt.want, opts...); diff != "" {
				t.Errorf("Unexpected result; Diff %v", diff)
			}
		})
	}
}

func TestManager_String(t *testing.T) {
	tests := []struct {
		description string
		manager     *Manager
		want        string
	}{
		{
			description: "return a list of assets",
			manager: &Manager{
				localAssets: []*LocalAsset{
					// No HTML asset on purpose

					// Two PNG assets on purpose
					{
						id:        "png-a",
						assetType: assets.PNG,
					},
					{
						id:        "png-b",
						assetType: assets.PNG,
					},

					{
						id:        "inline css-1",
						assetType: assets.InlineCSS,
					},
					{
						id:        "inline css-2",
						assetType: assets.InlineCSS,
					},
					{
						id:        "async css",
						assetType: assets.AsyncCSS,
					},
					{
						id:        "sync js",
						assetType: assets.SyncJS,
					},
					{
						id:        "preload js",
						assetType: assets.PreloadJS,
					},
				},
				remoteAssets: []*RemoteAsset{
					{
						id:        "json",
						assetType: assets.JSON,
					},
					{
						id:        "jpeg",
						assetType: assets.JPEG,
					},
					{
						id:        "sync css",
						assetType: assets.SyncCSS,
					},
					{
						id:        "inline js",
						assetType: assets.InlineJS,
					},
					{
						id:        "async js",
						assetType: assets.AsyncJS,
					},
				},
			},
			want: `JSON (Total: 1)

PNG (Total: 2)

JPEG (Total: 1)

Inline CSS (Total: 2)
------------------------
| Key          | Count |
------------------------
| inline css-1 | 1     |
| inline css-2 | 1     |
------------------------

Sync CSS (Total: 1)
--------------------
| Key      | Count |
--------------------
| sync css | 1     |
--------------------

Async CSS (Total: 1)
---------------------
| Key       | Count |
---------------------
| async css | 1     |
---------------------

Inline JS (Total: 1)
---------------------
| Key       | Count |
---------------------
| inline js | 1     |
---------------------

Sync JS (Total: 1)
-------------------
| Key     | Count |
-------------------
| sync js | 1     |
-------------------

Async JS (Total: 1)
--------------------
| Key      | Count |
--------------------
| async js | 1     |
--------------------

Preload JS (Total: 1)
----------------------
| Key        | Count |
----------------------
| preload js | 1     |
----------------------`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := tt.manager.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Unexpected result; Diff %v", diff)
			}
		})
	}
}

func TestLocalAsset_Path(t *testing.T) {
	tests := []struct {
		description string
		asset       *LocalAsset
		want        string
	}{
		{
			description: "return path",
			asset: &LocalAsset{
				path: "/example/example.html",
			},
			want: "/example/example.html",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := tt.asset.Path()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Unexpected result; Diff %v", diff)
			}
		})
	}
}

func TestLocalAsset_UpdatePath(t *testing.T) {
	tests := []struct {
		description string
		asset       *LocalAsset
		path        string
		want        string
	}{
		{
			description: "return path",
			asset: &LocalAsset{
				path: "/example/original.html",
			},
			path: "/example/updated.html",
			want: "/example/updated.html",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			tt.asset.UpdatePath(tt.path)
			if diff := cmp.Diff(tt.asset.path, tt.want); diff != "" {
				t.Errorf("Unexpected result; Diff %v", diff)
			}
		})
	}
}

func TestLocalAsset_Contents(t *testing.T) {
	tests := []struct {
		description string
		asset       *LocalAsset
		want        string
		wantError   error
	}{
		{
			description: "return error if read fails",
			asset: &LocalAsset{
				path: "/example/original.html",
				readFile: func(filename string) ([]byte, error) {
					return nil, errInjected
				},
			},
			wantError: errReadFailed,
		},
		{
			description: "return file contents",
			asset: &LocalAsset{
				path: "/example/original.html",
				readFile: func(filename string) ([]byte, error) {
					return []byte("Hello world"), nil
				},
			},
			want: "Hello world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got, err := tt.asset.Contents()
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Unexpected result; Diff %v", diff)
			}
		})
	}
}

func TestLocalAsset_URL(t *testing.T) {
	tests := []struct {
		description string
		asset       *LocalAsset
		want        string
		wantError   error
	}{
		{
			description: "return error if read fails",
			asset: &LocalAsset{
				relativeDir: "..",
				path:        "/example/original.html",
			},
			wantError: errRelPath,
		},
		{
			description: "return file contents",
			asset: &LocalAsset{
				relativeDir: "/example/",
				path:        "/example/original.html",
			},
			want: "/original.html",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got, err := tt.asset.URL()
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Unexpected result; Diff %v", diff)
			}
		})
	}
}

func TestLocalAsset_IsLocal(t *testing.T) {
	a := &LocalAsset{}
	got := a.IsLocal()
	want := true
	if got != want {
		t.Errorf("Unexpected result; got %v, want %v", got, want)
	}
}

func TestLocalAsset_Debug(t *testing.T) {
	tests := []struct {
		description string
		d           string
		asset       *LocalAsset
		want        bool
	}{
		{
			description: "return false for different files",
			d:           "other-path.html",
			asset: &LocalAsset{
				originalPath: "/example/original.html",
			},
			want: false,
		},
		{
			description: "return true for same files",
			d:           "/example/original.html",
			asset: &LocalAsset{
				originalPath: "/example/original.html",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := tt.asset.Debug(tt.d)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Unexpected result; Diff %v", diff)
			}
		})
	}
}

func TestLocalAsset_String(t *testing.T) {
	tests := []struct {
		description string
		asset       *LocalAsset
		want        string
	}{
		{
			description: "return string value for same paths",
			asset: &LocalAsset{
				originalPath: "/example/original.html",
				path:         "/example/original.html",
			},
			want: `<Local Asset: "/example/original.html">`,
		},
		{
			description: "return string value for different paths",
			asset: &LocalAsset{
				originalPath: "/example/original.html",
				path:         "/example/new.html",
			},
			want: `<Local Asset: "/example/original.html" | "/example/new.html">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := tt.asset.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Unexpected result; Diff %v", diff)
			}
		})
	}
}

func TestNewRemoteAsset(t *testing.T) {
	tests := []struct {
		description string
		id          string
		url         string
		assetType   assets.Type
		want        *RemoteAsset
	}{
		{
			description: "return new remote asset",
			id:          "example-id",
			url:         "http://example.com/example.css",
			assetType:   assets.InlineCSS,
			want: &RemoteAsset{
				id:        "example-id",
				url:       "http://example.com/example.css",
				assetType: assets.InlineCSS,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := NewRemoteAsset(tt.id, tt.url, tt.assetType)
			opts := []cmp.Option{
				cmp.AllowUnexported(RemoteAsset{}),
			}
			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Errorf("Unexpected result; Diff %v", diff)
			}
		})
	}
}

func TestRemoteAsset_IsLocal(t *testing.T) {
	a := &RemoteAsset{}
	got := a.IsLocal()
	want := false
	if got != want {
		t.Errorf("Unexpected result; got %v, want %v", got, want)
	}
}

func TestRemoteAsset_Debug(t *testing.T) {
	tests := []struct {
		description string
		d           string
		asset       *RemoteAsset
		want        bool
	}{
		{
			description: "return false for different files",
			d:           "http://example.com/other.html",
			asset: &RemoteAsset{
				url: "http://example.com/url.html",
			},
			want: false,
		},
		{
			description: "return true for same files",
			d:           "http://example.com/url.html",
			asset: &RemoteAsset{
				url: "http://example.com/url.html",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := tt.asset.Debug(tt.d)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Unexpected result; Diff %v", diff)
			}
		})
	}
}

func TestRemoteAsset_String(t *testing.T) {
	tests := []struct {
		description string
		asset       *RemoteAsset
		want        string
	}{
		{
			description: "return string value for same paths",
			asset: &RemoteAsset{
				url: "http://example.com/url.html",
			},
			want: `<Remote Asset: "http://example.com/url.html">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := tt.asset.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Unexpected result; Diff %v", diff)
			}
		})
	}
}

func TestRemoteAsset_Contents(t *testing.T) {
	tests := []struct {
		description string
		asset       *RemoteAsset
		want        string
		wantError   error
	}{
		{
			description: "return error if read fails",
			asset: &RemoteAsset{
				url: "http://example.com/url.html",
			},
			wantError: errNoContents,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got, err := tt.asset.Contents()
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Unexpected result; Diff %v", diff)
			}
		})
	}
}

func TestRemoteAsset_URL(t *testing.T) {
	tests := []struct {
		description string
		asset       *RemoteAsset
		want        string
		wantError   error
	}{
		{
			description: "return error if read fails",
			asset: &RemoteAsset{
				url: "http://example.com/url.html",
			},
			want: "http://example.com/url.html",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got, err := tt.asset.URL()
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Unexpected result; Diff %v", diff)
			}
		})
	}
}
