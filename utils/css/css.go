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
	LayoutType    CSSType = "l"
	UtilityType   CSSType = "u"
)

func Format(typ CSSType, body string, opts ...FormatOption) string {
	opt := &formationOption{
		namespace: "ham",
		cssType:   typ,
		body:      body,
	}

	for _, o := range opts {
		o(opt)
	}

	// Add namespace
	cssName := ""
	if opt.namespace != "" {
		cssName += fmt.Sprintf("n-%v-", opt.namespace)
	}

	cssName += fmt.Sprintf("%v-%v", typ, body)

	if opt.element != "" {
		cssName += fmt.Sprintf("__%v", opt.element)
	}

	if opt.modifier != "" {
		cssName += fmt.Sprintf("--%v", opt.modifier)
	}

	return cssName
}

func WithType(t CSSType) FormatOption {
	return func(o *formationOption) {
		o.cssType = t
	}
}

func WithNamspace(s string) FormatOption {
	return func(o *formationOption) {
		o.namespace = s
	}
}

func WithElement(s string) FormatOption {
	return func(o *formationOption) {
		o.element = s
	}
}

func WithModifier(s string) FormatOption {
	return func(o *formationOption) {
		o.modifier = s
	}
}

type formationOption struct {
	namespace string
	cssType   CSSType
	body      string
	element   string
	modifier  string
}

type FormatOption func(o *formationOption)
