package revisionassets

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/gauntface/go-html-asset-manager/v5/assets"
	"github.com/gauntface/go-html-asset-manager/v5/assets/assetmanager"
	"github.com/gauntface/go-html-asset-manager/v5/assets/assetstubs"
	"github.com/gauntface/go-html-asset-manager/v5/preprocessors"
)

var errInjected = errors.New("injected error")

func writeTempFile(t *testing.T, dir, name, contents string) string {
	t.Helper()

	p := filepath.Join(dir, name)
	if err := os.WriteFile(p, []byte(contents), 0644); err != nil {
		t.Fatalf("Failed to write temp file: %v", err)
	}
	return p
}

func Test_Preprocessor(t *testing.T) {
	t.Run("does nothing if there are no assets", func(t *testing.T) {
		manager := &assetstubs.Manager{}
		if err := Preprocessor(preprocessors.Runtime{Assets: manager}); err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("skips remote assets", func(t *testing.T) {
		manager := &assetstubs.Manager{
			AllReturn: []assetmanager.Asset{
				assetmanager.NewRemoteAsset("example", "https://example.com/sync.js", nil, assets.SyncJS),
			},
		}
		if err := Preprocessor(preprocessors.Runtime{Assets: manager}); err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("skips local assets with a non-revisionable type", func(t *testing.T) {
		dir := t.TempDir()
		p := writeTempFile(t, dir, "index.html", "<html></html>")
		asset := assetstubs.MustNewLocalAsset(t, dir, p)

		manager := &assetstubs.Manager{
			AllReturn: []assetmanager.Asset{asset},
		}
		if err := Preprocessor(preprocessors.Runtime{Assets: manager}); err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if asset.Path() != p {
			t.Fatalf("Expected path to be unchanged; got %v, want %v", asset.Path(), p)
		}
		if _, err := os.Stat(p); err != nil {
			t.Fatalf("Expected original file to still exist: %v", err)
		}
	})

	t.Run("renames revisionable local assets with a content hash and updates their path", func(t *testing.T) {
		dir := t.TempDir()
		p := writeTempFile(t, dir, "styles-sync.css", "body { color: red; }")
		asset := assetstubs.MustNewLocalAsset(t, dir, p)

		manager := &assetstubs.Manager{
			AllReturn: []assetmanager.Asset{asset},
		}
		if err := Preprocessor(preprocessors.Runtime{Assets: manager}); err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if asset.Path() == p {
			t.Fatalf("Expected path to be updated with a revision hash, still %v", p)
		}
		if _, err := os.Stat(p); err == nil {
			t.Fatalf("Expected original file to no longer exist at %v", p)
		}
		if _, err := os.Stat(asset.Path()); err != nil {
			t.Fatalf("Expected revisioned file to exist at %v: %v", asset.Path(), err)
		}
	})

	t.Run("returns an error if hashing the file fails", func(t *testing.T) {
		dir := t.TempDir()
		// Reference a file that doesn't exist so files.Hash fails to open it.
		p := filepath.Join(dir, "styles-sync.css")
		asset := assetstubs.MustNewLocalAsset(t, dir, p)

		manager := &assetstubs.Manager{
			AllReturn: []assetmanager.Asset{asset},
		}
		if err := Preprocessor(preprocessors.Runtime{Assets: manager}); err == nil {
			t.Fatalf("Expected an error, got nil")
		}
	})

	t.Run("returns a wrapped error if renaming fails", func(t *testing.T) {
		origRename := osRename
		defer func() { osRename = origRename }()
		osRename = func(oldpath, newpath string) error {
			return errInjected
		}

		dir := t.TempDir()
		p := writeTempFile(t, dir, "styles-sync.css", "body { color: red; }")
		asset := assetstubs.MustNewLocalAsset(t, dir, p)

		manager := &assetstubs.Manager{
			AllReturn: []assetmanager.Asset{asset},
		}
		err := Preprocessor(preprocessors.Runtime{Assets: manager})
		if !errors.Is(err, errRenameFailed) {
			t.Fatalf("Unexpected error; got %v, want %v", err, errRenameFailed)
		}
	})
}
