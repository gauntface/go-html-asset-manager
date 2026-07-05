package embedassets

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_CopyAssets(t *testing.T) {
	t.Run("copies every embedded asset file into <staticDir>/__ham", func(t *testing.T) {
		dir := t.TempDir()

		files, err := CopyAssets(dir)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if len(files) == 0 {
			t.Fatalf("Expected at least one file to be copied")
		}

		outputDir := filepath.Join(dir, "__ham")
		for _, f := range files {
			if !strings.HasPrefix(f, outputDir) {
				t.Fatalf("Expected copied file %q to be under %q", f, outputDir)
			}

			info, err := os.Stat(f)
			if err != nil {
				t.Fatalf("Expected copied file %q to exist: %v", f, err)
			}
			if info.Size() == 0 {
				t.Fatalf("Expected copied file %q to be non-empty", f)
			}
		}
	})

	t.Run("returns an error if the output directory cannot be created", func(t *testing.T) {
		dir := t.TempDir()

		// Pre-create a regular file where CopyAssets needs to create the
		// __ham directory, so os.MkdirAll fails.
		if err := os.WriteFile(filepath.Join(dir, "__ham"), []byte("not a directory"), 0644); err != nil {
			t.Fatalf("Failed to set up test: %v", err)
		}

		_, err := CopyAssets(dir)
		if err == nil {
			t.Fatalf("Expected an error, got nil")
		}
	})

	t.Run("returns an error if a file cannot be written", func(t *testing.T) {
		dir := t.TempDir()

		// Pre-create a directory at the exact path one of the known
		// embedded files would be written to, so os.WriteFile fails trying
		// to write a file over an existing directory. This is coupled to
		// embedassets/assets/js/bootstrap/always-async.js existing; if that
		// file is ever renamed/removed, update this path to match.
		conflictDir := filepath.Join(dir, "__ham", "assets", "js", "bootstrap", "always-async.js")
		if err := os.MkdirAll(conflictDir, 0755); err != nil {
			t.Fatalf("Failed to set up test: %v", err)
		}

		_, err := CopyAssets(dir)
		if err == nil {
			t.Fatalf("Expected an error, got nil")
		}
	})
}
