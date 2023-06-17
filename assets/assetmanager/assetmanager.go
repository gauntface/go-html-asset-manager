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
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gauntface/go-html-asset-manager/v5/assets"
	"github.com/gauntface/go-html-asset-manager/v5/assets/assetid"
	"github.com/gauntface/go-html-asset-manager/v5/utils/files"
	"github.com/gauntface/go-html-asset-manager/v5/utils/stringui"
	"golang.org/x/net/html"
)

var (
	errRelPath    = errors.New("unable to calculate relative path")
	errReadFailed = errors.New("unable to read file")
	errNoContents = errors.New("unable to get contents for asset")

	filesFind = files.Find
)

type Manager struct {
	htmlDir   string
	staticDir string
	jsonDir   string

	localAssets  []*LocalAsset
	remoteAssets []*RemoteAsset
}

func NewManager(htmlDir, staticDir, jsonDir string) (*Manager, error) {
	htmlAssets, err := findLocalAssets(htmlDir, ".html")
	if err != nil {
		return nil, err
	}

	staticAssets, err := findLocalAssets(staticDir, ".css", ".js", ".png", ".jpg", ".jpeg", ".webp", ".avif")
	if err != nil {
		return nil, err
	}

	jsonAssets, err := findLocalAssets(jsonDir, ".json")
	if err != nil {
		return nil, err
	}

	la := []*LocalAsset{}
	la = append(la, htmlAssets...)
	la = append(la, staticAssets...)
	la = append(la, jsonAssets...)
	return &Manager{
		htmlDir:   htmlDir,
		staticDir: staticDir,
		jsonDir:   jsonDir,

		localAssets:  la,
		remoteAssets: []*RemoteAsset{},
	}, nil
}

func (m *Manager) StaticDir() string {
	return m.staticDir
}

func (m *Manager) All() []Asset {
	as := []Asset{}
	for _, a := range m.localAssets {
		as = append(as, a)
	}
	for _, a := range m.remoteAssets {
		as = append(as, a)
	}
	return as
}

func (m *Manager) WithType(t assets.Type) []Asset {
	as := []Asset{}

	for _, a := range m.All() {
		if a.Type() == t {
			as = append(as, a)
		}
	}

	return as
}

func (m *Manager) WithID(id string) map[assets.Type][]Asset {
	as := map[assets.Type][]Asset{}

	for _, a := range m.All() {
		if a.ID() != id {
			continue
		}

		if _, ok := as[a.Type()]; !ok {
			as[a.Type()] = []Asset{}
		}
		as[a.Type()] = append(as[a.Type()], a)
	}

	return as
}

func (m *Manager) AddRemote(a *RemoteAsset) {
	m.remoteAssets = append(m.remoteAssets, a)
}

func (m *Manager) AddLocal(a *LocalAsset) {
	m.localAssets = append(m.localAssets, a)
}

func (m *Manager) String() string {
	groups := []struct {
		Title     string
		Type      assets.Type
		CountOnly bool
	}{
		{
			Title:     "HTML",
			Type:      assets.HTML,
			CountOnly: true,
		},
		{
			Title:     "JSON",
			Type:      assets.JSON,
			CountOnly: true,
		},
		{
			Title:     "PNG",
			Type:      assets.PNG,
			CountOnly: true,
		},
		{
			Title:     "JPEG",
			Type:      assets.JPEG,
			CountOnly: true,
		},

		{
			Title:     "WEBP",
			Type:      assets.WEBP,
			CountOnly: true,
		},
		{
			Title:     "AVIF",
			Type:      assets.AVIF,
			CountOnly: true,
		},
		{
			Title: "Inline CSS",
			Type:  assets.InlineCSS,
		},
		{
			Title: "Sync CSS",
			Type:  assets.SyncCSS,
		},
		{
			Title: "Async CSS",
			Type:  assets.AsyncCSS,
		},
		{
			Title: "Preload CSS",
			Type:  assets.PreloadCSS,
		},
		{
			Title: "Inline JS",
			Type:  assets.InlineJS,
		},
		{
			Title: "Sync JS",
			Type:  assets.SyncJS,
		},
		{
			Title: "Async JS",
			Type:  assets.AsyncJS,
		},
		{
			Title: "Preload JS",
			Type:  assets.PreloadJS,
		},
	}

	headings := []string{
		"Key",
		"Count",
	}

	sections := []string{}
	for _, g := range groups {
		assets := m.WithType(g.Type)
		if len(assets) == 0 {
			continue
		}

		keyCount := map[string]int{}
		for _, a := range assets {
			keyCount[a.ID()]++
		}

		lines := []string{
			fmt.Sprintf("%v (Total: %v)", g.Title, len(assets)),
		}

		rows := [][]string{}
		if !g.CountOnly {
			for k, v := range keyCount {
				rows = append(rows, []string{k, fmt.Sprintf("%v", v)})
			}

			sort.Slice(rows, func(i, j int) bool {
				return rows[i][0] < rows[j][0]
			})

			lines = append(lines, stringui.Table(headings, rows))
		}

		sections = append(sections, strings.Join(lines, "\n"))
	}
	return strings.Join(sections, "\n\n")
}

