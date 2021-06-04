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
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/gauntface/go-html-asset-manager/assets"
	"github.com/gauntface/go-html-asset-manager/assets/assetmanager"
	"github.com/gauntface/go-html-asset-manager/manipulations"
	"github.com/gauntface/go-html-asset-manager/manipulations/asyncsrc"
	"github.com/gauntface/go-html-asset-manager/manipulations/iframedefaultsize"
	"github.com/gauntface/go-html-asset-manager/manipulations/imgtopicture"
	"github.com/gauntface/go-html-asset-manager/manipulations/injectassets"
	"github.com/gauntface/go-html-asset-manager/manipulations/lazyload"
	"github.com/gauntface/go-html-asset-manager/manipulations/opengraphimg"
	"github.com/gauntface/go-html-asset-manager/manipulations/ratiowrapper"
	"github.com/gauntface/go-html-asset-manager/manipulations/stripassets"
	"github.com/gauntface/go-html-asset-manager/manipulations/vimeoclean"
	"github.com/gauntface/go-html-asset-manager/manipulations/youtubeclean"
	"github.com/gauntface/go-html-asset-manager/preprocessors"
	"github.com/gauntface/go-html-asset-manager/preprocessors/jsonassets"
	"github.com/gauntface/go-html-asset-manager/preprocessors/revisionassets"
	"github.com/gauntface/go-html-asset-manager/utils/config"
	"github.com/gauntface/go-html-asset-manager/utils/html/htmlencoding"
	"github.com/gauntface/go-html-asset-manager/utils/vimeoapi"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/net/html"
)

var (
	configPath = flag.String("config", "asset-manager.json", "The path of the Config file.")
	vimeoToken = flag.String("vimeo", "", "Personal access token for Vimeo API")
	debug      = flag.String("debug", "", "Provide a HTML file name to log debug info as required")

	errRunFailed  = errors.New("failed to run successfully")
	errRelPath    = errors.New("unable to calculate relative path")
	errManipulate = errors.New("failed to manipulate HTML")

	configGet              = config.Get
	homedirExpand          = homedir.Expand
	assetmanagerNewManager = assetmanager.NewManager
)

func main() {
	c, err := newClient()
	if err != nil {
		fmt.Printf("Could not initialize client: %v", err)
		os.Exit(1)
	}
	if err := c.run(); err != nil {
		fmt.Printf("Run was not successful: %v", err)
		os.Exit(1)
	}
}

type client struct {
	homedirExpand          func(path string) (string, error)
	htmlParse              func(r io.Reader) (*html.Node, error)
	htmlRender             func(w io.Writer, n *html.Node) error
	ioutilWriteFile        func(filename string, data []byte, perm os.FileMode) error
	assetmanagerNewManager func(htmlDir, staticDir, jsonDir string) (*assetmanager.Manager, error)

	config        *config.Config
	manager       assetmanagerManager
	vimeo         *vimeoapi.Client
	preprocessors []preprocessors.Preprocessor
	manipulators  []manipulations.Manipulator
}

func newClient() (*client, error) {
	flag.Parse()

	absConfigPath, err := homedirExpand(*configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path for config flag: %w", err)
	}
	fmt.Printf("📁 Getting config file: %q\n", absConfigPath)

	c, err := configGet(absConfigPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path for html_dir flag: %w", err)
	}

	fmt.Printf("📁 Looking for HTML files in:   %q\n", c.HTMLDir)
	fmt.Printf("📁 Looking for Static assets in: %q\n", c.Assets.StaticDir)
	fmt.Printf("📁 Looking for JSON assets in: %q\n", c.Assets.JSONDir)
	if c.GenAssets != nil {
		fmt.Printf("📁 Looking for generated assets in: %q\n", c.GenAssets.OutputDir)
	}
	fmt.Println("")

	manager, err := assetmanagerNewManager(c.HTMLDir, c.Assets.StaticDir, c.Assets.JSONDir)
	if err != nil {
		return nil, err
	}

	var vimeo *vimeoapi.Client
	if *vimeoToken != "" {
		vimeo = vimeoapi.New(*vimeoToken)
	}

	return &client{
		htmlRender:      html.Render,
		htmlParse:       html.Parse,
		ioutilWriteFile: ioutil.WriteFile,

		config:  c,
		manager: manager,
		vimeo:   vimeo,
		preprocessors: []preprocessors.Preprocessor{
			jsonassets.Preprocessor,
			revisionassets.Preprocessor,
		},
		manipulators: []manipulations.Manipulator{
			opengraphimg.Manipulator,
			youtubeclean.Manipulator,
			vimeoclean.Manipulator,
			iframedefaultsize.Manipulator,
			imgtopicture.Manipulator,
			ratiowrapper.Manipulator,
			lazyload.Manipulator,
			asyncsrc.Manipulator,
			stripassets.Manipulator,
			injectassets.Manipulator,
		},
	}, nil
}

