package nasa

import (
	"net/http"
)

const (
	url = "https://api.nasa.gov/"
)

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

func NewClient(key string) *Client {
	return &Client{
		BaseURL:    url,
		apiKey:     key,
		HTTPClient: &http.Client{},
	}
}
