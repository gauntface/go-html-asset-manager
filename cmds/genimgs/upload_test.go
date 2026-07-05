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
	"os"
	"path/filepath"
	"testing"

	"github.com/aws/aws-sdk-go-v2/feature/s3/transfermanager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"golang.org/x/sync/semaphore"
)

// fakeS3APIClient implements transfermanager.S3APIClient. Only PutObject is
// expected to be called for the small test files uploaded here; the other
// methods exist purely to satisfy the interface.
type fakeS3APIClient struct {
	putObjectCalls []*s3.PutObjectInput
	putObjectErr   error
}

func (f *fakeS3APIClient) PutObject(ctx context.Context, in *s3.PutObjectInput, opts ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	f.putObjectCalls = append(f.putObjectCalls, in)
	return &s3.PutObjectOutput{}, f.putObjectErr
}

func (f *fakeS3APIClient) UploadPart(ctx context.Context, in *s3.UploadPartInput, opts ...func(*s3.Options)) (*s3.UploadPartOutput, error) {
	return nil, errors.New("UploadPart: not implemented by fake")
}

func (f *fakeS3APIClient) CreateMultipartUpload(ctx context.Context, in *s3.CreateMultipartUploadInput, opts ...func(*s3.Options)) (*s3.CreateMultipartUploadOutput, error) {
	return nil, errors.New("CreateMultipartUpload: not implemented by fake")
}

func (f *fakeS3APIClient) CompleteMultipartUpload(ctx context.Context, in *s3.CompleteMultipartUploadInput, opts ...func(*s3.Options)) (*s3.CompleteMultipartUploadOutput, error) {
	return nil, errors.New("CompleteMultipartUpload: not implemented by fake")
}

func (f *fakeS3APIClient) AbortMultipartUpload(ctx context.Context, in *s3.AbortMultipartUploadInput, opts ...func(*s3.Options)) (*s3.AbortMultipartUploadOutput, error) {
	return nil, errors.New("AbortMultipartUpload: not implemented by fake")
}

func (f *fakeS3APIClient) GetObject(ctx context.Context, in *s3.GetObjectInput, opts ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	return nil, errors.New("GetObject: not implemented by fake")
}

func (f *fakeS3APIClient) HeadObject(ctx context.Context, in *s3.HeadObjectInput, opts ...func(*s3.Options)) (*s3.HeadObjectOutput, error) {
	return nil, errors.New("HeadObject: not implemented by fake")
}

func (f *fakeS3APIClient) ListObjectsV2(ctx context.Context, in *s3.ListObjectsV2Input, opts ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	return nil, errors.New("ListObjectsV2: not implemented by fake")
}

func Test_uploadImage(t *testing.T) {
	t.Run("uploads the file with the resolved cache-control header and S3 key", func(t *testing.T) {
		dir := t.TempDir()
		outputPath := filepath.Join(dir, "static", "img", "a.webp")
		if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
			t.Fatalf("Failed to set up test: %v", err)
		}
		if err := os.WriteFile(outputPath, []byte("fake image bytes"), 0644); err != nil {
			t.Fatalf("Failed to set up test: %v", err)
		}

		fake := &fakeS3APIClient{}
		c := &client{
			staticdir: filepath.Join(dir, "static") + string(filepath.Separator),
			s3Bucket:  "example-bucket",
			s3Manager: transfermanager.New(fake),
			s3Sem:     semaphore.NewWeighted(1),
		}

		if err := c.uploadImage(context.Background(), generateImage{outputPath: outputPath}); err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if len(fake.putObjectCalls) != 1 {
			t.Fatalf("Expected exactly one PutObject call, got %v", len(fake.putObjectCalls))
		}

		got := fake.putObjectCalls[0]
		if got.Bucket == nil || *got.Bucket != "example-bucket" {
			t.Fatalf("Unexpected bucket: %v", got.Bucket)
		}
		if got.Key == nil || *got.Key != "img/a.webp" {
			t.Fatalf("Unexpected key: %v", got.Key)
		}
		wantCacheControl := cacheControlHeader(*cacheControlAge)
		if got.CacheControl == nil || *got.CacheControl != wantCacheControl {
			t.Fatalf("Unexpected Cache-Control header; got %v, want %v", got.CacheControl, wantCacheControl)
		}
	})

	t.Run("returns an error if the upload fails", func(t *testing.T) {
		dir := t.TempDir()
		outputPath := filepath.Join(dir, "a.webp")
		if err := os.WriteFile(outputPath, []byte("fake image bytes"), 0644); err != nil {
			t.Fatalf("Failed to set up test: %v", err)
		}

		errInjected := errors.New("injected upload error")
		fake := &fakeS3APIClient{putObjectErr: errInjected}
		c := &client{
			staticdir: dir + string(filepath.Separator),
			s3Bucket:  "example-bucket",
			s3Manager: transfermanager.New(fake),
			s3Sem:     semaphore.NewWeighted(1),
		}

		err := c.uploadImage(context.Background(), generateImage{outputPath: outputPath})
		if !errors.Is(err, errInjected) {
			t.Fatalf("Unexpected error; got %v, want %v", err, errInjected)
		}
	})

	t.Run("returns an error if the file cannot be opened", func(t *testing.T) {
		fake := &fakeS3APIClient{}
		c := &client{
			s3Manager: transfermanager.New(fake),
			s3Sem:     semaphore.NewWeighted(1),
		}

		err := c.uploadImage(context.Background(), generateImage{outputPath: "/does/not/exist.webp"})
		if err == nil {
			t.Fatalf("Expected an error, got nil")
		}
		if len(fake.putObjectCalls) != 0 {
			t.Fatalf("Expected no PutObject calls, got %v", len(fake.putObjectCalls))
		}
	})
}
