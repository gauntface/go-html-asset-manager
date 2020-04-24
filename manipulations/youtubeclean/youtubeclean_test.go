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

package youtubeclean

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/gauntface/go-html-asset-injector/manipulations"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/html"
)

func Test_Manipulator(t *testing.T) {
	tests := []struct {
		description string
		runtime     manipulations.Runtime
		doc         *html.Node
		want        string
		wantError   error
	}{
		{
			description: "do nothing if no iframes or images exist",
			doc:         MustGetNode(t, ""),
			want:        "<html><head></head><body></body></html>",
		},
		{
			description: "do nothing if src attribute does not",
			doc:         MustGetNode(t, `<iframe></iframe>`),
			want:        `<html><head></head><body><iframe></iframe></body></html>`,
		},
		{
			description: "do nothing for non-youtube iframe without protocol",
			doc:         MustGetNode(t, `<iframe src="other.com/example"></iframe>`),
			want:        `<html><head></head><body><iframe src="other.com/example"></iframe></body></html>`,
		},
		{
			description: "do nothing for non-youtube iframe with protocol",
			doc:         MustGetNode(t, `<iframe src="//other.com/example"></iframe>`),
			want:        `<html><head></head><body><iframe src="//other.com/example"></iframe></body></html>`,
		},
		{
			description: "return error if src cannot be parsed as a URL",
			doc:         MustGetNode(t, `<iframe src=":"></iframe>`),
			want:        `<html><head></head><body><iframe src=":"></iframe></body></html>`,
			wantError:   errURLParse,
		},
		{
			description: "do nothing for non embed youtube URL",
			doc:         MustGetNode(t, `<iframe src="www.youtube.com/view/1234-abcd?random=searchparam" iframeborder="0" other="test"></iframe>`),
			want:        `<html><head></head><body><iframe src="www.youtube.com/view/1234-abcd?random=searchparam" iframeborder="0" other="test"></iframe></body></html>`,
		},
		{
			description: "clean attributes and URL for youtube src without protocol",
			doc:         MustGetNode(t, `<iframe src="www.youtube.com/embed/1234-abcd?random=searchparam" iframeborder="0" other="test"></iframe>`),
			want:        `<html><head></head><body><div class="n-hopin-u-ratio-container"><div class="n-hopin-u-ratio-container__wrapper" style="position:relative;padding-bottom:75%;"><div class="n-hopin-lite-yt" videoid="1234-abcd" style="position:absolute;width:100%;height:100%;"><a href="https://www.youtube.com/watch?v=1234-abcd" class="n-hopin-lite-yt__link" target="_blank"><img src="https://i.ytimg.com/vi/1234-abcd/hqdefault.jpg" style="width: 100%; height: 100%; object-fit: contain;"/></a></div></div></div></body></html>`,
		},
		{
			description: "clean attributes and URL for youtube src with http protocol",
			doc:         MustGetNode(t, `<iframe src="http://www.youtube.com/embed/1234-abcd?random=searchparam" iframeborder="0" other="test"></iframe>`),
			want:        `<html><head></head><body><div class="n-hopin-u-ratio-container"><div class="n-hopin-u-ratio-container__wrapper" style="position:relative;padding-bottom:75%;"><div class="n-hopin-lite-yt" videoid="1234-abcd" style="position:absolute;width:100%;height:100%;"><a href="https://www.youtube.com/watch?v=1234-abcd" class="n-hopin-lite-yt__link" target="_blank"><img src="https://i.ytimg.com/vi/1234-abcd/hqdefault.jpg" style="width: 100%; height: 100%; object-fit: contain;"/></a></div></div></div></body></html>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			err := Manipulator(tt.runtime, tt.doc)
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Different error returned; got %v, want %v", err, tt.wantError)
			}

			if diff := cmp.Diff(MustRenderNode(t, tt.doc), tt.want); diff != "" {
				t.Fatalf("Unexpected HTML files; diff %v", diff)
			}
		})
	}
}

func MustGetNode(t *testing.T, input string) *html.Node {
	t.Helper()

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Failed to parse HTML: %v", err)
	}
	return doc
}

func MustRenderNode(t *testing.T, n *html.Node) string {
	t.Helper()

	if n == nil {
		return ""
	}

	var buf bytes.Buffer
	err := html.Render(&buf, n)
	if err != nil {
		t.Fatalf("failed to render html node to string: %v", err)
	}

	return buf.String()
}
