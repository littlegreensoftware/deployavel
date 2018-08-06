package api

import (
	"net/http"
)

// Token models an AuthToken
type Token struct {
	Auth struct {
		Value string `yaml:"AuthToken"`
	} `yaml:"config"`
}

// Request interface for interacting with an api
type Request interface {
	Get(endpoint string) (*http.Response, error)
	Post(endpoint string, data interface{}) (*http.Response, error)
	Delete(endpoint string) (*http.Response, error)
	MakeRequest(method string, url string, body string) (*http.Request, error)
}
