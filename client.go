package nasa

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
	rateLimit  int
}

type ClientOption func(*Client)

/* Returns new Client */
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

/* Set API Key option */
func WithKey(key string) ClientOption {
	return func(c *Client) {
		c.apiKey = key
	}
}

/* Set HTTP Client option */
func WithClient(h *http.Client) ClientOption {
	return func(client *Client) {
		client.httpClient = h
	}
}

func WithBaseURL(url string) ClientOption {
	return func(c *Client) {
		c.baseURL = url
	}
}

func (c *Client) RateLimit() int {
	return c.rateLimit
}

func (c *Client) Key() string {
	return c.apiKey
}
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

	// add queries
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
