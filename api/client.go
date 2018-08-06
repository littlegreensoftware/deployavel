package api

import (
	"net/http"

	"bitbucket.org/littlegreensoftware/deployavel/config"
)

var client Client

// NewHTTPClient creates a new global client
func NewHTTPClient() {
	client = Client{
		Client: http.Client{
			Timeout: config.Timeout,
		},
		URL: config.URL,
	}
}

// GlobalHTTPClient returns a singleton client
func GlobalHTTPClient() *Client {
	return &client
}

// Client handles requests
type Client struct {
	http.Client
	URL string
}
