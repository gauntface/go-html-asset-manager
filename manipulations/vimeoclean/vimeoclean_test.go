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

package vimeoclean

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/gauntface/go-html-asset-manager/v2/manipulations"
	"github.com/gauntface/go-html-asset-manager/v2/utils/config"
	"github.com/gauntface/go-html-asset-manager/v2/utils/vimeoapi"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/html"
)

var errInjected = errors.New("injected error")

func Test_Manipulator(t *testing.T) {
	tests := []struct {
		description string
		runtime     manipulations.Runtime
		doc         *html.Node
		want        string
		wantError   error
	}{
		{
			description: "do nothing if no vimeo client is available",
			runtime:     manipulations.Runtime{},
			doc:         MustGetNode(t, ""),
			want:        "<html><head></head><body></body></html>",
		},
		{
			description: "do nothing if no iframes or images exist",
			runtime: manipulations.Runtime{
				HasVimeo: true,
			},
			doc:  MustGetNode(t, ""),
			want: "<html><head></head><body></body></html>",
		},
		{
			description: "do nothing if src attribute does not",
			runtime: manipulations.Runtime{
				HasVimeo: true,
			},
			doc:  MustGetNode(t, `<iframe></iframe>`),
			want: `<html><head></head><body><iframe></iframe></body></html>`,
		},
		{
			description: "do nothing for non-vimeo iframe without protocol",
			runtime: manipulations.Runtime{
				HasVimeo: true,
			},
			doc:  MustGetNode(t, `<iframe src="other.com/example"></iframe>`),
			want: `<html><head></head><body><iframe src="other.com/example"></iframe></body></html>`,
		},
		{
			description: "return error if src cannot be parsed as a URL",
			runtime: manipulations.Runtime{
				HasVimeo: true,
			},
			doc:       MustGetNode(t, `<iframe src=":"></iframe>`),
			want:      `<html><head></head><body><iframe src=":"></iframe></body></html>`,
			wantError: errURLParse,
		},
		{
			description: "do nothing for non embed vimeo URL",
			runtime: manipulations.Runtime{
				HasVimeo: true,
			},
			doc:  MustGetNode(t, `<iframe src="https://vimeo.com/1234567" iframeborder="0" other="test"></iframe>`),
			want: `<html><head></head><body><iframe src="https://vimeo.com/1234567" iframeborder="0" other="test"></iframe></body></html>`,
		},
		{
			description: "do nothing for non-video embed",
			runtime: manipulations.Runtime{
				HasVimeo: true,
			},
			doc:  MustGetNode(t, `<iframe src="https://player.vimeo.com/other/1234567" iframeborder="0" other="test"></iframe>`),
			want: `<html><head></head><body><iframe src="https://player.vimeo.com/other/1234567" iframeborder="0" other="test"></iframe></body></html>`,
		},
		{
			description: "return nothing if there are no sizes",
			runtime: manipulations.Runtime{
				HasVimeo: true,
				Vimeo: &vimeoapiStub{
					VideoError: map[string]error{
						"1234567": errInjected,
					},
				},
				Config: &config.Config{},
			},
			doc:  MustGetNode(t, `<iframe src="player.vimeo.com/video/1234567?random=searchparam" iframeborder="0" other="test"></iframe>`),
			want: `<html><head></head><body><iframe src="player.vimeo.com/video/1234567?random=searchparam" iframeborder="0" other="test"></iframe></body></html>`,
		},
		{
			description: "return error if looking up vimeo video fails",
			runtime: manipulations.Runtime{
				HasVimeo: true,
				Vimeo: &vimeoapiStub{
					VideoError: map[string]error{
						"1234567": errInjected,
					},
				},
				Config: &config.Config{
					ImgToPicture: []*config.ImgToPicConfig{
						{
							ID: "body",
							SourceSizes: []string{
								"min-width(800px) 800w",
								"100vw",
							},
						},
					},
				},
			},
			wantError: errInjected,
			doc:       MustGetNode(t, `<iframe src="player.vimeo.com/video/1234567?random=searchparam" iframeborder="0" other="test"></iframe>`),
			want:      `<html><head></head><body><iframe src="player.vimeo.com/video/1234567?random=searchparam" iframeborder="0" other="test"></iframe></body></html>`,
		},
		{
			description: "do nothing if api returns no sizes",
			runtime: manipulations.Runtime{
				HasVimeo: true,
				Vimeo: &vimeoapiStub{
					VideoResult: map[string]*vimeoapi.Video{
						"1234567": {},
					},
				},
				Config: &config.Config{
					ImgToPicture: []*config.ImgToPicConfig{
						{
							ID: "body",
							SourceSizes: []string{
								"min-width(800px) 800w",
								"100vw",
							},
						},
					},
				},
			},
			doc:  MustGetNode(t, `<iframe src="player.vimeo.com/video/1234567?random=searchparam" iframeborder="0" other="test"></iframe>`),
			want: `<html><head></head><body><iframe src="player.vimeo.com/video/1234567?random=searchparam" iframeborder="0" other="test"></iframe></body></html>`,
		},
		{
			description: "clean attributes and URL for vimeo src without protocol",
			runtime: manipulations.Runtime{
				HasVimeo: true,
				Vimeo: &vimeoapiStub{
					VideoResult: map[string]*vimeoapi.Video{
						"1234567": {
							Name:        "Example name",
							Description: "Example description",
							Link:        "example.vimeo.com/1234567",
							Width:       4,
							Height:      3,
							Pictures: vimeoapi.Pictures{
								Sizes: []vimeoapi.Picture{
									{
										Width:              2,
										LinkWithPlayButton: "https://cdn.vimeo.com/1234567/2",
									},
									{
										Width:              1,
										LinkWithPlayButton: "https://cdn.vimeo.com/1234567/1",
									},
									{
										Width:              3,
										LinkWithPlayButton: "https://cdn.vimeo.com/1234567/3",
									},
								},
							},
						},
					},
				},
				Config: &config.Config{
					ImgToPicture: []*config.ImgToPicConfig{
						{
							ID: "body",
							SourceSizes: []string{
								"min-width(800px) 800w",
								"100vw",
							},
						},
					},
				},
			},
			doc:  MustGetNode(t, `<iframe src="player.vimeo.com/video/1234567?random=searchparam" iframeborder="0" other="test"></iframe>`),
			want: `<html><head></head><body><div class="n-hopin-u-ratio-container"><div class="n-hopin-u-ratio-container__wrapper" style="position:relative;padding-bottom:75%;"><div class="n-hopin-lite-vi" videoid="1234567" style="position:absolute;width:100%;height:100%;"><a href="example.vimeo.com/1234567" class="n-hopin-lite-vi__link" target="_blank"><img style="width: 100%; height: 100%; object-fit: contain;" sizes="min-width(800px) 800w,100vw" srcset="https://cdn.vimeo.com/1234567/1 1w,https://cdn.vimeo.com/1234567/2 2w,https://cdn.vimeo.com/1234567/3 3w" src="https://cdn.vimeo.com/1234567/3" alt="Vimeo video &#34;Example name&#34; described as &#34;Example description&#34;"/></a></div></div></div></body></html>`,
		},
		{
			description: "clean attributes and URL for vimeo src with http protocol",
			runtime: manipulations.Runtime{
				HasVimeo: true,
				Vimeo: &vimeoapiStub{
					VideoResult: map[string]*vimeoapi.Video{
						"1234567": {
							Name:        "Example name",
							Description: "Example description",
							Link:        "example.vimeo.com/1234567",
							Width:       4,
							Height:      3,
							Pictures: vimeoapi.Pictures{
								Sizes: []vimeoapi.Picture{
									{
										Width:              2,
										LinkWithPlayButton: "https://cdn.vimeo.com/1234567/2",
									},
									{
										Width:              1,
										LinkWithPlayButton: "https://cdn.vimeo.com/1234567/1",
									},
									{
										Width:              3,
										LinkWithPlayButton: "https://cdn.vimeo.com/1234567/3",
									},
								},
							},
						},
					},
				},
				Config: &config.Config{
					ImgToPicture: []*config.ImgToPicConfig{
						{
							ID: "body",
							SourceSizes: []string{
								"min-width(800px) 800w",
								"100vw",
							},
						},
					},
				},
			},
			doc:  MustGetNode(t, `<iframe src="http://player.vimeo.com/video/1234567?random=searchparam" iframeborder="0" other="test"></iframe>`),
			want: `<html><head></head><body><div class="n-hopin-u-ratio-container"><div class="n-hopin-u-ratio-container__wrapper" style="position:relative;padding-bottom:75%;"><div class="n-hopin-lite-vi" videoid="1234567" style="position:absolute;width:100%;height:100%;"><a href="example.vimeo.com/1234567" class="n-hopin-lite-vi__link" target="_blank"><img style="width: 100%; height: 100%; object-fit: contain;" sizes="min-width(800px) 800w,100vw" srcset="https://cdn.vimeo.com/1234567/1 1w,https://cdn.vimeo.com/1234567/2 2w,https://cdn.vimeo.com/1234567/3 3w" src="https://cdn.vimeo.com/1234567/3" alt="Vimeo video &#34;Example name&#34; described as &#34;Example description&#34;"/></a></div></div></div></body></html>`,
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

type vimeoapiStub struct {
	VideoError  map[string]error
	VideoResult map[string]*vimeoapi.Video
}

func (v vimeoapiStub) Video(videoID string) (*vimeoapi.Video, error) {
	return v.VideoResult[videoID], v.VideoError[videoID]
}
