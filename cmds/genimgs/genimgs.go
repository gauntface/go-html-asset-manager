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
	"context"
	"errors"
	"flag"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"github.com/gauntface/go-html-asset-injector/assets"
	"github.com/gauntface/go-html-asset-injector/assets/assetmanager"
	"github.com/gauntface/go-html-asset-injector/files"
	"github.com/gauntface/go-html-asset-injector/sets"
	"github.com/mitchellh/go-homedir"
	"github.com/schollz/progressbar/v3"
)

var (
	assetsdir = flag.String("assets_dir", "", "the path to a directory containing CSS and JS files")
	outputdir = flag.String("output_dir", "", "the path to a directory to generate sized images too")

	homedirExpand = homedir.Expand
)

func main() {
	ctx := context.Background()
	c, err := newClient(ctx)
	if err != nil {
		fmt.Printf("☠️ Failed to initialize new client: %v\n", err)
		os.Exit(1)
	}
	if err := c.run(ctx); err != nil {
		fmt.Printf("☠️ Run was not successful: %v\n", err)
		os.Exit(1)
	}
}

type client struct {
	staticdir string
	outputdir string

	staticManager    *assetmanager.Manager
	generatedManager *assetmanager.Manager
}

func newClient(ctx context.Context) (*client, error) {
	flag.Parse()

	if *assetsdir == "" {
		return nil, errors.New("assets_dir is required")
	}

	if *outputdir == "" {
		return nil, errors.New("output_dir is required")
	}

	absStaticDir, err := homedirExpand(*assetsdir)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path for html_dir flag: %w", err)
	}

	absOutputDir, err := homedirExpand(*outputdir)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path for html_dir flag: %w", err)
	}

	fmt.Printf("📁 Looking for Static assets in: %q\n", absStaticDir)
	fmt.Printf("📁 Will output imgs to: %q\n", absOutputDir)

	err = os.MkdirAll(absOutputDir, 0777)
	if err != nil {
		return nil, fmt.Errorf("failed to create output directory: %v", err)
	}

	staticManager, err := assetmanager.NewManager("", absStaticDir, "")
	if err != nil {
		return nil, err
	}

	generatedManager, err := assetmanager.NewManager("", absOutputDir, "")
	if err != nil {
		return nil, err
	}

	return &client{
		staticdir:        absStaticDir,
		outputdir:        absOutputDir,
		staticManager:    staticManager,
		generatedManager: generatedManager,
	}, nil
}

func (c *client) run(ctx context.Context) error {
	pngs := c.staticManager.WithType(assets.PNG)
	jpegs := c.staticManager.WithType(assets.JPEG)
	webps := c.staticManager.WithType(assets.WEBP)
	all := append(pngs, jpegs...)
	all = append(all, webps...)

	fmt.Printf("📷 Found %v images\n", len(all))

	fullImgSet, err := c.generateImageList(all)
	if err != nil {
		return err
	}

	fmt.Printf("📸 This should result in %v images\n", len(fullImgSet))

	toCreate, toDelete := c.assessAssets(fullImgSet)

	fmt.Printf("🖌️ Need to create %v images\n", len(toCreate))
	fmt.Printf("🗑️ Need to delete %v images\n", len(toDelete))

	err = c.createImages(toCreate)
	if err != nil {
		return err
	}

	err = c.deleteImages(toDelete)
	if err != nil {
		return err
	}

	fmt.Printf("✅ Done.\n")

	return nil
}

