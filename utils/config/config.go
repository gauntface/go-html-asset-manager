package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

var (
	homedirExpand = homedir.Expand
)

// Config defines all config options for a config file.
type Config struct {
	// The path to a directory containing HTML files
	HTMLDir string `json:"html-dir"`

	// The base URL of the site
	BaseURL string `json:"base-url"`

	// The assets to be used by the html-asset-manager
	Assets *AssetsConfig `json:"assets"`

	// The gen-images to generate and use by the html-asset-manager
	GenAssets *GeneratedImagesConfig `json:"gen-assets"`

	// The img-to-picture manipulation config
	ImgToPicture []*ImgToPicConfig `json:"img-to-picture"`

	// The ratio-wrapper manipulation config
	RatioWrapper []string `json:"ratio-wrapper"`
}

// AssetsConfig defines config options for assets
type AssetsConfig struct {
	// path to a directory containing CSS and JS files
	StaticDir string `json:"static-dir"`
	// path to a directory containing JSON files for asset injection
	JSONDir string `json:"json-dir"`
	// path to a directory containing generate files
	GeneratedDir string `json:"generated-dir"`
}

// GeneratedImagesConfig defines config options for gen-imgs cmd
type GeneratedImagesConfig struct {
	// path to a directory containing static assets
	StaticDir string `json:"static-dir"`
	// The path to a directory containing generated files
	OutputDir string `json:"output-dir"`
	// The maximum width in CSS pixels images should be
	MaxWidth int64 `json:"max-width"`
	// The maximum density to cater for
	MaxDensity int64 `json:"max-density"`
}

// ImgToPicConfig defines config options for the img-to-picture manipulation
type ImgToPicConfig struct {
	// The classname or tag of elements to replace img to picture
	ID string `json:"id"`
	// The maximum width the image will be in CSS pixels
	MaxWidth int64 `json:"max-width"`
	// The source sizes the picture should have
	SourceSizes []string `json:"source-sizes"`
	// Class to apply to the picture element
	Class string `json:"class"`
}

// Get reads and parses a Config file
func Get(filePath string) (*Config, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %q: %v", filePath, err)
	}

	var conf Config
	err = json.Unmarshal(b, &conf)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file %q: %v", filePath, err)
	}

	dir := filepath.Dir(filePath)
	conf.HTMLDir = abs(dir, conf.HTMLDir)

	if conf.Assets != nil {
		conf.Assets.StaticDir = abs(dir, conf.Assets.StaticDir)
		conf.Assets.JSONDir = abs(dir, conf.Assets.JSONDir)
		conf.Assets.GeneratedDir = abs(dir, conf.Assets.GeneratedDir)
	}
	if conf.GenAssets != nil {
		conf.GenAssets.StaticDir = abs(dir, conf.GenAssets.StaticDir)
		conf.GenAssets.OutputDir = abs(dir, conf.GenAssets.OutputDir)
	}

	return &conf, nil
}

func abs(relDir, path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(relDir, path)
}
