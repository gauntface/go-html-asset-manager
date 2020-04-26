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
)

type StringSet map[string]bool

func NewStringSet(strs ...string) StringSet {
	ss := StringSet{}
	for _, s := range strs {
		ss.Add(s)
	}
	return ss
}

func (s StringSet) Add(k string) {
	s[k] = true
}

func (s StringSet) Merge(st StringSet) {
	for k := range st {
		s.Add(k)
	}
}

func (s StringSet) Slice() []string {
	sl := []string{}
	for k := range s {
		sl = append(sl, k)
	}
	return sl
}

func (s StringSet) Sorted() []string {
	sl := s.Slice()
	sort.Strings(sl)
	return sl
}

func (s StringSet) Contains(k string) bool {
	_, ok := s[k]
	return ok
}
