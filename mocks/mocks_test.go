package mocks

import (
	"net/http"
	"testing"
)

func TestMocks(t *testing.T) {

	server := DefaultTestServer()
	defer server.Close()

	c := MockHTTPClient(server.URL)

	resp, err := c.Client.Get(server.URL)
	if err != nil {
		t.Errorf("should not return a error %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("should return a %d status code", http.StatusOK)
	}
}
