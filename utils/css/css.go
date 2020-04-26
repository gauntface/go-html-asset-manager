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

import "fmt"

type CSSType string

const (
	ComponentType CSSType = "c"
	LayoutType            = "l"
	UtilityType           = "u"

	HopinNamespace = "hopin"
)

func Format(namespace string, typ CSSType, body, element, modifier string) string {
	cssName := ""

	// Add namespace
	if namespace != "" {
		cssName += fmt.Sprintf("n-%v", namespace)
	}

	if cssName != "" {
		cssName += "-"
	}
	cssName += fmt.Sprintf("%v-%v", typ, body)

	if element != "" {
		cssName += fmt.Sprintf("__%v", element)
	}

	if modifier != "" {
		cssName += fmt.Sprintf("--%v", modifier)
	}

	return cssName
}
