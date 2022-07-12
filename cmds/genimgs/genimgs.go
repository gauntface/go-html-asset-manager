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

	"github.com/Kagami/go-avif"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	s3manager "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	awstypes "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"github.com/gauntface/go-html-asset-manager/v3/assets"
	"github.com/gauntface/go-html-asset-manager/v3/assets/assetmanager"
	"github.com/gauntface/go-html-asset-manager/v3/utils/config"
	"github.com/gauntface/go-html-asset-manager/v3/utils/files"
	"github.com/gauntface/go-html-asset-manager/v3/utils/sets"
	"github.com/mitchellh/go-homedir"
	"github.com/schollz/progressbar/v3"
)

var (
	configPath      = flag.String("config", "asset-manager.json", "The path of the Config file.")
	cacheControlAge = flag.Int64("cache_control", 31104000, "The max age for caching images")
	homedirExpand   = homedir.Expand
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
	staticdir   string
	outputdir   string
	s3Bucket    string
	s3BucketDir string
	maxWidth    int64

	staticManager    *assetmanager.Manager
	generatedManager *assetmanager.Manager
	s3               *s3.Client
	s3Manager        *s3manager.Uploader
}

func newClient(ctx context.Context) (*client, error) {
	flag.Parse()

	absConfigPath, err := homedirExpand(*configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path for config flag: %w", err)
	}
	fmt.Printf("📁 Getting config file: %q\n", absConfigPath)

	c, err := config.Get(absConfigPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path for html_dir flag: %w", err)
	}

	fmt.Printf("📁 Looking for Static assets in: %q\n", c.GenAssets.StaticDir)
	fmt.Printf("📁 Will output imgs to: [🪣 %v]/%v\n", c.GenAssets.OutputBucket, c.GenAssets.OutputBucketDir)

	err = os.MkdirAll(c.GenAssets.OutputDir, 0777)
	if err != nil {
		return nil, fmt.Errorf("failed to create output directory: %v", err)
	}

	staticManager, err := assetmanager.NewManager("", c.GenAssets.StaticDir, "")
	if err != nil {
		return nil, err
	}

	generatedManager, err := assetmanager.NewManager("", c.GenAssets.OutputDir, "")
	if err != nil {
		return nil, err
	}

	maxWidth := c.GenAssets.MaxWidth * c.GenAssets.MaxDensity
	fmt.Printf("📏 Max width will be %v (CSS px) x %v (Density) = %v\n", c.GenAssets.MaxWidth, c.GenAssets.MaxDensity, maxWidth)

	cfg, err := awsconfig.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config, %v", err)
	}

	s3Client := s3.NewFromConfig(cfg)
	s3Manager := s3manager.NewUploader(s3Client)

	return &client{
		staticdir:        c.GenAssets.StaticDir,
		outputdir:        c.GenAssets.OutputDir,
		s3Bucket:         c.GenAssets.OutputBucket,
		s3BucketDir:      c.GenAssets.OutputBucketDir,
		maxWidth:         maxWidth,
		staticManager:    staticManager,
		generatedManager: generatedManager,
		s3:               s3Client,
		s3Manager:        s3Manager,
	}, nil
}

func (c *client) run(ctx context.Context) error {
	pngs := c.staticManager.WithType(assets.PNG)
	jpegs := c.staticManager.WithType(assets.JPEG)
	webps := c.staticManager.WithType(assets.WEBP)
	avifs := c.staticManager.WithType(assets.AVIF)
	all := append(pngs, jpegs...)
	all = append(all, webps...)
	all = append(all, avifs...)

	fmt.Printf("📷 Found %v images\n", len(all))

	fullImgSet, err := c.generateImageList(all)
	if err != nil {
		return err
	}

	fmt.Printf("📸 This should result in %v images\n", len(fullImgSet))

	s3Imgs, err := c.getS3GenImages(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("🪣 S3 has %v images\n", len(s3Imgs))

	toCreate, toDelete := c.assessAssets(ctx, fullImgSet, s3Imgs)
	if err != nil {
		return err
	}

	fmt.Printf("🖌️ Need to create %v images\n", len(toCreate))
	fmt.Printf("🗑️ Need to delete %v images\n", len(toDelete))

	err = c.createImages(toCreate)
	if err != nil {
		return err
	}

	/* err = c.deleteImages(toDelete)
	if err != nil {
		return err
	}*/

	fmt.Printf("✅ Done.\n")

	return nil
}

func (c *client) getS3GenImages(ctx context.Context) ([]awstypes.Object, error) {
	params := &s3.ListObjectsV2Input{
		Bucket: &c.s3Bucket,
		Prefix: &c.s3BucketDir,
	}

	// Create the Paginator for the ListObjectsV2 operation.
	p := s3.NewListObjectsV2Paginator(c.s3, params)

	// Iterate through the S3 object pages, printing each object returned.
	objs := []awstypes.Object{}
	for p.HasMorePages() {
		// Next Page takes a new context for each page retrieval. This is where
		// you could add timeouts or deadlines.
		page, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		// Log the objects found
		objs = append(objs, page.Contents...)
	}
	return objs, nil
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
		go c.imgCreatorWorker(w, jobs, results)
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
			fmt.Printf("\nfailed to create image: %v\n", err)
			errCount++
		}
	}

	if errCount > 0 {
		return fmt.Errorf("%v errors occured while creating images", errCount)
	}

	return nil
}

