package nasa

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const (
	layoutISO = "2006-01-02"
)

type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
	rateLimit  int
}

// ClientOption is the type constructor options for NewClient(...)
type ClientOption func(*Client)

// NewClient constructs a new Client that can make requests to
// NASA's Open APIs
func NewClient(options ...ClientOption) *Client {
	c := &Client{}

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

// WithKey configures Client with an API key
func WithKey(key string) ClientOption {
	return func(c *Client) {
		c.apiKey = key
	}
}

// WithClient configures Client with an http.Client
func WithClient(h *http.Client) ClientOption {
	return func(client *Client) {
		client.httpClient = h
	}
}

// WithBaseURL configures Client to use a custom base url
func WithBaseURL(url string) ClientOption {
	return func(c *Client) {
		c.baseURL = url
	}
}

// RateLimit returns the current remaining api calls
func (c *Client) RateLimit() int {
	return c.rateLimit
}

// Returns Client's api key
func (c *Client) Key() string {
	return c.apiKey
}

// Returns Client's http.Client
func (c *Client) HttpClient() *http.Client {
	return c.httpClient
}

type apiConfig struct {
	host string
	path string
}

type apiRequest interface {
	params() url.Values
}

func (c *Client) get(config *apiConfig, apiReq apiRequest) (*http.Response, error) {

	host := config.host
	if c.baseURL != "" {
		host = c.baseURL
	}

	httpReq, err := http.NewRequest("GET", host+config.path, nil)

	if err != nil {
		return nil, err
	}

	httpReq.URL.RawQuery = c.generateQuery(apiReq.params())

	return c.httpClient.Do(httpReq)
}

func (c *Client) getJSON(config *apiConfig, apiReq apiRequest, d interface{}) error {
	res, err := c.get(config, apiReq)

	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}
	fmt.Sscan(res.Header.Get("X-RateLimit-Remaining"), &c.rateLimit)
	return json.NewDecoder(res.Body).Decode(d)
}

func (c *Client) generateQuery(q url.Values) string {
	q.Set("api_key", c.apiKey)

	return q.Encode()
}