func (c *client) deleteImages(imgs []string) error {
	dirs := sets.NewStringSet()
	for _, i := range imgs {
		err := os.Remove(i)
		if err != nil {
			return err
		}
		dirs.Add(filepath.Dir(i))
	}

	for _, d := range dirs.Sorted() {
		files, err := ioutil.ReadDir(d)
		if err != nil {
			return err
		}
		if len(files) != 0 {
			continue
		}
		err = os.Remove(d)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *client) createImages(imgs []generateImage) error {
	sort.Slice(imgs, func(i, j int) bool {
		return imgs[i].outputPath < imgs[j].outputPath
	})

	bar := progressbar.NewOptions(
		len(imgs),
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionSetDescription("🖼️ Creating Images"),
		progressbar.OptionShowCount(),
		progressbar.OptionSetPredictTime(false),
	)

	jobs := make(chan generateImage, len(imgs))
	results := make(chan error, len(imgs))

	for w := 1; w <= runtime.NumCPU(); w++ {
		go imgCreatorWorker(w, jobs, results)
	}

	for _, i := range imgs {
		jobs <- i
	}

	close(jobs)

	errCount := 0
	for a := 1; a <= len(imgs); a++ {
		err := <-results
		bar.Add(1)
		if err != nil {
			fmt.Printf("failed to create image: %v\n", err)
			errCount++
		}
	}

	if errCount > 0 {
		return fmt.Errorf("%v errors occured while creating images", errCount)
	}

	return nil
}

func (c *client) assessAssets(allImages []generateImage) ([]generateImage, []string) {
	generatedPNGs := c.generatedManager.WithType(assets.PNG)
	generatedJPEGs := c.generatedManager.WithType(assets.JPEG)
	generatedWEBPs := c.generatedManager.WithType(assets.WEBP)
	allGenerated := append(generatedPNGs, generatedJPEGs...)
	allGenerated = append(allGenerated, generatedWEBPs...)

	generatedSet := sets.NewStringSet()
	for _, g := range allGenerated {
		img := g.(*assetmanager.LocalAsset)
		generatedSet.Add(img.Path())
	}

	requiredMap := map[string]generateImage{}
	for _, i := range allImages {
		requiredMap[i.outputPath] = i
	}

	imgsToGenerate := []generateImage{}
	for path, r := range requiredMap {
		if generatedSet.Contains(path) {
			continue
		}
		imgsToGenerate = append(imgsToGenerate, r)
	}

	filesToRm := []string{}
	for _, g := range generatedSet.Slice() {
		if _, ok := requiredMap[g]; !ok {
			filesToRm = append(filesToRm, g)
		}
	}

	return imgsToGenerate, filesToRm
}

func (c *client) generateImageList(imgs []assetmanager.Asset) ([]generateImage, error) {
	genImgs := []generateImage{}
	errs := []error{}

	var wg sync.WaitGroup
	var imgMu sync.Mutex
	var errMu sync.Mutex

	wg.Add(len(imgs))

	for _, i := range imgs {
		go func(i assetmanager.Asset) {
			defer wg.Done()

			img := i.(*assetmanager.LocalAsset)
			// Ignore an image if it's in the generated path
			imgPath := img.Path()
			if strings.HasPrefix(imgPath, c.outputdir) {
				return
			}

			gi, err := c.generateImageSet(imgPath)
			if err != nil {
				errMu.Lock()
				defer errMu.Unlock()
				errs = append(errs, err)
				return
			}

			imgMu.Lock()
			defer imgMu.Unlock()
			genImgs = append(genImgs, gi...)
		}(i)
	}

	wg.Wait()

	if len(errs) > 0 {
		fmt.Printf("☠️ %v errors occurred while generating the image list:", len(errs))
		for i, e := range errs {
			fmt.Printf("    - %v) %v", i, e)
		}
		return nil, fmt.Errorf("%v errors occurred while generating image list. See logs for details.", len(errs))
	}

	return genImgs, nil
}

func (c *client) generateImageSet(imgPath string) ([]generateImage, error) {
	srcImg, err := imaging.Open(imgPath)
	if err != nil {
		return nil, err
	}

	outputDir, err := c.generatedDir(imgPath)
	if err != nil {
		return nil, err
	}

	sizes := generateSizes(srcImg)

	genImgs := []generateImage{}
	for _, s := range sizes {
		ext := filepath.Ext(imgPath)
		resizedImg := path.Join(outputDir, fmt.Sprintf("%v%v", s, ext))
		genImgs = append(genImgs,
			generateImage{
				originalPath: imgPath,
				width:        s,
				outputPath:   resizedImg,
			},
			generateImage{
				originalPath: imgPath,
				width:        s,
				outputPath:   path.Join(outputDir, fmt.Sprintf("%v%v", s, ".webp")),
			},
		)
	}

	return genImgs, nil
}

func (c *client) generatedDir(imgPath string) (string, error) {
	hash, err := files.Hash(imgPath)
	if err != nil {
		return "", err
	}
	ext := filepath.Ext(imgPath)
	fullFilename := path.Base(imgPath)
	filename := fullFilename[:len(fullFilename)-len(ext)]
	return filepath.Join(c.outputdir, fmt.Sprintf("%v.%v", filename, hash)), nil
}

func generateSizes(img image.Image) []int {
	origSize := img.Bounds().Size()

	widths := []int{}
	currentWidth := 400
	interval := 200
	for {
		if currentWidth >= origSize.X {
			break
		}
		widths = append(widths, currentWidth)
		currentWidth += interval
	}
	widths = append(widths, origSize.X)
	return widths
}

func imgCreatorWorker(id int, jobs <-chan generateImage, results chan<- error) {
	for j := range jobs {
		err := createImage(j)
		results <- err
	}
}

func createImage(img generateImage) error {
	err := os.MkdirAll(filepath.Dir(img.outputPath), 0777)
	if err != nil {
		return err
	}

	ext := filepath.Ext(img.outputPath)
	switch ext {
	case ".png", ".jpg", ".jpeg":
		return createImagingImage(img)
	case ".webp":
		return createWebpImage(img)
	default:
		return fmt.Errorf("unsupported file: %q with extension%q", img.outputPath, ext)
	}
}

func createImagingImage(img generateImage) error {
	srcImg, err := imaging.Open(img.originalPath)
	if err != nil {
		return err
	}

	dst := imaging.Resize(srcImg, img.width, 0, imaging.Lanczos)
	err = imaging.Save(dst, img.outputPath)
	if err != nil {
		return err
	}
	return nil
}

func createWebpImage(img generateImage) error {
	srcImg, err := imaging.Open(img.originalPath)
	if err != nil {
		return err
	}

	dst := imaging.Resize(srcImg, img.width, 0, imaging.Lanczos)

	f, err := os.Create(img.outputPath)
	if err != nil {
		return err
	}

	err = webp.Encode(f, dst, nil)
	if err != nil {
		return err
	}
	return nil
}

type generateImage struct {
	originalPath string
	width        int
	outputPath   string
}