func (c *client) assessAssets(ctx context.Context, allImages []generateImage, s3Images []types.Object) ([]generateImage, []string) {
	requiredMap := map[string]generateImage{}
	for _, i := range allImages {
		k := strings.TrimPrefix(i.outputPath, c.staticdir)
		k = strings.TrimPrefix(k, "/")
		requiredMap[k] = i
	}

	imgsToGenerate := []generateImage{}
	for k, r := range requiredMap {
		found := false
		for _, i := range s3Images {
			if k == *i.Key {
				found = true
				break
			}
		}
		if !found {
			imgsToGenerate = append(imgsToGenerate, r)
		}
	}

	filesToRm := []string{}
	for _, g := range s3Images {
		if _, ok := requiredMap[*g.Key]; !ok {
			filesToRm = append(filesToRm, *g.Key)
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
		fmt.Printf("☠️ %v errors occurred while generating the image list:\n", len(errs))
		for i, e := range errs {
			fmt.Printf("    - %v) %v\n", i, e)
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

	sizes := generateSizes(srcImg, c.maxWidth)

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

		// avif library doesn't support alpha channel
		if !strings.HasSuffix(imgPath, ".png") {
			genImgs = append(genImgs, generateImage{
				originalPath: imgPath,
				width:        s,
				outputPath:   path.Join(outputDir, fmt.Sprintf("%v%v", s, ".avif")),
			})
		}
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

func generateSizes(img image.Image, maxWidth int64) []int {
	origSize := img.Bounds().Size()

	widths := []int{}
	currentWidth := 400
	interval := 200
	for {
		if currentWidth >= origSize.X || int64(currentWidth) > maxWidth {
			break
		}
		widths = append(widths, currentWidth)
		currentWidth += interval
	}

	if int64(origSize.X) <= maxWidth {
		widths = append(widths, origSize.X)
	}
	return widths
}

func (c *client) imgCreatorWorker(id int, jobs <-chan generateImage, results chan<- error) {
	ctx := context.Background()
	for j := range jobs {
		err := c.createAndUploadImage(ctx, j)
		results <- err
	}
}

func (c *client) createAndUploadImage(ctx context.Context, img generateImage) error {
	err := createImage(img)
	if err != nil {
		return err
	}
	return c.uploadImage(ctx, img)
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
	case ".avif":
		return createAvifImage(img)
	default:
		return fmt.Errorf("unsupported file: %q with extension%q", img.outputPath, ext)
	}
}

func (c *client) uploadImage(ctx context.Context, img generateImage) error {
	f, err := os.Open(img.outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	key := strings.TrimPrefix(img.outputPath, c.staticdir)
	key = strings.TrimPrefix(key, "/")

	cc := fmt.Sprintf("max-age=%v", cacheControlAge)

	_, err = c.s3Manager.Upload(ctx, &s3.PutObjectInput{
		Bucket:       &c.s3Bucket,
		ACL:          awstypes.ObjectCannedACLPublicRead,
		CacheControl: &cc,
		Key:          &key,
		Body:         f,
	})
	return err
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

func createAvifImage(img generateImage) error {
	// go-avif doesn't support alpha channel
	if path.Ext(img.originalPath) == ".png" {
		return nil
	}
	srcImg, err := imaging.Open(img.originalPath)
	if err != nil {
		return err
	}

	dst := imaging.Resize(srcImg, img.width, 0, imaging.Lanczos)

	f, err := os.Create(img.outputPath)
	if err != nil {
		return err
	}

	err = avif.Encode(f, dst, nil)
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
