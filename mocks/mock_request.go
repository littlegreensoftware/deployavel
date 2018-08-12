package mocks

import (
	"net/http"
	"os"
)

// MockAPIRequest implements the api.Request interface used for testing
type MockAPIRequest struct {
	StatusCode int
	data       *os.File
}

// Get returns a mocked http response
func (r MockAPIRequest) Get(endpoint string) (*http.Response, error) {
	return r.response()
}

// Post returns a mocked http response
func (r MockAPIRequest) Post(endpoint string, data interface{}) (*http.Response, error) {
	return r.response()
}

// Delete returns a mocked http response
func (r MockAPIRequest) Delete(endpoint string) (*http.Response, error) {
	return r.response()
}

// MakeRequest returns a mocked http request
func (r MockAPIRequest) MakeRequest(method string, url string, body string) (*http.Request, error) {
	return &http.Request{}, nil
}

func (r MockAPIRequest) response() (*http.Response, error) {
	return &http.Response{
		StatusCode: r.StatusCode,
		Body:       r.data,
	}, nil
}

// MakeMockAPIRequest creates and returns a MockAPIRequest
// data is a file pointer to the data to return for the request
// code is the status code that should be returned for the request
func MakeMockAPIRequest(data *os.File, code int) MockAPIRequest {
	return MockAPIRequest{
		StatusCode: code,
		data:       data,
	}
}
