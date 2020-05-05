package vimeoapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	host = "https://api.vimeo.com"
)

type Client struct {
	apiKey     string
	httpClient *http.Client
}

func New(apiKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: http.DefaultClient,
	}
}

func (c *Client) Video(videoID string) (*Video, error) {
	api := fmt.Sprintf("%v/videos/%v", host, videoID)

	req, err := http.NewRequest(http.MethodGet, api, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.apiKey))
	req.Header.Set("User-Agent", "go-html-asset-manager")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	video := Video{}
	jsonErr := json.Unmarshal(body, &video)
	if jsonErr != nil {
		log.Fatal(jsonErr)
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
