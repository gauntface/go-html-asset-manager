package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeConfig(t *testing.T, dir, contents string) string {
	t.Helper()

	p := filepath.Join(dir, "asset-manager.json")
	if err := os.WriteFile(p, []byte(contents), 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}
	return p
}

func Test_Get(t *testing.T) {
	t.Run("returns an error if the file cannot be read", func(t *testing.T) {
		_, err := Get(filepath.Join(t.TempDir(), "does-not-exist.json"))
		if err == nil {
			t.Fatalf("Expected an error, got nil")
		}
	})

	t.Run("returns an error if the file is not valid JSON", func(t *testing.T) {
		dir := t.TempDir()
		p := writeConfig(t, dir, `not json`)

		_, err := Get(p)
		if err == nil {
			t.Fatalf("Expected an error, got nil")
		}
	})

	t.Run("resolves a relative html-dir against the config file's directory", func(t *testing.T) {
		dir := t.TempDir()
		p := writeConfig(t, dir, `{"html-dir": "public"}`)

		conf, err := Get(p)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		want := filepath.Join(dir, "public")
		if conf.HTMLDir != want {
			t.Fatalf("Unexpected HTMLDir; got %v, want %v", conf.HTMLDir, want)
		}
	})

	t.Run("leaves an absolute html-dir untouched", func(t *testing.T) {
		dir := t.TempDir()
		p := writeConfig(t, dir, `{"html-dir": "/already/absolute"}`)

		conf, err := Get(p)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if conf.HTMLDir != "/already/absolute" {
			t.Fatalf("Unexpected HTMLDir; got %v, want %v", conf.HTMLDir, "/already/absolute")
		}
	})

	t.Run("does not resolve assets paths if assets is not set", func(t *testing.T) {
		dir := t.TempDir()
		p := writeConfig(t, dir, `{"html-dir": "public"}`)

		conf, err := Get(p)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if conf.Assets != nil {
			t.Fatalf("Expected Assets to be nil, got %+v", conf.Assets)
		}
	})

	t.Run("resolves relative assets static-dir and json-dir", func(t *testing.T) {
		dir := t.TempDir()
		p := writeConfig(t, dir, `{
			"html-dir": "public",
			"assets": {
				"static-dir": "public/static",
				"json-dir": "data"
			}
		}`)

		conf, err := Get(p)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if want := filepath.Join(dir, "public/static"); conf.Assets.StaticDir != want {
			t.Fatalf("Unexpected Assets.StaticDir; got %v, want %v", conf.Assets.StaticDir, want)
		}
		if want := filepath.Join(dir, "data"); conf.Assets.JSONDir != want {
			t.Fatalf("Unexpected Assets.JSONDir; got %v, want %v", conf.Assets.JSONDir, want)
		}
	})

	t.Run("leaves empty assets static-dir and json-dir untouched", func(t *testing.T) {
		dir := t.TempDir()
		p := writeConfig(t, dir, `{
			"html-dir": "public",
			"assets": {}
		}`)

		conf, err := Get(p)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if conf.Assets.StaticDir != "" {
			t.Fatalf("Expected Assets.StaticDir to stay empty, got %v", conf.Assets.StaticDir)
		}
		if conf.Assets.JSONDir != "" {
			t.Fatalf("Expected Assets.JSONDir to stay empty, got %v", conf.Assets.JSONDir)
		}
	})

	t.Run("resolves relative gen-assets static-dir and output-dir", func(t *testing.T) {
		dir := t.TempDir()
		p := writeConfig(t, dir, `{
			"html-dir": "public",
			"gen-assets": {
				"static-dir": "static",
				"output-dir": "static/generated"
			}
		}`)

		conf, err := Get(p)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if want := filepath.Join(dir, "static"); conf.GenAssets.StaticDir != want {
			t.Fatalf("Unexpected GenAssets.StaticDir; got %v, want %v", conf.GenAssets.StaticDir, want)
		}
		if want := filepath.Join(dir, "static/generated"); conf.GenAssets.OutputDir != want {
			t.Fatalf("Unexpected GenAssets.OutputDir; got %v, want %v", conf.GenAssets.OutputDir, want)
		}
	})

	t.Run("resolves an empty gen-assets static-dir/output-dir to the config file's directory", func(t *testing.T) {
		// Unlike assets.static-dir/json-dir, gen-assets' fields are resolved
		// unconditionally rather than only when non-empty, so an omitted
		// value collapses to the config file's own directory.
		dir := t.TempDir()
		p := writeConfig(t, dir, `{
			"html-dir": "public",
			"gen-assets": {}
		}`)

		conf, err := Get(p)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if !strings.HasPrefix(conf.GenAssets.StaticDir, dir) {
			t.Fatalf("Expected GenAssets.StaticDir to resolve under %v, got %v", dir, conf.GenAssets.StaticDir)
		}
	})

	t.Run("parses img-to-picture and ratio-wrapper", func(t *testing.T) {
		dir := t.TempDir()
		p := writeConfig(t, dir, `{
			"html-dir": "public",
			"img-to-picture": [
				{"id": "c-example__img", "max-width": 620, "source-sizes": ["100vw"], "class": "c-picture"}
			],
			"ratio-wrapper": ["iframe", "c-example"]
		}`)

		conf, err := Get(p)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if len(conf.ImgToPicture) != 1 {
			t.Fatalf("Unexpected number of ImgToPicture entries; got %v, want 1", len(conf.ImgToPicture))
		}
		itp := conf.ImgToPicture[0]
		if itp.ID != "c-example__img" || itp.MaxWidth != 620 || itp.Class != "c-picture" {
			t.Fatalf("Unexpected ImgToPicture entry: %+v", itp)
		}

		wantRatioWrapper := []string{"iframe", "c-example"}
		if len(conf.RatioWrapper) != len(wantRatioWrapper) {
			t.Fatalf("Unexpected RatioWrapper; got %v, want %v", conf.RatioWrapper, wantRatioWrapper)
		}
		for i, v := range wantRatioWrapper {
			if conf.RatioWrapper[i] != v {
				t.Fatalf("Unexpected RatioWrapper[%v]; got %v, want %v", i, conf.RatioWrapper[i], v)
			}
		}
	})
}
