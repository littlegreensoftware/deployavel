package api_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/littlegreensoftware/deployavel/api"
	"github.com/littlegreensoftware/deployavel/mocks"
)

func TestForgeRequestGet(t *testing.T) {
	testServer := mocks.DefaultTestServer()
	defer testServer.Close()

	fr := api.ForgeRequest{
		Token:  "eyJ0eXAiOiJKV1QiLCJhbGc",
		Client: mocks.MockHTTPClient(testServer.URL),
	}

	res, err := fr.Get("servers/1")
	if err != nil {
		t.Errorf("Should not include an error: %v", err)
	}

	if res.StatusCode == 0 {
		t.Error("Expected a response code.")
	}

	if res.Body == nil {
		t.Error("Expected a body")
	}
}

func TestForgeRequestPost(t *testing.T) {
	testServer := mocks.DefaultTestServer()
	defer testServer.Close()

	fr := api.ForgeRequest{
		Token:  "ieyJ0eXAiOiJKV1QiLCJhbGc",
		Client: mocks.MockHTTPClient(testServer.URL),
	}

	res, err := fr.Post("servers", []byte("body"))
	if err != nil {
		t.Errorf("Should not return an error: %v", err)
	}

	if res.StatusCode == 0 {
		t.Error("Expected a response code.")
	}

	if res.Body == nil {
		t.Error("Expected a body")
	}
}

func TestForgeRequestDelete(t *testing.T) {
	testServer := mocks.DefaultTestServer()
	defer testServer.Close()

	fr := api.ForgeRequest{
		Token:  "ieyJ0eXAiOiJKV1QiLCJhbGc",
		Client: mocks.MockHTTPClient(testServer.URL),
	}

	res, err := fr.Delete("servers")
	if err != nil {
		t.Errorf("Should return an error: %v", err)
	}

	if res.StatusCode == 0 {
		t.Error("Expected a response code.")
	}

	if res.Body == nil {
		t.Error("Expected a body")
	}
}

func TestForgeRequestMakeRequest(t *testing.T) {
	url := "test.com"
	endpoint := "servers"
	token := "ieyJ0eXAiOiJKV1QiLCJhbGc"
	headers := map[string]string{
		"Accept":        "application/json",
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}

	fr := api.ForgeRequest{
		Token:  token,
		Client: mocks.MockHTTPClient(url),
	}

	request, err := fr.MakeRequest(http.MethodGet, endpoint, "data")
	if err != nil {
		t.Errorf("Should not return an error: %v", err)
	}

	if request.URL.Path != fmt.Sprintf("%s/%s", url, endpoint) {
		t.Errorf("Expected URL to be %s/%s\n", url, endpoint)
	}

	for header := range headers {
		if request.Header.Get(header) != headers[header] {
			t.Errorf("Expected %s Header to be %s\n", header, headers[header])
		}
	}
}

func TestForgeRequestMakeRequestWithBadInput(t *testing.T) {
	input := map[string]map[string]string{
		"invalid_url": map[string]string{
			"method":   http.MethodGet,
			"endpoint": "",
			"token":    "ieyJ0eXAiOiJKV1QiLCJhbGc",
			"url":      ":j:j",
		},
		"invalid_method": map[string]string{
			"method":   " ",
			"endpoint": "servers",
			"token":    "ieyJ0eXAiOiJKV1QiLCJhbGc",
			"url":      "test.com",
		},
	}

	for in := range input {
		fr := api.ForgeRequest{
			Token:  input[in]["token"],
			Client: mocks.MockHTTPClient(input[in]["url"]),
		}

		_, err := fr.MakeRequest(input[in]["method"], input[in]["endpoint"], "data")

		if err == nil {
			t.Errorf("Expected %s", in)
		}
	}
}
