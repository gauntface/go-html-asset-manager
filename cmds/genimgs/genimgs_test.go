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

import "testing"

func TestCacheControlHeader(t *testing.T) {
	tests := []struct {
		name          string
		maxAgeSeconds int64
		want          string
	}{
		{
			name:          "default max age",
			maxAgeSeconds: 31104000,
			want:          "max-age=31104000",
		},
		{
			name:          "zero max age",
			maxAgeSeconds: 0,
			want:          "max-age=0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cacheControlHeader(tt.maxAgeSeconds)
			if got != tt.want {
				t.Fatalf("cacheControlHeader(%v) = %q, want %q", tt.maxAgeSeconds, got, tt.want)
			}
		})
	}
}
