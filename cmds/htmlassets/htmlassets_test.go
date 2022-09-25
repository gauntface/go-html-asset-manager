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
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/gauntface/go-html-asset-manager/v4/assets"
	"github.com/gauntface/go-html-asset-manager/v4/assets/assetmanager"
	"github.com/gauntface/go-html-asset-manager/v4/assets/assetstubs"
	"github.com/gauntface/go-html-asset-manager/v4/manipulations"
	"github.com/gauntface/go-html-asset-manager/v4/preprocessors"
	"github.com/gauntface/go-html-asset-manager/v4/utils/config"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/mitchellh/go-homedir"
	"github.com/otiai10/copy"
	"golang.org/x/net/html"
)

var errInjected = errors.New("injected error")

var reset func()

func TestMain(m *testing.M) {
	origDebug := debug
	origConfigPath := configPath
	origNewManager := assetmanagerNewManager
	origHomeDirExpand := homedirExpand
	origConfigGet := configGet
	origVimeo := vimeoToken

	reset = func() {
		debug = origDebug
		configPath = origConfigPath
		assetmanagerNewManager = origNewManager
		homedirExpand = origHomeDirExpand
		configGet = origConfigGet
		vimeoToken = origVimeo
	}

	os.Exit(m.Run())
}

func Test_prettyPrintAssets(t *testing.T) {
	tests := []struct {
		description string
		debug       string
		assets      *assetmanager.Manager
	}{
		{
			description: "do nothing when debug is empty",
			debug:       "",
		},
		{
			description: "log assets",
			debug:       "example.html",
			assets:      &assetmanager.Manager{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			defer reset()

			debug = &tt.debug

			prettyPrintAssets(tt.assets)
		})
	}
}