type LocalAsset struct {
	assetType    assets.Type
	assetMedia   string
	id           string
	originalPath string
	path         string
	relativeDir  string
	url          string

	readFile func(string) ([]byte, error)
}

func NewLocalAsset(relDir, assetPath string) (*LocalAsset, error) {
	t, m, err := assetid.IdentifyType(assetPath)
	if err != nil {
		return nil, err
	}

	return &LocalAsset{
		assetType:    t,
		assetMedia:   m,
		id:           assetid.Generate(assetPath),
		originalPath: assetPath,
		path:         assetPath,
		relativeDir:  relDir,

		readFile: ioutil.ReadFile,
	}, nil
}

func (l *LocalAsset) Type() assets.Type {
	return l.assetType
}

func (l *LocalAsset) Media() string {
	return l.assetMedia
}

func (l *LocalAsset) ID() string {
	return l.id
}

func (l *LocalAsset) Path() string {
	return l.path
}

func (l *LocalAsset) UpdatePath(p string) {
	l.path = p
}

func (l *LocalAsset) Contents() (string, error) {
	// Read file
	b, err := l.readFile(l.path)
	if err != nil {
		return "", fmt.Errorf("%w with path %q; %v", errReadFailed, l.path, err)
	}
	return string(b), err
}

func (l *LocalAsset) URL() (string, error) {
	relPath, err := filepath.Rel(l.relativeDir, l.path)
	if err != nil {
		return "", fmt.Errorf("%w for directory %q and file %q; %v", errRelPath, l.relativeDir, l.path, err)
	}

	return path.Join("/", relPath), nil
}

func (l *LocalAsset) Attributes() []html.Attribute {
	return nil
}

func (l *LocalAsset) IsLocal() bool {
	return true
}

func (l *LocalAsset) Debug(d string) bool {
	return strings.Contains(l.originalPath, d)
}

func (l *LocalAsset) String() string {
	if l.originalPath == l.path {
		return fmt.Sprintf("<Local Asset: %q>", l.originalPath)
	}
	return fmt.Sprintf("<Local Asset: %q | %q>", l.originalPath, l.path)
}

type RemoteAsset struct {
	id         string
	url        string
	attributes []html.Attribute
	assetType  assets.Type
}

func NewRemoteAsset(ID, src string, attributes []html.Attribute, ty assets.Type) *RemoteAsset {
	return &RemoteAsset{
		id:         ID,
		url:        src,
		attributes: attributes,
		assetType:  ty,
	}
}

func (r *RemoteAsset) Type() assets.Type {
	return r.assetType
}

func (r *RemoteAsset) Media() string {
	return ""
}

func (r *RemoteAsset) ID() string {
	return r.id
}

func (r *RemoteAsset) URL() (string, error) {
	return r.url, nil
}

func (r *RemoteAsset) Attributes() []html.Attribute {
	return r.attributes
}

func (r *RemoteAsset) Contents() (string, error) {
	return "", fmt.Errorf("%w for %q", errNoContents, r.url)
}

func (r *RemoteAsset) IsLocal() bool {
	return false
}

func (r *RemoteAsset) String() string {
	attrs := []string{}
	for _, a := range r.attributes {
		attrs = append(attrs, fmt.Sprintf("%v=%q", a.Key, a.Val))
	}
	return fmt.Sprintf("<Remote Asset: %q Attributes: [%v]>", r.url, strings.Join(attrs, ", "))
}

func (r *RemoteAsset) Debug(d string) bool {
	return strings.Contains(r.url, d)
}

func findLocalAssets(dir string, exts ...string) ([]*LocalAsset, error) {
	if dir == "" {
		return nil, nil
	}

	files, err := filesFind(dir, exts...)
	if err != nil {
		return nil, err
	}

	assets := []*LocalAsset{}
	for _, f := range files {
		a, err := NewLocalAsset(dir, f)
		if err != nil {
			return nil, err
		}
		assets = append(assets, a)
	}
	return assets, nil
}

type Asset interface {
	ID() string
	Type() assets.Type
	Media() string
	URL() (string, error)
	Contents() (string, error)
	Attributes() []html.Attribute
	IsLocal() bool
	String() string
	Debug(d string) bool
}
