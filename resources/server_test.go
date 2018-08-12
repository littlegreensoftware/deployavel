package resources_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/littlegreensoftware/deployavel/mocks"
	"github.com/littlegreensoftware/deployavel/resources"
)

var errors = map[string]int{
	"bad_request":           http.StatusBadRequest,
	"unauthorized":          http.StatusUnauthorized,
	"not_found":             http.StatusNotFound,
	"unprocessable_entity":  http.StatusUnprocessableEntity,
	"too_many_requests":     http.StatusTooManyRequests,
	"internal_server_error": http.StatusInternalServerError,
	"service_unavailable":   http.StatusServiceUnavailable,
	"server_error":          1000,
}

func TestServerList(t *testing.T) {
	data := JsonData("servers.json")
	defer data.Close()

	r := mocks.MakeMockAPIRequest(data, http.StatusOK)

	servers, err := resources.ServerList(r)
	if err != nil {
		t.Errorf("Should not have an error: %v", err)
	}

	if len(servers.All) == 0 {
		t.Error("Should have servers")
	}
}

func TestServerListErrors(t *testing.T) {
	data := JsonData("servers.json")
	defer data.Close()

	for key := range errors {
		r := mocks.MakeMockAPIRequest(data, errors[key])

		servers, err := resources.ServerList(r)
		if err == nil {
			t.Errorf("Should have an error")
		}

		if len(servers.All) != 0 {
			t.Error("Should not have a server")
		}
	}
}

func TestServerGet(t *testing.T) {
	data := JsonData("server.json")
	defer data.Close()

	r := mocks.MakeMockAPIRequest(data, http.StatusOK)

	server, err := resources.ServerRead(r, 100)
	if err != nil {
		t.Error("Should not have an error")
	}

	if server.ID == 0 {
		t.Error("Should have a non zero server")
	}
}

func TestServerReadErrors(t *testing.T) {
	data := JsonData("server.json")
	defer data.Close()

	for key := range errors {
		r := mocks.MakeMockAPIRequest(data, errors[key])

		server, err := resources.ServerRead(r, 100)
		if err == nil {
			t.Errorf("Should have an error")
		}

		if server.ID != 0 {
			t.Error("Should not have a server")
		}
	}
}

func TestServerEnableOpCache(t *testing.T) {
	data := JsonData("ok.json")
	defer data.Close()

	r := mocks.MakeMockAPIRequest(data, http.StatusOK)

	err := resources.EnableOpCache(r, 100)
	if err != nil {
		t.Error("Should not have an error")
	}
}

func TestServerEnableOpCacheErrors(t *testing.T) {
	data := JsonData("ok.json")
	defer data.Close()

	for key := range errors {
		r := mocks.MakeMockAPIRequest(data, errors[key])

		err := resources.EnableOpCache(r, 100)
		if err == nil {
			t.Errorf("Should have an error")
		}
	}
}

func TestServerDisableOpCache(t *testing.T) {
	data := JsonData("ok.json")
	defer data.Close()

	r := mocks.MakeMockAPIRequest(data, http.StatusOK)

	err := resources.DisableOpCache(r, 100)
	if err != nil {
		t.Error("Should not have an error")
	}
}

func TestServerDisableOpCacheErrors(t *testing.T) {
	data := JsonData("ok.json")
	defer data.Close()

	for key := range errors {
		r := mocks.MakeMockAPIRequest(data, errors[key])

		err := resources.DisableOpCache(r, 100)
		if err == nil {
			t.Errorf("Should have an error")
		}
	}
}

// JsonData returns a file pointer to data
func JsonData(dat string) *os.File {
	data, err := os.Open("../test_data/v1/" + dat)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
