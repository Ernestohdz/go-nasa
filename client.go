package nasa

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	url = "https://api.nasa.gov/"
)

type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

type ClientOption func(*Client)

func NewClient(options ...ClientOption) *Client {
	c := &Client{
		baseURL: url,
	}

	for _, op := range options {
		op(c)
	}
	if c.apiKey == "" {
		c.apiKey = "DEMO_KEY"
	}
	if c.httpClient == nil {
		c.httpClient = http.DefaultClient
	}
	return c
}
func WithKey(key string) ClientOption {
	return func(c *Client) {
		c.apiKey = key
	}
}
func WithClient(h *http.Client) ClientOption {
	return func(client *Client) {
		client.httpClient = h
	}
}

func (c *Client) Key() string {
	return c.apiKey
}
func (c *Client) HttpClient() *http.Client {
	return c.httpClient
}

func (c *Client) send(req *http.Request, d interface{}) error {
	q := req.URL.Query()
	q.Add("api_key", c.apiKey)
	req.URL.RawQuery = q.Encode()

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(res.Status)
	}
	return json.NewDecoder(res.Body).Decode(d)
}
