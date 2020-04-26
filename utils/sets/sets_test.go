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

package sets

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_StringSet_Add(t *testing.T) {
	tests := []struct {
		description string
		input       string
		want        StringSet
	}{
		{
			description: "add string to set",
			input:       "example",
			want: StringSet{
				"example": true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			ss := StringSet{}
			ss.Add(tt.input)
			ss.Add(tt.input)

			if diff := cmp.Diff(ss, tt.want); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}

func Test_StringSet_Merge(t *testing.T) {
	tests := []struct {
		description string
		start       StringSet
		input       StringSet
		want        StringSet
	}{
		{
			description: "merge empty string set",
			start:       StringSet{},
			input:       StringSet{},
			want:        StringSet{},
		},
		{
			description: "merge string set with single item",
			start:       StringSet{},
			input: StringSet{
				"example": true,
			},
			want: StringSet{
				"example": true,
			},
		},
		{
			description: "merge string set with multiple items",
			start: StringSet{
				"start-unique-example": true,
				"start-shared-example": true,
			},
			input: StringSet{
				"start-shared-example": true,
				"input-example":        true,
			},
			want: StringSet{
				"start-unique-example": true,
				"start-shared-example": true,
				"input-example":        true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			tt.start.Merge(tt.input)

			if diff := cmp.Diff(tt.start, tt.want); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}

func Test_StringSet_Slice(t *testing.T) {
	tests := []struct {
		description string
		input       StringSet
		want        []string
	}{
		{
			description: "return empty slice",
			input:       StringSet{},
			want:        []string{},
		},
		{
			description: "return slice of set of single item",
			input: StringSet{
				"example": true,
			},
			want: []string{"example"},
		},
		{
			description: "return slice of set of multiple items",
			input: StringSet{
				"example":   true,
				"example-2": true,
			},
			want: []string{"example", "example-2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := tt.input.Slice()

			sort.Strings(got)
			sort.Strings(tt.want)

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}

func Test_StringSet_Sorted(t *testing.T) {
	tests := []struct {
		description string
		input       StringSet
		want        []string
	}{
		{
			description: "return empty slice",
			input:       StringSet{},
			want:        []string{},
		},
		{
			description: "return slice of set of single item",
			input: StringSet{
				"example": true,
			},
			want: []string{"example"},
		},
		{
			description: "return slice of set of multiple items",
			input: StringSet{
				"example-2": true,
				"example-1": true,
			},
			want: []string{"example-1", "example-2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := tt.input.Sorted()

			sort.Strings(got)
			sort.Strings(tt.want)

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}

func Test_NewStringSet(t *testing.T) {
	tests := []struct {
		description string
		init        []string
		want        StringSet
	}{
		{
			description: "init with no values",
			want:        StringSet{},
		},
		{
			description: "init with one value",
			init:        []string{"example"},
			want: StringSet{
				"example": true,
			},
		},
		{
			description: "init with one value",
			init:        []string{"example", "example-2"},
			want: StringSet{
				"example":   true,
				"example-2": true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := NewStringSet(tt.init...)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}

func Test_StringSet_Contains(t *testing.T) {
	tests := []struct {
		description string
		input       StringSet
		key         string
		want        bool
	}{
		{
			description: "return false if not in set slice",
			input:       StringSet{},
			key:         "example",
			want:        false,
		},
		{
			description: "return true for key in set",
			input: StringSet{
				"example": true,
			},
			key:  "example",
			want: true,
		},
		{
			description: "return true for key in set",
			input: StringSet{
				"example": true,
			},
			key:  "doesnotexist",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := tt.input.Contains(tt.key)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}
