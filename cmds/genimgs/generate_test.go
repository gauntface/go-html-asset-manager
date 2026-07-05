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
	"image"
	"sort"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func Test_generateSizes(t *testing.T) {
	tests := []struct {
		description string
		width       int
		height      int
		maxWidth    int64
		want        []int
	}{
		{
			description: "returns just the original size if it's already narrower than the first step",
			width:       300,
			height:      200,
			maxWidth:    700,
			want:        []int{300},
		},
		{
			description: "steps up in 200px increments below the original width",
			width:       900,
			height:      600,
			maxWidth:    2000,
			want:        []int{400, 600, 800, 900},
		},
		{
			description: "stops stepping at maxWidth and does not include the original size if it exceeds maxWidth",
			width:       2000,
			height:      1000,
			maxWidth:    900,
			want:        []int{400, 600, 800},
		},
		{
			description: "includes the original size exactly at maxWidth",
			width:       900,
			height:      600,
			maxWidth:    900,
			want:        []int{400, 600, 800, 900},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			img := image.NewRGBA(image.Rect(0, 0, tt.width, tt.height))
			got := generateSizes(img, tt.maxWidth)

			if len(got) != len(tt.want) {
				t.Fatalf("Unexpected sizes; got %v, want %v", got, tt.want)
			}
			for i, w := range tt.want {
				if got[i] != w {
					t.Fatalf("Unexpected size at %v; got %v, want %v (full: got %v, want %v)", i, got[i], w, got, tt.want)
				}
			}
		})
	}
}

func Test_assessAssets(t *testing.T) {
	strPtr := func(s string) *string { return &s }

	tests := []struct {
		description string
		staticdir   string
		allImages   []generateImage
		s3Images    []types.Object
		wantCreate  []string
		wantDelete  []string
	}{
		{
			description: "generates everything if S3 has no images",
			staticdir:   "/static/",
			allImages: []generateImage{
				{outputPath: "/static/img/a.webp"},
				{outputPath: "/static/img/b.webp"},
			},
			wantCreate: []string{"/static/img/a.webp", "/static/img/b.webp"},
		},
		{
			description: "does not recreate images already in S3",
			staticdir:   "/static/",
			allImages: []generateImage{
				{outputPath: "/static/img/a.webp"},
				{outputPath: "/static/img/b.webp"},
			},
			s3Images: []types.Object{
				{Key: strPtr("img/a.webp")},
			},
			wantCreate: []string{"/static/img/b.webp"},
		},
		{
			description: "deletes S3 images no longer required locally",
			staticdir:   "/static/",
			allImages: []generateImage{
				{outputPath: "/static/img/a.webp"},
			},
			s3Images: []types.Object{
				{Key: strPtr("img/a.webp")},
				{Key: strPtr("img/stale.webp")},
			},
			wantDelete: []string{"img/stale.webp"},
		},
		{
			description: "does nothing if local images and S3 already match",
			staticdir:   "/static/",
			allImages: []generateImage{
				{outputPath: "/static/img/a.webp"},
			},
			s3Images: []types.Object{
				{Key: strPtr("img/a.webp")},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			c := &client{staticdir: tt.staticdir}
			toCreate, toDelete := c.assessAssets(context.Background(), tt.allImages, tt.s3Images)

			gotCreate := []string{}
			for _, g := range toCreate {
				gotCreate = append(gotCreate, g.outputPath)
			}
			sort.Strings(gotCreate)
			sort.Strings(tt.wantCreate)

			if len(gotCreate) != len(tt.wantCreate) {
				t.Fatalf("Unexpected toCreate; got %v, want %v", gotCreate, tt.wantCreate)
			}
			for i, w := range tt.wantCreate {
				if gotCreate[i] != w {
					t.Fatalf("Unexpected toCreate; got %v, want %v", gotCreate, tt.wantCreate)
				}
			}

			sort.Strings(toDelete)
			sort.Strings(tt.wantDelete)
			if len(toDelete) != len(tt.wantDelete) {
				t.Fatalf("Unexpected toDelete; got %v, want %v", toDelete, tt.wantDelete)
			}
			for i, w := range tt.wantDelete {
				if toDelete[i] != w {
					t.Fatalf("Unexpected toDelete; got %v, want %v", toDelete, tt.wantDelete)
				}
			}
		})
	}
}