func (c *client) run() error {
	prettyPrintAssets(c.manager)

	// Step 1: Run preprocessprs
	errs := c.preprocesses(c.manager, c.preprocessors)
	if len(errs) > 0 {
		return logReturn(errRunFailed, errs)
	}

	// Step 2: Run HTML manipulation steps
	errs = c.manipulations(c.manager, c.manipulators)
	if len(errs) > 0 {
		return logReturn(errRunFailed, errs)
	}

	return nil
}

func (c *client) preprocesses(manager assetmanagerManager, preprocesses []preprocessors.Preprocessor) []error {
	errs := []error{}

	runtime := preprocessors.Runtime{
		Assets: manager,
	}
	for i, p := range preprocesses {
		err := p(runtime)
		if err != nil {
			errs = append(errs, fmt.Errorf("preprocessor %v failed: %w", i, err))
		}
	}

	return errs
}

func (c *client) manipulations(manager assetmanagerManager, manipulators []manipulations.Manipulator) []error {
	htmlAssets := manager.WithType(assets.HTML)

	las := []assetmanagerLocalAsset{}
	for _, a := range htmlAssets {
		if !a.IsLocal() {
			continue
		}

		las = append(las, a.(*assetmanager.LocalAsset))
	}

	return c.manipulateHTMLFiles(las, manager, manipulators)
}

func (c *client) manipulateHTMLFiles(assets []assetmanagerLocalAsset, manager assetmanagerManager, manipulators []manipulations.Manipulator) []error {
	var wg sync.WaitGroup
	wg.Add(len(assets))

	errs := []error{}
	var errMu sync.Mutex

	for _, htmlAsset := range assets {
		go func(htmlAsset assetmanagerLocalAsset) {
			defer wg.Done()

			err := c.manipulateHTMLFile(htmlAsset, manager, manipulators)
			if err != nil {
				errMu.Lock()
				defer errMu.Unlock()
				errs = append(errs, fmt.Errorf("%w %q: %v", errManipulate, htmlAsset.Path(), err))
			}
		}(htmlAsset)
	}

	wg.Wait()

	return errs
}

func (c *client) manipulateHTMLFile(asset assetmanagerLocalAsset, manager assetmanagerManager, manips []manipulations.Manipulator) error {
	html, err := asset.Contents()
	if err != nil {
		return err
	}

	doc, err := c.htmlParse(strings.NewReader(html))
	if err != nil {
		return fmt.Errorf("failed to parse file %q: %w", asset, err)
	}

	debug := *debug != "" && asset.Debug(*debug)
	r := manipulations.Runtime{
		Debug:    debug,
		Assets:   manager,
		Vimeo:    c.vimeo,
		HasVimeo: c.vimeo != nil,
		Config:   c.config,
	}
	for i, m := range manips {
		if err := m(r, doc); err != nil {
			return fmt.Errorf(`Manipulation %v failed: %w`, i, err)
		}
	}

	err = c.writeChanges(asset.Path(), doc)
	if err != nil {
		return fmt.Errorf("failed to write changes: %w", err)
	}

	return nil
}

func (c *client) writeChanges(htmlFile string, doc *html.Node) error {
	htmlencoding.EncodeNodes(doc)
	var buf bytes.Buffer
	err := c.htmlRender(&buf, doc)
	if err != nil {
		return fmt.Errorf("failed to render html node to string: %w", err)
	}

	err = c.ioutilWriteFile(htmlFile, []byte(buf.String()), 0644)
	if err != nil {
		return fmt.Errorf("failed to write changes to %q: %w", htmlFile, err)
	}
	return nil
}

func prettyPrintAssets(assets assetmanagerManager) {
	if *debug == "" {
		return
	}

	fmt.Printf("Found the following assets:\n\n")
	fmt.Printf("%v\n", assets.String())
}

func logReturn(e error, errs []error) error {
	for i, e := range errs {
		fmt.Printf("    - %v) %v\n", i+1, e)
	}
	return fmt.Errorf("%w: %v errors occurred", e, len(errs))
}

type assetmanagerManager interface {
	AddRemote(a *assetmanager.RemoteAsset)
	All() []assetmanager.Asset
	String() string
	WithID(id string) map[assets.Type][]assetmanager.Asset
	WithType(t assets.Type) []assetmanager.Asset
}

type assetmanagerLocalAsset interface {
	Contents() (string, error)
	Debug(string) bool
	Path() string
}
