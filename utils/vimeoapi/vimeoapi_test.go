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

package vimeoapi

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var errInjected = errors.New("injected error")

func TestNew(t *testing.T) {
	tests := []struct {
		description string
		apiKey      string
		want        *Client
	}{
		{
			description: "return new client",
			apiKey:      "abcd1234",
			want: &Client{
				apiKey:     "abcd1234",
				httpClient: http.DefaultClient,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := New(tt.apiKey)

			opts := []cmp.Option{
				cmpopts.IgnoreUnexported(Client{}),
			}
			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}

func TestVideo(t *testing.T) {
	tests := []struct {
		description   string
		videoID       string
		newRequest    func(method, url string, body io.Reader) (*http.Request, error)
		ioutilReadAll func(r io.Reader) ([]byte, error)
		httpClient    *HTTPClientStub
		want          *Video
		wantError     error
	}{
		{
			description: "return error if new request fails",
			newRequest: func(method, url string, body io.Reader) (*http.Request, error) {
				return nil, errInjected
			},
			wantError: errInjected,
		},
		{
			description: "return error if api call fails",
			videoID:     "abcd1234",
			newRequest:  http.NewRequest,
			httpClient: &HTTPClientStub{
				DoError: map[string]error{
					"https://api.vimeo.com/videos/abcd1234": errInjected,
				},
			},
			wantError: errInjected,
		},
		{
			description: "return error if reading the response fails",
			videoID:     "abcd1234",
			newRequest:  http.NewRequest,
			ioutilReadAll: func(r io.Reader) ([]byte, error) {
				return nil, errInjected
			},
			httpClient: &HTTPClientStub{
				DoReturn: map[string]*http.Response{
					"https://api.vimeo.com/videos/abcd1234": {},
				},
			},
			wantError: errInjected,
		},
		{
			description:   "return error if response cannot be parsed",
			videoID:       "abcd1234",
			newRequest:    http.NewRequest,
			ioutilReadAll: ioutil.ReadAll,
			httpClient: &HTTPClientStub{
				DoReturn: map[string]*http.Response{
					"https://api.vimeo.com/videos/abcd1234": {
						Body: ioutil.NopCloser(strings.NewReader(`This is no JSON`)),
					},
				},
			},
			wantError: errJSONParse,
		},
		{
			description:   "return video from example response",
			videoID:       "abcd1234",
			newRequest:    http.NewRequest,
			ioutilReadAll: ioutil.ReadAll,
			httpClient: &HTTPClientStub{
				DoReturn: map[string]*http.Response{
					"https://api.vimeo.com/videos/abcd1234": {
						Body: ioutil.NopCloser(strings.NewReader(exampleVideoResponse)),
					},
				},
			},
			want: &Video{
				Name:        "Production Build of gauntface.com - 2020-05-04",
				Description: "A demo of the production build of gauntface.com loading in Chrome with and without network throttling.",
				Link:        "https://vimeo.com/414998702",
				Width:       1280,
				Height:      720,
				Pictures: Pictures{
					URI:    "/videos/414998702/pictures/888821225",
					Active: true,
					Sizes: []Picture{
						{
							Width:              100,
							Height:             75,
							Link:               "https://i.vimeocdn.com/video/888821225_100x75.jpg?r=pad",
							LinkWithPlayButton: "https://i.vimeocdn.com/filter/overlay?src0=https%3A%2F%2Fi.vimeocdn.com%2Fvideo%2F888821225_100x75.jpg&src1=http%3A%2F%2Ff.vimeocdn.com%2Fp%2Fimages%2Fcrawler_play.png",
						},
						{
							Width:              200,
							Height:             150,
							Link:               "https://i.vimeocdn.com/video/888821225_200x150.jpg?r=pad",
							LinkWithPlayButton: "https://i.vimeocdn.com/filter/overlay?src0=https%3A%2F%2Fi.vimeocdn.com%2Fvideo%2F888821225_200x150.jpg&src1=http%3A%2F%2Ff.vimeocdn.com%2Fp%2Fimages%2Fcrawler_play.png",
						},
						{
							Width:              295,
							Height:             166,
							Link:               "https://i.vimeocdn.com/video/888821225_295x166.jpg?r=pad",
							LinkWithPlayButton: "https://i.vimeocdn.com/filter/overlay?src0=https%3A%2F%2Fi.vimeocdn.com%2Fvideo%2F888821225_295x166.jpg&src1=http%3A%2F%2Ff.vimeocdn.com%2Fp%2Fimages%2Fcrawler_play.png",
						},
						{
							Width:              640,
							Height:             360,
							Link:               "https://i.vimeocdn.com/video/888821225_640x360.jpg?r=pad",
							LinkWithPlayButton: "https://i.vimeocdn.com/filter/overlay?src0=https%3A%2F%2Fi.vimeocdn.com%2Fvideo%2F888821225_640x360.jpg&src1=http%3A%2F%2Ff.vimeocdn.com%2Fp%2Fimages%2Fcrawler_play.png",
						},
						{
							Width:              1280,
							Height:             720,
							Link:               "https://i.vimeocdn.com/video/888821225_1280x720.jpg?r=pad",
							LinkWithPlayButton: "https://i.vimeocdn.com/filter/overlay?src0=https%3A%2F%2Fi.vimeocdn.com%2Fvideo%2F888821225_1280x720.jpg&src1=http%3A%2F%2Ff.vimeocdn.com%2Fp%2Fimages%2Fcrawler_play.png",
						},
						{
							Width:              1920,
							Height:             1080,
							Link:               "https://i.vimeocdn.com/video/888821225_1920x1080.jpg?r=pad",
							LinkWithPlayButton: "https://i.vimeocdn.com/filter/overlay?src0=https%3A%2F%2Fi.vimeocdn.com%2Fvideo%2F888821225_1920x1080.jpg&src1=http%3A%2F%2Ff.vimeocdn.com%2Fp%2Fimages%2Fcrawler_play.png",
						},
						{
							Width:              960,
							Height:             540,
							Link:               "https://i.vimeocdn.com/video/888821225_960x540.jpg?r=pad",
							LinkWithPlayButton: "https://i.vimeocdn.com/filter/overlay?src0=https%3A%2F%2Fi.vimeocdn.com%2Fvideo%2F888821225_960x540.jpg&src1=http%3A%2F%2Ff.vimeocdn.com%2Fp%2Fimages%2Fcrawler_play.png",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			c := &Client{
				httpNewRequest: tt.newRequest,
				ioutilReadAll:  tt.ioutilReadAll,
				httpClient:     tt.httpClient,
			}

			got, err := c.Video(tt.videoID)
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			opts := []cmp.Option{
				cmp.AllowUnexported(Client{}),
			}
			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Fatalf("Unexpected result; diff %v", diff)
			}
		})
	}
}

type HTTPClientStub struct {
	DoReturn map[string]*http.Response
	DoError  map[string]error
}

func (h *HTTPClientStub) Do(req *http.Request) (*http.Response, error) {
	return h.DoReturn[req.URL.String()], h.DoError[req.URL.String()]
}

const exampleVideoResponse = `{
	"uri": "/videos/414998702",
	"name": "Production Build of gauntface.com - 2020-05-04",
	"description": "A demo of the production build of gauntface.com loading in Chrome with and without network throttling.",
	"type": "video",
	"link": "https://vimeo.com/414998702",
	"duration": 35,
	"width": 1280,
	"language": "en",
	"height": 720,
	"embed": {
	  "buttons": {
		"like": true,
		"watchlater": true,
		"share": true,
		"embed": true,
		"hd": false,
		"fullscreen": true,
		"scaling": true
	  },
	  "logos": {
		"vimeo": true,
		"custom": {
		  "active": false,
		  "link": null,
		  "sticky": false
		}
	  },
	  "title": {
		"name": "user",
		"owner": "user",
		"portrait": "user"
	  },
	  "playbar": true,
	  "volume": true,
	  "speed": false,
	  "color": "00adef",
	  "uri": null,
	  "html": "<iframe src=\"https://player.vimeo.com/video/414998702?badge=0&amp;autopause=0&amp;player_id=0&amp;app_id=174185\" width=\"1280\" height=\"720\" frameborder=\"0\" allow=\"autoplay; fullscreen\" allowfullscreen title=\"Production Build of gauntface.com - 2020-05-04\"></iframe>",
	  "badges": {
		"hdr": false,
		"live": {
		  "streaming": false,
		  "archived": false
		},
		"staff_pick": {
		  "normal": false,
		  "best_of_the_month": false,
		  "best_of_the_year": false,
		  "premiere": false
		},
		"vod": false,
		"weekend_challenge": false
	  }
	},
	"created_time": "2020-05-05T02:44:56+00:00",
	"modified_time": "2020-05-06T03:06:48+00:00",
	"release_time": "2020-05-05T02:44:56+00:00",
	"content_rating": [ "safe" ],
	"license": null,
	"privacy": {
	  "view": "anybody",
	  "embed": "public",
	  "download": false,
	  "add": true,
	  "comments": "anybody"
	},
	"pictures": {
	  "uri": "/videos/414998702/pictures/888821225",
	  "active": true,
	  "type": "custom",
	  "sizes": [
		{
		  "width": 100,
		  "height": 75,
		  "link": "https://i.vimeocdn.com/video/888821225_100x75.jpg?r=pad",
		  "link_with_play_button": "https://i.vimeocdn.com/filter/overlay?src0=https%3A%2F%2Fi.vimeocdn.com%2Fvideo%2F888821225_100x75.jpg&src1=http%3A%2F%2Ff.vimeocdn.com%2Fp%2Fimages%2Fcrawler_play.png"
		},
		{
		  "width": 200,
		  "height": 150,
		  "link": "https://i.vimeocdn.com/video/888821225_200x150.jpg?r=pad",
		  "link_with_play_button": "https://i.vimeocdn.com/filter/overlay?src0=https%3A%2F%2Fi.vimeocdn.com%2Fvideo%2F888821225_200x150.jpg&src1=http%3A%2F%2Ff.vimeocdn.com%2Fp%2Fimages%2Fcrawler_play.png"
		},
		{
		  "width": 295,
		  "height": 166,
		  "link": "https://i.vimeocdn.com/video/888821225_295x166.jpg?r=pad",
		  "link_with_play_button": "https://i.vimeocdn.com/filter/overlay?src0=https%3A%2F%2Fi.vimeocdn.com%2Fvideo%2F888821225_295x166.jpg&src1=http%3A%2F%2Ff.vimeocdn.com%2Fp%2Fimages%2Fcrawler_play.png"
		},
		{
		  "width": 640,
		  "height": 360,
		  "link": "https://i.vimeocdn.com/video/888821225_640x360.jpg?r=pad",
		  "link_with_play_button": "https://i.vimeocdn.com/filter/overlay?src0=https%3A%2F%2Fi.vimeocdn.com%2Fvideo%2F888821225_640x360.jpg&src1=http%3A%2F%2Ff.vimeocdn.com%2Fp%2Fimages%2Fcrawler_play.png"
		},
		{
		  "width": 1280,
		  "height": 720,
		  "link": "https://i.vimeocdn.com/video/888821225_1280x720.jpg?r=pad",
		  "link_with_play_button": "https://i.vimeocdn.com/filter/overlay?src0=https%3A%2F%2Fi.vimeocdn.com%2Fvideo%2F888821225_1280x720.jpg&src1=http%3A%2F%2Ff.vimeocdn.com%2Fp%2Fimages%2Fcrawler_play.png"
		},
		{
		  "width": 1920,
		  "height": 1080,
		  "link": "https://i.vimeocdn.com/video/888821225_1920x1080.jpg?r=pad",
		  "link_with_play_button": "https://i.vimeocdn.com/filter/overlay?src0=https%3A%2F%2Fi.vimeocdn.com%2Fvideo%2F888821225_1920x1080.jpg&src1=http%3A%2F%2Ff.vimeocdn.com%2Fp%2Fimages%2Fcrawler_play.png"
		},
		{
		  "width": 960,
		  "height": 540,
		  "link": "https://i.vimeocdn.com/video/888821225_960x540.jpg?r=pad",
		  "link_with_play_button": "https://i.vimeocdn.com/filter/overlay?src0=https%3A%2F%2Fi.vimeocdn.com%2Fvideo%2F888821225_960x540.jpg&src1=http%3A%2F%2Ff.vimeocdn.com%2Fp%2Fimages%2Fcrawler_play.png"
		}
	  ],
	  "resource_key": "b4a812c750d41c7a0f88d83658921a5f46c89f90"
	},
	"tags": [
	  {
		"uri": "/tags/html",
		"name": "html",
		"tag": "html",
		"canonical": "html",
		"metadata": {
		  "connections": {
			"videos": {
			  "uri": "/tags/html/videos",
			  "options": [ "GET" ],
			  "total": 7298
			}
		  }
		},
		"resource_key": "cb3ec423ffacd013e9fb5e04eac6a33335e8dc25"
	  },
	  {
		"uri": "/tags/webdev",
		"name": "webdev",
		"tag": "webdev",
		"canonical": "webdev",
		"metadata": {
		  "connections": {
			"videos": {
			  "uri": "/tags/webdev/videos",
			  "options": [ "GET" ],
			  "total": 670
			}
		  }
		},
		"resource_key": "843d02db6cf144072cfa32c291a79418bfa4fcf3"
	  }
	],
	"stats": { "plays": 0 },
	"categories": [  ],
	"metadata": {
	  "connections": {
		"comments": {
		  "uri": "/videos/414998702/comments",
		  "options": [ "GET", "POST" ],
		  "total": 0
		},
		"credits": {
		  "uri": "/videos/414998702/credits",
		  "options": [ "GET", "POST" ],
		  "total": 0
		},
		"likes": {
		  "uri": "/videos/414998702/likes",
		  "options": [ "GET" ],
		  "total": 0
		},
		"pictures": {
		  "uri": "/videos/414998702/pictures",
		  "options": [ "GET", "POST" ],
		  "total": 2
		},
		"texttracks": {
		  "uri": "/videos/414998702/texttracks",
		  "options": [ "GET", "POST" ],
		  "total": 0
		},
		"related": null,
		"recommendations": {
		  "uri": "/videos/414998702/recommendations",
		  "options": [ "GET" ]
		},
		"albums": {
		  "uri": "/videos/414998702/albums",
		  "options": [ "GET", "PATCH" ],
		  "total": 0
		},
		"available_albums": {
		  "uri": "/videos/414998702/available_albums",
		  "options": [ "GET" ],
		  "total": 0
		},
		"available_channels": {
		  "uri": "/videos/414998702/available_channels",
		  "options": [ "GET" ],
		  "total": 0
		}
	  },
	  "interactions": {
		"watchlater": {
		  "uri": "/users/5352297/watchlater/414998702",
		  "options": [
			"GET",
			"PUT",
			"DELETE"
		  ],
		  "added": false,
		  "added_time": null
		},
		"report": {
		  "uri": "/videos/414998702/report",
		  "options": [ "POST" ],
		  "reason": [
			"pornographic",
			"harassment",
			"advertisement",
			"ripoff",
			"incorrect rating",
			"spam",
			"causes harm"
		  ]
		}
	  }
	},
	"user": {
	  "uri": "/users/5352297",
	  "name": "Matt Gaunt",
	  "link": "https://vimeo.com/gauntface",
	  "location": "Bristol, UK",
	  "gender": "m",
	  "bio": "I'm a senior developer at Mubaloo, developing and managing Android, Blackberry platforms, with some other bits and bobs thrown in for good measure :)",
	  "short_bio": null,
	  "created_time": "2010-11-29T21:49:33+00:00",
	  "pictures": {
		"uri": "/users/5352297/pictures/1410581",
		"active": true,
		"type": "custom",
		"sizes": [
		  {
			"width": 30,
			"height": 30,
			"link": "https://i.vimeocdn.com/portrait/1410581_30x30"
		  },
		  {
			"width": 75,
			"height": 75,
			"link": "https://i.vimeocdn.com/portrait/1410581_75x75"
		  },
		  {
			"width": 100,
			"height": 100,
			"link": "https://i.vimeocdn.com/portrait/1410581_100x100"
		  },
		  {
			"width": 300,
			"height": 300,
			"link": "https://i.vimeocdn.com/portrait/1410581_300x300"
		  },
		  {
			"width": 72,
			"height": 72,
			"link": "https://i.vimeocdn.com/portrait/1410581_72x72"
		  },
		  {
			"width": 144,
			"height": 144,
			"link": "https://i.vimeocdn.com/portrait/1410581_144x144"
		  },
		  {
			"width": 216,
			"height": 216,
			"link": "https://i.vimeocdn.com/portrait/1410581_216x216"
		  },
		  {
			"width": 288,
			"height": 288,
			"link": "https://i.vimeocdn.com/portrait/1410581_288x288"
		  },
		  {
			"width": 360,
			"height": 360,
			"link": "https://i.vimeocdn.com/portrait/1410581_360x360"
		  }
		],
		"resource_key": "45ec3eddc5d7e6177c60126bb876bdac97a4a76a"
	  },
	  "websites": [
		{
		  "name": null,
		  "link": "www.gauntface.co.uk",
		  "type": "link",
		  "description": null
		}
	  ],
	  "metadata": {
		"connections": {
		  "albums": {
			"uri": "/users/5352297/albums",
			"options": [ "GET" ],
			"total": 0
		  },
		  "appearances": {
			"uri": "/users/5352297/appearances",
			"options": [ "GET" ],
			"total": 0
		  },
		  "categories": {
			"uri": "/users/5352297/categories",
			"options": [ "GET" ],
			"total": 0
		  },
		  "channels": {
			"uri": "/users/5352297/channels",
			"options": [ "GET" ],
			"total": 0
		  },
		  "feed": {
			"uri": "/users/5352297/feed",
			"options": [ "GET" ]
		  },
		  "followers": {
			"uri": "/users/5352297/followers",
			"options": [ "GET" ],
			"total": 0
		  },
		  "following": {
			"uri": "/users/5352297/following",
			"options": [ "GET" ],
			"total": 0
		  },
		  "groups": {
			"uri": "/users/5352297/groups",
			"options": [ "GET" ],
			"total": 0
		  },
		  "likes": {
			"uri": "/users/5352297/likes",
			"options": [ "GET" ],
			"total": 0
		  },
		  "membership": {
			"uri": "/users/5352297/membership/",
			"options": [ "PATCH" ]
		  },
		  "moderated_channels": {
			"uri": "/users/5352297/channels?filter=moderated",
			"options": [ "GET" ],
			"total": 0
		  },
		  "portfolios": {
			"uri": "/users/5352297/portfolios",
			"options": [ "GET" ],
			"total": 0
		  },
		  "videos": {
			"uri": "/users/5352297/videos",
			"options": [ "GET" ],
			"total": 7
		  },
		  "watchlater": {
			"uri": "/users/5352297/watchlater",
			"options": [ "GET" ],
			"total": 0
		  },
		  "shared": {
			"uri": "/users/5352297/shared/videos",
			"options": [ "GET" ],
			"total": 0
		  },
		  "pictures": {
			"uri": "/users/5352297/pictures",
			"options": [ "GET", "POST" ],
			"total": 1
		  },
		  "watched_videos": {
			"uri": "/me/watched/videos",
			"options": [ "GET" ],
			"total": 0
		  },
		  "folders": {
			"uri": "/me/folders",
			"options": [ "GET", "POST" ],
			"total": 0
		  },
		  "block": {
			"uri": "/me/block",
			"options": [ "GET" ],
			"total": 0
		  }
		}
	  },
	  "location_details": {
		"formatted_address": "Bristol, UK",
		"latitude": null,
		"longitude": null,
		"city": null,
		"state": null,
		"neighborhood": null,
		"sub_locality": null,
		"state_iso_code": null,
		"country": null,
		"country_iso_code": null
	  },
	  "skills": [  ],
	  "available_for_hire": false,
	  "preferences": {
		"videos": {
		  "privacy": {
			"view": "anybody",
			"comments": "anybody",
			"embed": "public",
			"download": true,
			"add": true
		  }
		}
	  },
	  "content_filter": [
		"language",
		"drugs",
		"violence",
		"nudity",
		"safe",
		"unrated"
	  ],
	  "resource_key": "396ca720cf3d730b69bfdf17f2240c7e0a2fd5ff",
	  "account": "basic"
	},
	"review_page": {
	  "active": true,
	  "link": "https://vimeo.com/gauntface/review/414998702/c22a421785"
	},
	"parent_folder": null,
	"last_user_action_event_date": "2020-05-06T03:06:48+00:00",
	"app": {
	  "name": "Parallel Uploader",
	  "uri": "/apps/87099"
	},
	"status": "available",
	"resource_key": "5e6371d8a5cf0dd7a2daae3ecc1c6835a3cbc51d",
	"upload": {
	  "status": "complete",
	  "link": null,
	  "upload_link": null,
	  "complete_uri": null,
	  "form": null,
	  "approach": null,
	  "size": null,
	  "redirect_url": null
	},
	"transcode": { "status": "complete" }
  }`
