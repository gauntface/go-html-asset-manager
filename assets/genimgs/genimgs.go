package genimgs

import (
	"errors"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/gauntface/go-html-asset-manager/utils/config"
	"github.com/gauntface/go-html-asset-manager/utils/files"
)

var (
	errFileHash = errors.New("failed to get file hash")
	errRelPath  = errors.New("unable to get relative path")

	imagingOpen   = imaging.Open
	filesHash     = files.Hash
	ioutilReadDir = ioutil.ReadDir
)

func getPath(conf *config.Config, imgPath string) string {
	return filepath.Join(conf.Assets.StaticDir, imgPath)
}

func Open(conf *config.Config, imgPath string) (image.Image, error) {
	return imagingOpen(getPath(conf, imgPath))
}

func LookupSizes(conf *config.Config, imgPath string) ([]GenImg, error) {
	srcPath := getPath(conf, imgPath)

	hash, err := filesHash(srcPath)
	if err != nil {
		return nil, fmt.Errorf("%w for img %q", errFileHash, srcPath)
	}

	// Get available sizes of the image
	sizes, err := getImageSizes(conf, srcPath, hash)
	if err != nil {
		return nil, err
	}
	return sizes, nil
}

func getImageSizes(conf *config.Config, srcPath, hash string) ([]GenImg, error) {
	filename := strings.TrimSuffix(filepath.Base(srcPath), filepath.Ext(srcPath))
	genDirName := fmt.Sprintf("%v.%v", filename, hash)
	dirPath := filepath.Join(conf.GenAssets.OutputDir, genDirName)

	contents, err := ioutilReadDir(dirPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}

	generatedDirURL, err := filepath.Rel(conf.Assets.StaticDir, dirPath)
	if err != nil {
		return nil, fmt.Errorf("%w from %q to %q: %v", errRelPath, conf.GenAssets.OutputDir, dirPath, err)
	}

	maxSize := conf.GenAssets.MaxWidth * conf.GenAssets.MaxDensity
	imgs := []GenImg{}
	for _, c := range contents {
		if c.IsDir() {
			continue
		}

		ext := filepath.Ext(c.Name())
		filename := strings.TrimSuffix(c.Name(), ext)
		size, err := strconv.ParseInt(filename, 10, 64)
		if err != nil {
			continue
		}

		if size > maxSize {
			continue
		}

		var typ string
		switch ext {
		case ".webp":
			typ = "image/webp"
			break
		}

		imgs = append(imgs, GenImg{
			URL:  filepath.Join("/", generatedDirURL, c.Name()),
			Type: typ,
			Size: size,
		})
	}

	return imgs, nil
}

func GroupByType(imgs []GenImg) map[string][]GenImg {
	// Group by type
	sourceSetByType := map[string][]GenImg{}
	for _, s := range imgs {
		_, ok := sourceSetByType[s.Type]
		if !ok {
			sourceSetByType[s.Type] = []GenImg{}
		}

		sourceSetByType[s.Type] = append(sourceSetByType[s.Type], s)
	}
	for t := range sourceSetByType {
		sort.Slice(sourceSetByType[t], func(i, j int) bool {
			return sourceSetByType[t][i].Size < sourceSetByType[t][j].Size
		})
	}
	return sourceSetByType
}

type GenImg struct {
	URL  string
	Type string
	Size int64
}