func Test_newClient(t *testing.T) {
	tests := []struct {
		description   string
		configPath    string
		vimeoToken    string
		newManager    func(htmlDir, staticDir, jsonDir string) (*assetmanager.Manager, error)
		homedirExpand func(path string) (string, error)
		configGet     func(path string) (*config.Config, error)
		want          *client
		wantError     error
	}{
		{
			description: "return error if expanding the config path fails",
			configPath:  "/config.json",
			homedirExpand: func(path string) (string, error) {
				if path == "/config.json" {
					return "", errInjected
				}
				return "", nil
			},
			wantError: errInjected,
		},
		{
			description: "return error if creating new manager fails",
			configPath:  "/config.json",
			configGet: func(path string) (*config.Config, error) {
				return &config.Config{
					Assets: &config.AssetsConfig{},
				}, nil
			},
			homedirExpand: homedir.Expand,
			newManager: func(htmlDir, staticDir, jsonDir string) (*assetmanager.Manager, error) {
				return nil, errInjected
			},
			wantError: errInjected,
		},
		{
			description: "return client without optional values",
			configPath:  "/config.json",
			configGet: func(path string) (*config.Config, error) {
				return &config.Config{
					Assets: &config.AssetsConfig{},
				}, nil
			},
			homedirExpand: homedir.Expand,
			newManager: func(htmlDir, staticDir, jsonDir string) (*assetmanager.Manager, error) {
				return &assetmanager.Manager{}, nil
			},
			want: &client{},
		},
		{
			description: "return client without all values",
			configPath:  "/config.json",
			vimeoToken:  "example-vimeo-token",
			configGet: func(path string) (*config.Config, error) {
				return &config.Config{
					Assets: &config.AssetsConfig{},
				}, nil
			},
			homedirExpand: homedir.Expand,
			newManager: func(htmlDir, staticDir, jsonDir string) (*assetmanager.Manager, error) {
				return &assetmanager.Manager{}, nil
			},
			want: &client{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			defer reset()

			configPath = &tt.configPath
			assetmanagerNewManager = tt.newManager
			homedirExpand = tt.homedirExpand
			configGet = tt.configGet
			vimeoToken = &tt.vimeoToken

			got, err := newClient()
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			opts := []cmp.Option{
				cmpopts.IgnoreUnexported(client{}),
			}
			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}

func Test_manipulateHTMLFile(t *testing.T) {
	tests := []struct {
		description   string
		asset         *assetstubs.Asset
		manager       *assetmanager.Manager
		manipulations []manipulations.Manipulator
		open          func(f string) (*os.File, error)
		htmlParse     func(r io.Reader) (*html.Node, error)
		htmlRender    func(w io.Writer, n *html.Node) error
		writeFile     func(filename string, data []byte, perm os.FileMode) error
		wantError     error
	}{
		{
			description: "return error if opening file fails",
			asset: &assetstubs.Asset{
				ContentsError: errInjected,
			},
			wantError: errInjected,
		},
		{
			description: "return error if parsing the HTML fails",
			asset: &assetstubs.Asset{
				ContentsReturn: "example",
			},
			htmlParse: func(r io.Reader) (*html.Node, error) {
				return nil, errInjected
			},
			wantError: errInjected,
		},
		{
			description: "return error if inserting keys fails",
			asset: &assetstubs.Asset{
				ContentsReturn: "example",
			},
			htmlParse: func(r io.Reader) (*html.Node, error) {
				return MustGetNode(t, ""), nil
			},
			manipulations: []manipulations.Manipulator{
				func(runtime manipulations.Runtime, doc *html.Node) error {
					return errInjected
				},
			},
			wantError: errInjected,
		},
		{
			description: "return error if writing file fails",
			asset: &assetstubs.Asset{
				ContentsReturn: "example",
			},
			htmlParse: func(r io.Reader) (*html.Node, error) {
				return MustGetNode(t, ""), nil
			},
			htmlRender: func(w io.Writer, n *html.Node) error {
				return errInjected
			},
			wantError: errInjected,
		},
		{
			description: "return nothing on success",
			asset: &assetstubs.Asset{
				ContentsReturn: "example",
			},
			htmlParse: func(r io.Reader) (*html.Node, error) {
				return MustGetNode(t, ""), nil
			},
			htmlRender: func(w io.Writer, n *html.Node) error {
				return nil
			},
			writeFile: func(filename string, data []byte, perm os.FileMode) error {
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			c := &client{
				htmlParse:       tt.htmlParse,
				htmlRender:      tt.htmlRender,
				ioutilWriteFile: tt.writeFile,
			}
			err := c.manipulateHTMLFile(tt.asset, tt.manager, tt.manipulations)
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Unexpected error; got %v, want %v", err, tt.wantError)
			}
		})
	}
}

func Test_preprocesses(t *testing.T) {
	tests := []struct {
		description   string
		manager       *assetstubs.Manager
		preprocessors []preprocessors.Preprocessor
		wantErrors    []error
	}{
		{
			description: "return errors if processing file fails",
			preprocessors: []preprocessors.Preprocessor{
				func(runtime preprocessors.Runtime) error {
					return errInjected
				},
			},
			wantErrors: []error{
				errInjected,
			},
		},
		{
			description: "return no errors on success",
			manager:     &assetstubs.Manager{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			c := &client{}
			errs := c.preprocesses(tt.manager, tt.preprocessors)
			if len(errs) != len(tt.wantErrors) {
				t.Fatalf("Unexpected errors; got %v, want %v", errs, tt.wantErrors)
			}
			for i, e := range errs {
				if !errors.Is(e, tt.wantErrors[i]) {
					t.Fatalf("Unexpected error at %v; got %v, want %v", i, e, tt.wantErrors[i])
				}
			}
		})
	}
}

func Test_manipulations(t *testing.T) {
	tests := []struct {
		description   string
		manager       *assetstubs.Manager
		manipulations []manipulations.Manipulator
		wantErrors    []error
	}{
		{
			description: "return errors if processing file fails",
			manager: &assetstubs.Manager{
				WithTypeReturn: map[assets.Type][]assetmanager.Asset{
					assets.HTML: []assetmanager.Asset{
						assetstubs.MustNewLocalAsset(t, "/example/", "/example/example-1.html"),
						assetmanager.NewRemoteAsset("example", "http://example.com/123", assets.Unknown),
					},
				},
			},
			wantErrors: []error{
				errManipulate,
			},
		},
		{
			description: "return no errors on success",
			manager:     &assetstubs.Manager{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			c := &client{}
			errs := c.manipulations(tt.manager, tt.manipulations)
			if len(errs) != len(tt.wantErrors) {
				t.Fatalf("Unexpected errors; got %v, want %v", errs, tt.wantErrors)
			}
			for i, e := range errs {
				if !errors.Is(e, tt.wantErrors[i]) {
					t.Fatalf("Unexpected error at %v; got %v, want %v", i, e, tt.wantErrors[i])
				}
			}
		})
	}
}

func Test_writeChanges(t *testing.T) {
	tests := []struct {
		description string
		htmlFile    string
		node        *html.Node
		render      func(w io.Writer, n *html.Node) error
		writeFile   func(filename string, data []byte, perm os.FileMode) error
		wantError   error
	}{
		{
			description: "return error if render fails",
			render: func(w io.Writer, n *html.Node) error {
				return errInjected
			},
			wantError: errInjected,
		},
		{
			description: "return error if write fails",
			render: func(w io.Writer, n *html.Node) error {
				return nil
			},
			writeFile: func(filename string, data []byte, perm os.FileMode) error {
				return errInjected
			},
			wantError: errInjected,
		},
		{
			description: "return nothing on success",
			render: func(w io.Writer, n *html.Node) error {
				return nil
			},
			writeFile: func(filename string, data []byte, perm os.FileMode) error {
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			c := &client{
				htmlRender:      tt.render,
				ioutilWriteFile: tt.writeFile,
			}
			err := c.writeChanges(tt.htmlFile, tt.node)
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Unexpected error; got %v, want %v", err, tt.wantError)
			}
		})
	}
}

func Test_run(t *testing.T) {
	tests := []struct {
		description   string
		manager       *assetstubs.Manager
		preprocessors []preprocessors.Preprocessor
		manipulations []manipulations.Manipulator
		wantError     error
	}{
		{
			description: "return error if preprocessors fail",
			manager:     &assetstubs.Manager{},
			preprocessors: []preprocessors.Preprocessor{
				func(runtime preprocessors.Runtime) error {
					return errInjected
				},
			},
			wantError: errRunFailed,
		},
		{
			description: "return error if manipulations fail",
			manager: &assetstubs.Manager{
				WithTypeReturn: map[assets.Type][]assetmanager.Asset{
					assets.HTML: {
						assetstubs.MustNewLocalAsset(t, "testdata/noassets/", "index.html"),
					},
				},
			},
			preprocessors: []preprocessors.Preprocessor{},
			manipulations: []manipulations.Manipulator{
				func(runtime manipulations.Runtime, doc *html.Node) error {
					return errInjected
				},
			},
			wantError: errRunFailed,
		},
		{
			description: "return nothing on success",
			manager:     &assetstubs.Manager{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			c := &client{
				manager:       tt.manager,
				preprocessors: tt.preprocessors,
				manipulators:  tt.manipulations,
			}
			err := c.run()
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Unexpected error; got %v, want %v", err, tt.wantError)
			}
		})
	}
}

