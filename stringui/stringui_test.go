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

package stringui

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_dividerString(t *testing.T) {
	tests := []struct {
		description string
		widths      []int
		want        string
	}{
		{
			description: "return divider",
			widths:      []int{1, 1},
			want:        "---------",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := dividerString(tt.widths)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}

func Test_rowString(t *testing.T) {
	tests := []struct {
		description string
		values      []string
		widths      []int
		want        string
	}{
		{
			description: "return row with correct widths",
			values:      []string{"a", "b", "c"},
			widths:      []int{1, 2, 3},
			want:        "| a | b  | c   |",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := rowString(tt.values, tt.widths)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}

func Test_rowsString(t *testing.T) {
	tests := []struct {
		description string
		rows        [][]string
		widths      []int
		want        string
	}{
		{
			description: "return row with correct widths",
			rows: [][]string{
				[]string{"a", "b", "c"},
				[]string{"d", "e", "f"},
			},
			widths: []int{1, 2, 3},
			want: `| a | b  | c   |
| d | e  | f   |`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := rowsString(tt.rows, tt.widths)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}

func Test_padRight(t *testing.T) {
	tests := []struct {
		description string
		s           string
		pad         string
		length      int
		want        string
	}{
		{
			description: "return string of current length",
			s:           "A",
			pad:         "B",
			length:      1,
			want:        "A",
		},
		{
			description: "return string with padding",
			s:           "A",
			pad:         "B",
			length:      3,
			want:        "ABB",
		},
		{
			description: "return string as is if longer than length",
			s:           "AAAA",
			pad:         "B",
			length:      3,
			want:        "AAAA",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := padRight(tt.s, tt.pad, tt.length)
			if got != tt.want {
				t.Fatalf("Unexpected result; got %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_columnWidths(t *testing.T) {
	tests := []struct {
		description string
		rows        [][]string
		want        []int
	}{
		{
			description: "return longest column width",
			rows: [][]string{
				[]string{"1234567890", "1", "1234"},
				[]string{"1", "1234", "1234567890"},
				[]string{"1234", "1234567890", "1"},
			},
			want: []int{
				10, 10, 10,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := columnWidths(tt.rows)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}

func Test_Table(t *testing.T) {
	tests := []struct {
		description string
		headings    []string
		rows        [][]string
		want        string
	}{
		{
			description: "return formatted table",
			headings: []string{
				"Column 1",
				"Column 2",
				"Column 3",
			},
			rows: [][]string{
				[]string{
					"1.1",
					"1.2",
					"1.3",
				},
				[]string{
					"2.1",
					"2.2",
					"2.3",
				},
				[]string{
					"3.1",
					"3.2",
					"3.3",
				},
			},
			want: `
----------------------------------
| Column 1 | Column 2 | Column 3 |
----------------------------------
| 1.1      | 1.2      | 1.3      |
| 2.1      | 2.2      | 2.3      |
| 3.1      | 3.2      | 3.3      |
----------------------------------
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := Table(tt.headings, tt.rows)
			if diff := cmp.Diff(got, strings.TrimSpace(tt.want)); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}
