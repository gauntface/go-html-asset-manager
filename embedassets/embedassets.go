package embedassets

import (
	"embed"
	"errors"
	"fmt"
	"os"
	"path"
)

var (
	//go:embed assets
	assetsfs embed.FS

	errReadFailed    = errors.New("failed to read directory")
	errMakeDirFailed = errors.New("failed to make directory")
	errWriteFailed   = errors.New("failed to write file")
)

func CopyAssets(staticDir string) ([]string, error) {
	outputDir := path.Join(staticDir, "__ham")

	dirs := []string{
		"assets",
	}
	var currentDir string
	files := []string{}
	for len(dirs) > 0 {
		currentDir, dirs = dirs[0], dirs[1:]

		dirContents, err := assetsfs.ReadDir(currentDir)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", errReadFailed, err)
		}

		for _, d := range dirContents {
			if d.IsDir() {
				dirs = append(dirs, path.Join(currentDir, d.Name()))
				continue
			}

			err := os.MkdirAll(path.Join(outputDir, currentDir), 0755)
			if err != nil {
				return nil, fmt.Errorf("%w: %v", errMakeDirFailed, err)
			}

			data, err := assetsfs.ReadFile(path.Join(currentDir, d.Name()))
			if err != nil {
				return nil, fmt.Errorf("%w: %v", errReadFailed, err)
			}

			fp := path.Join(outputDir, currentDir, d.Name())
			err = os.WriteFile(fp, data, 0755)
			if err != nil {
				return nil, fmt.Errorf("%w: %v", errWriteFailed, err)
			}

			files = append(files, fp)
		}
	}
	return files, nil
}
