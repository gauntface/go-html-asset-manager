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
	"fmt"
	"strings"
)

func Table(headings []string, rows [][]string) string {
	allRows := [][]string{
		headings,
	}
	allRows = append(allRows, rows...)

	colWidths := columnWidths(allRows)

	lines := []string{
		dividerString(colWidths),
		rowString(allRows[0], colWidths),
		dividerString(colWidths),
		rowsString(allRows[1:], colWidths),
		dividerString(colWidths),
	}

	return strings.Join(lines, "\n")
}

func columnWidths(rows [][]string) []int {
	colWidths := make([]int, len(rows[0]))
	for _, cols := range rows {
		for i, col := range cols {
			if len(col) > colWidths[i] {
				colWidths[i] = len(col)
			}
		}
	}
	return colWidths
}

func dividerString(widths []int) string {
	s := "-"
	for _, c := range widths {
		s += strings.Repeat("-", c+3)
	}
	return s
}

func rowString(values []string, widths []int) string {
	s := "|"
	for i, v := range values {
		s += fmt.Sprintf(" %v |", padRight(v, " ", widths[i]))
	}
	return s
}

func rowsString(rows [][]string, widths []int) string {
	lines := []string{}
	for _, r := range rows {
		lines = append(lines, rowString(r, widths))
	}
	return strings.Join(lines, "\n")
}

func padRight(s, pad string, length int) string {
	for {
		if len(s) >= length {
			return s
		}

		s = s + pad
	}
}
