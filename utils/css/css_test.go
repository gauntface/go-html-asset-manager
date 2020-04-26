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

package css

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Format(t *testing.T) {
	tests := []struct {
		description string
		namespace   string
		typ         CSSType
		body        string
		element     string
		modifier    string
		want        string
	}{
		{
			description: "return component with body only",
			namespace:   "",
			typ:         ComponentType,
			body:        "example",
			element:     "",
			modifier:    "",
			want:        `c-example`,
		},
		{
			description: "return layout with body and namespace",
			namespace:   "example-namespace",
			typ:         LayoutType,
			body:        "example-body",
			element:     "",
			modifier:    "",
			want:        `n-example-namespace-l-example-body`,
		},
		{
			description: "return utility with body and element",
			namespace:   "",
			typ:         UtilityType,
			body:        "example-body",
			element:     "example-element",
			modifier:    "",
			want:        `u-example-body__example-element`,
		},
		{
			description: "return utility with body and modifier",
			namespace:   "",
			typ:         UtilityType,
			body:        "example-body",
			element:     "",
			modifier:    "example-modifier",
			want:        `u-example-body--example-modifier`,
		},
		{
			description: "return component with everything",
			namespace:   "example-namespace",
			typ:         ComponentType,
			body:        "example-body",
			element:     "example-element",
			modifier:    "example-modifier",
			want:        `n-example-namespace-c-example-body__example-element--example-modifier`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := Format(tt.namespace, tt.typ, tt.body, tt.element, tt.modifier)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("Unexpected HTML files; diff %v", diff)
			}
		})
	}
}
