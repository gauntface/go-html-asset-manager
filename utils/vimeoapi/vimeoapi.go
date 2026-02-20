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
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	host = "https://api.vimeo.com"
)

var (
	errJSONParse = errors.New("unable to parse JSON")
)

type Client struct {
	apiKey     string
	httpClient httpClient

	httpNewRequest func(method, url string, body io.Reader) (*http.Request, error)
	ioutilReadAll  func(r io.Reader) ([]byte, error)
}

func New(apiKey string) *Client {
	return &Client{
		apiKey:         apiKey,
		httpClient:     http.DefaultClient,
		httpNewRequest: http.NewRequest,
		ioutilReadAll:  ioutil.ReadAll,
	}
}

func (c *Client) Video(videoID string) (*Video, error) {
	api := fmt.Sprintf("%v/videos/%v", host, videoID)

	req, err := c.httpNewRequest(http.MethodGet, api, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.apiKey))
	req.Header.Set("User-Agent", "go-html-asset-manager")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := c.ioutilReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	video := Video{}
	err = json.Unmarshal(body, &video)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errJSONParse, err)
	}

	return &video, nil
}

type Pictures struct {
	URI    string    `json:"uri,omitempty"`
	Active bool      `json:"active"`
	Sizes  []Picture `json:"sizes,omitempty"`
}

type Picture struct {
	Width              int    `json:"width,omitempty"`
	Height             int    `json:"height,omitempty"`
	Link               string `json:"link,omitempty"`
	LinkWithPlayButton string `json:"link_with_play_button,omitempty"`
}

type Video struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Link        string   `json:"link,omitempty"`
	Width       int      `json:"width,omitempty"`
	Height      int      `json:"height,omitempty"`
	Pictures    Pictures `json:"pictures,omitempty"`
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
