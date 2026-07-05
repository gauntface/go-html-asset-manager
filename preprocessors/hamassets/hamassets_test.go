package hamassets

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gauntface/go-html-asset-manager/v5/assets/assetstubs"
	"github.com/gauntface/go-html-asset-manager/v5/preprocessors"
)

func Test_Preprocessor(t *testing.T) {
	t.Run("does nothing if static dir is empty", func(t *testing.T) {
		manager := &assetstubs.Manager{}
		err := Preprocessor(preprocessors.Runtime{Assets: manager})
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if len(manager.AddLocalCalls) != 0 {
			t.Fatalf("Expected no local assets to be added, got %v", len(manager.AddLocalCalls))
		}
	})

	t.Run("copies embedded assets and registers them as local assets", func(t *testing.T) {
		dir := t.TempDir()
		manager := &assetstubs.Manager{
			StaticDirReturn: dir,
		}

		err := Preprocessor(preprocessors.Runtime{Assets: manager})
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if len(manager.AddLocalCalls) == 0 {
			t.Fatalf("Expected at least one local asset to be registered")
		}

		for _, a := range manager.AddLocalCalls {
			p := a.Path()
			if !filepath.IsAbs(p) {
				p = filepath.Join(dir, p)
			}
			if _, err := os.Stat(p); err != nil {
				t.Fatalf("Expected copied asset to exist on disk at %q: %v", p, err)
			}
		}
	})
}
