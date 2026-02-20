package genimgs

import (
	"context"
	"fmt"
	"image"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/disintegration/imaging"
	"github.com/gauntface/go-html-asset-manager/v5/utils/config"
	"golang.org/x/sync/singleflight"
)

type mockS3Client struct {
	S3ClientInterface
	callCount int32
	delay     time.Duration
	mu        sync.Mutex
	active    int32
	maxActive int32
}

func (m *mockS3Client) ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	atomic.AddInt32(&m.callCount, 1)
	currActive := atomic.AddInt32(&m.active, 1)
	defer atomic.AddInt32(&m.active, -1)

	m.mu.Lock()
	if currActive > m.maxActive {
		m.maxActive = currActive
	}
	m.mu.Unlock()

	if m.delay > 0 {
		time.Sleep(m.delay)
	}

	key := fmt.Sprintf("%v/100.webp", *params.Prefix)
	return &s3.ListObjectsV2Output{
		Contents: []types.Object{
			{
				Key: &key,
			},
		},
		IsTruncated: nil,
	}, nil
}

func TestLookupSizes_CachingAndConcurrency(t *testing.T) {
	// Mock external dependencies
	oldImagingOpen := imagingOpen
	oldFilesHash := filesHash
	defer func() {
		imagingOpen = oldImagingOpen
		filesHash = oldFilesHash
	}()

	imagingOpen = func(path string, opts ...imaging.DecodeOption) (image.Image, error) {
		return image.NewRGBA(image.Rect(0, 0, 1000, 1000)), nil
	}
	filesHash = func(path string) (string, error) {
		return "mockhash", nil
	}

	conf := &config.Config{
		Assets: &config.AssetsConfig{
			StaticDir: "/static",
		},
		GenAssets: &config.GeneratedImagesConfig{
			StaticDir:       "/static",
			OutputDir:       "/output",
			OutputBucket:    "bucket",
			OutputBucketDir: "gen",
			MaxWidth:        1000,
			MaxDensity:      1,
		},
	}

	m := &mockS3Client{
		delay: 50 * time.Millisecond,
	}

	// Reset global state
	s3Cache = sync.Map{}
	s3Group = singleflight.Group{}

	const numCalls = 10
	var wg sync.WaitGroup
	wg.Add(numCalls)

	imgPath := "test.png"

	for i := 0; i < numCalls; i++ {
		go func() {
			defer wg.Done()
			_, err := LookupSizes(m, conf, imgPath)
			if err != nil {
				t.Errorf("LookupSizes failed: %v", err)
			}
		}()
	}

	wg.Wait()

	if m.callCount != 1 {
		t.Errorf("Expected 1 S3 call due to caching/singleflight, got %v", m.callCount)
	}

	// Test cache hit
	_, err := LookupSizes(m, conf, imgPath)
	if err != nil {
		t.Fatalf("LookupSizes failed: %v", err)
	}
	if m.callCount != 1 {
		t.Errorf("Expected callCount to remain 1 after cache hit, got %v", m.callCount)
	}
}

func TestLookupSizes_Semaphore(t *testing.T) {
	// Mock external dependencies
	oldImagingOpen := imagingOpen
	oldFilesHash := filesHash
	defer func() {
		imagingOpen = oldImagingOpen
		filesHash = oldFilesHash
	}()

	imagingOpen = func(path string, opts ...imaging.DecodeOption) (image.Image, error) {
		return image.NewRGBA(image.Rect(0, 0, 1000, 1000)), nil
	}
	// Return a different hash for each path to avoid singleflight
	filesHash = func(path string) (string, error) {
		return "hash-" + path, nil
	}

	conf := &config.Config{
		Assets: &config.AssetsConfig{
			StaticDir: "/static",
		},
		GenAssets: &config.GeneratedImagesConfig{
			StaticDir:       "/static",
			OutputDir:       "/static/output", // localDirPath must be relative to StaticDir for filepath.Rel to work easily without ".."
			OutputBucket:    "bucket",
			OutputBucketDir: "gen",
			MaxWidth:        1000,
			MaxDensity:      1,
		},
	}

	m := &mockS3Client{
		delay: 100 * time.Millisecond,
	}

	// Reset global state
	s3Cache = sync.Map{}
	s3Group = singleflight.Group{}

	const numCalls = 5
	var wg sync.WaitGroup
	wg.Add(numCalls)

	for i := 0; i < numCalls; i++ {
		imgPath := fmt.Sprintf("test-%d.png", i)
		go func(p string) {
			defer wg.Done()
			_, err := LookupSizes(m, conf, p)
			if err != nil {
				t.Errorf("LookupSizes failed: %v", err)
			}
		}(imgPath)
	}

	wg.Wait()

	if m.callCount != numCalls {
		t.Errorf("Expected %d S3 calls, got %v", numCalls, m.callCount)
	}

	if m.maxActive > maxS3ParallelRequests {
		t.Errorf("Expected max %v concurrent S3 calls, got %v", maxS3ParallelRequests, m.maxActive)
	}
}
