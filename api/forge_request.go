package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// ForgeRequest handles connections to forge
// implements api.Request
type ForgeRequest struct {
	Token  string
	Client Client
}

// Get issues a GET request
func (r ForgeRequest) Get(endpoint string) (*http.Response, error) {
	req, err := r.MakeRequest(http.MethodGet, endpoint, `{...}`)
	if err != nil {
		return nil, err
	}

	return r.Client.Do(req)
}

// Post issues a POST request
func (r ForgeRequest) Post(endpoint string, data interface{}) (*http.Response, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := r.MakeRequest(http.MethodPost, endpoint, string(body))
	if err != nil {
		return nil, err
	}

	return r.Client.Do(req)
}

// Delete issues a DELETE request
func (r ForgeRequest) Delete(endpoint string) (*http.Response, error) {
	req, err := r.MakeRequest(http.MethodDelete, endpoint, `{...}`)
	if err != nil {
		return nil, err
	}

	return r.Client.Do(req)
}

// MakeRequest handles formating a request
func (r ForgeRequest) MakeRequest(method string, endpoint string, body string) (*http.Request, error) {
	url := fmt.Sprintf("%s/%s", r.Client.URL, endpoint)

	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if req == nil {
		return req, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", r.Token))

	return req, err
}
