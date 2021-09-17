package nasa

import (
	"net/http"
	"testing"
	"time"
)

func TestClientDefault(t *testing.T) {
	client := NewClient()

	if client.apiKey != "DEMO_KEY" {
		t.Errorf("default api key not set")
	}
	if client.httpClient != http.DefaultClient {
		t.Errorf("default http client not set")
	}
}

func TestSetAPIKey(t *testing.T) {
	var key string = "random-key"
	client := NewClient(WithKey(key))

	if client.apiKey != key {
		t.Errorf("client api key not set")
	}
}

func TestSetHTTPClient(t *testing.T) {
	c := &http.Client{
		Timeout: 2 * time.Second,
	}
	client := NewClient(WithClient(c))

	if client.httpClient != c {
		t.Errorf("http client not set")
	}
}
