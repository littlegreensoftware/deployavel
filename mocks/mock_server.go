package mocks

import (
	"net/http"
	"net/http/httptest"
)

// DefaultTestServer creates a new http server
func DefaultTestServer() *httptest.Server {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("body"))
	}))

	return testServer
}
