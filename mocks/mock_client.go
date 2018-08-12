package mocks

import (
	"net/http"

	"github.com/littlegreensoftware/deployavel/api"
)

// MockHTTPClient mocks an http client
func MockHTTPClient(url string) api.Client {
	return api.Client{
		Client: *http.DefaultClient,
		URL:    url,
	}
}
