package genimgs

import (
	"context"
	"errors"
	"fmt"
	"image"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	awstypes "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/disintegration/imaging"
	"github.com/gauntface/go-html-asset-manager/v3/utils/config"
	"github.com/gauntface/go-html-asset-manager/v3/utils/files"
)

var (
	errFileHash = errors.New("failed to get file hash")
	errRelPath  = errors.New("unable to get relative path")

	imagingOpen = imaging.Open
	filesHash   = files.Hash
)

func getPath(conf *config.Config, imgPath string) string {
	return filepath.Join(conf.Assets.StaticDir, imgPath)
}

func Open(conf *config.Config, imgPath string) (image.Image, error) {
	return imagingOpen(getPath(conf, imgPath))
}

func LookupSizes(s3Client *s3.Client, conf *config.Config, imgPath string) ([]GenImg, error) {
	srcPath := getPath(conf, imgPath)

	hash, err := filesHash(srcPath)
	if err != nil {
		return nil, fmt.Errorf("%w for img %q", errFileHash, srcPath)
	}

	// Get available sizes of the image
	sizes, err := getImageSizes(s3Client, conf, srcPath, hash)
	if err != nil {
		return nil, err
	}
	return sizes, nil
}

func getImageSizes(s3Client *s3.Client, conf *config.Config, srcPath, hash string) ([]GenImg, error) {
	filename := strings.TrimSuffix(filepath.Base(srcPath), filepath.Ext(srcPath))
	genDirName := fmt.Sprintf("%v.%v", filename, hash)
	localDirPath := filepath.Join(conf.GenAssets.OutputDir, genDirName)
	bucketDirPath := filepath.Join(conf.GenAssets.OutputBucketDir, genDirName)

	objs, err := lookupS3Images(s3Client, conf, bucketDirPath)
	if err != nil {
		return nil, fmt.Errorf("unable to lookup S3 images in %v", bucketDirPath)
	}

	maxSize := conf.GenAssets.MaxWidth * conf.GenAssets.MaxDensity
	generatedDirURL, err := filepath.Rel(conf.GenAssets.StaticDir, localDirPath)
	if err != nil {
		fmt.Printf("GENIMGS\nStaticDir: %v\nDirPath: %v\nRel: %v\n\n", conf.GenAssets.StaticDir, localDirPath, generatedDirURL)
		return nil, fmt.Errorf("%w from %q to %q: %v", errRelPath, conf.GenAssets.StaticDir, localDirPath, err)
	}

	imgs := []GenImg{}
	for _, c := range objs {
		_, file := filepath.Split(*c.Key)
		ext := filepath.Ext(file)
		filename := strings.TrimSuffix(file, ext)

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
		case ".avif":
			typ = "image/avif"
		}

		imgs = append(imgs, GenImg{
			URL:  filepath.Join("/", generatedDirURL, file),
			Type: typ,
			Size: size,
		})
	}

	return imgs, nil
}

func lookupS3Images(s3Client *s3.Client, conf *config.Config, dir string) ([]awstypes.Object, error) {
	ctx := context.Background()
	params := &s3.ListObjectsV2Input{
		Bucket: &conf.GenAssets.OutputBucket,
		Prefix: &dir,
	}

	// Create the Paginator for the ListObjectsV2 operation.
	p := s3.NewListObjectsV2Paginator(s3Client, params)

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
