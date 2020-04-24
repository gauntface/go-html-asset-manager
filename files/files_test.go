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

package files

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var errInjected = errors.New("injected error")

var reset func()

func TestMain(m *testing.M) {
	origFilepathWalk := filepathWalk
	origOpen := osOpen
	origCopy := ioCopy

	reset = func() {
		filepathWalk = origFilepathWalk
		osOpen = origOpen
		ioCopy = origCopy
	}

	os.Exit(m.Run())
}

func TestFind(t *testing.T) {
	tests := []struct {
		description  string
		dir          string
		exts         []string
		filepathWalk func(root string, walkFn filepath.WalkFunc) error
		want         []string
		wantError    error
	}{
		{
			description: "return error if walk fails",
			dir:         "/example",
			exts:        []string{".example"},
			filepathWalk: func(root string, walkFn filepath.WalkFunc) error {
				wantRoot := "/example"
				if root != wantRoot {
					t.Fatalf("Unexpected root; got %v, want %v", root, wantRoot)
				}
				return walkFn("", nil, errInjected)
			},
			wantError: errInjected,
			want:      []string{},
		},
		{
			description: "do nothing if the path is a directory",
			filepathWalk: func(root string, walkFn filepath.WalkFunc) error {
				return walkFn("/example", &fileInfoStub{IsDirReturn: true}, nil)
			},
			want: []string{},
		},
		{
			description: "return HTML file",
			dir:         "/example",
			exts:        []string{".html"},
			filepathWalk: func(root string, walkFn filepath.WalkFunc) error {
				return walkFn("/example/index.html", &fileInfoStub{NameReturn: "index.html"}, nil)
			},
			want: []string{
				"/example/index.html",
			},
		},
		{
			description: "return ignore file without ext",
			dir:         "/example",
			exts:        []string{".html"},
			filepathWalk: func(root string, walkFn filepath.WalkFunc) error {
				return walkFn("/example/index.css", &fileInfoStub{NameReturn: "index.css"}, nil)
			},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			defer reset()

			filepathWalk = tt.filepathWalk

			got, err := Find(tt.dir, tt.exts...)
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("Unexpected results; diff %v", diff)
			}
		})
	}
}

func Test_Hash(t *testing.T) {
	tests := []struct {
		description string
		path        string
		osOpen      func(string) (*os.File, error)
		ioCopy      func(dst io.Writer, src io.Reader) (written int64, err error)
		want        string
		wantError   error
	}{
		{
			description: "return error if open fails",
			osOpen: func(path string) (*os.File, error) {
				return nil, errInjected
			},
			wantError: errOpenFailed,
		},
		{
			description: "return error if copy fails",
			osOpen: func(path string) (*os.File, error) {
				return &os.File{}, nil
			},
			ioCopy: func(dst io.Writer, src io.Reader) (written int64, err error) {
				return 0, errInjected
			},
			wantError: errHashFailed,
		},
		{
			description: "return hash on success",
			osOpen: func(path string) (*os.File, error) {
				return &os.File{}, nil
			},
			ioCopy: func(dst io.Writer, src io.Reader) (written int64, err error) {
				return 0, nil
			},
			want: "e3b0c44",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			defer reset()

			osOpen = tt.osOpen
			ioCopy = tt.ioCopy

			got, err := Hash(tt.path)
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Different error returned; got %v, want %v", err, tt.wantError)
			}

			if err != nil {
				return
			}

			if !cmp.Equal(got, tt.want) {
				t.Errorf("Unexpected result; Diff %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

type fileInfoStub struct {
	os.FileInfo

	IsDirReturn bool
	NameReturn  string
}

func (f *fileInfoStub) IsDir() bool {
	return f.IsDirReturn
}

func (f *fileInfoStub) Name() string {
	return f.NameReturn
}