func Test_integration_noassets(t *testing.T) {
	defer reset()

	tmpDir, err := ioutil.TempDir("/tmp", "htmlassets")
	if err != nil {
		t.Fatalf("Fatal to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	testdataDir := path.Join("..", "..", "testdata", "noassets")
	err = copy.Copy(testdataDir, tmpDir)
	if err != nil {
		t.Fatalf("Fatal to copy files to temporary directory: %v", err)
	}

	tmpConf := filepath.Join(tmpDir, "asset-manager.json")
	configPath = &tmpConf

	c, err := newClient()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	err = c.run()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	got := readTestFile(t, path.Join(tmpDir, "index.html"))
	want := readTestFile(t, path.Join(testdataDir, "index-want.html"))

	if diff := cmp.Diff(got, want); diff == "" {
		t.Fatalf("Unexpected result; diff %v", diff)
	}

}

func readTestFile(t *testing.T, file string) string {
	t.Helper()

	fileBuffer, err := ioutil.ReadFile(file)
	if err != nil {
		t.Fatalf("Failed to read tmp file: %v", err)
	}

	content := string(fileBuffer)
	content = strings.ReplaceAll(content, "\n", "")
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(content, " ")
}

func MustGetNode(t *testing.T, input string) *html.Node {
	t.Helper()

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Failed to parse HTML: %v", err)
	}
	return doc
}
