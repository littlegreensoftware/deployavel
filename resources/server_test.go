package resources_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/littlegreensoftware/deployavel/mocks"
	"github.com/littlegreensoftware/deployavel/resources"
)

var data *os.File

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

func ServerList(t *testing.T) {
	r := mocks.MakeMockAPIRequest(data, http.StatusOK)

	servers, err := resources.ServerList(r)
	if err != nil {
		t.Errorf("Should not have an error: %v", err)
	}

	if servers == nil {
		t.Error("Should have servers")
	}
}

func ServerListErrors(t *testing.T) {
	for key := range errors {
		r := mocks.MakeMockAPIRequest(data, errors[key])

		servers, err := resources.ServerList(r)
		if err == nil {
			t.Errorf("Should have an error")
		}

		if servers != nil {
			t.Error("Should not have a server")
		}
	}
}

func ServerRead(t *testing.T) {
	serverData := JsonData("server.json")
	defer serverData.Close()

	r := mocks.MakeMockAPIRequest(serverData, http.StatusOK)

	server, err := resources.ServerRead(r, 100)
	if err != nil {
		t.Error("Should not have an error")
	}

	if server == nil {
		t.Error("Should have a non zero server")
	}
}

func ServerReadErrors(t *testing.T) {
	for key := range errors {
		r := mocks.MakeMockAPIRequest(data, errors[key])

		server, err := resources.ServerRead(r, 100)
		if err == nil {
			t.Errorf("Should have an error")
		}

		if server != nil {
			t.Error("Should not have a server")
		}
	}
}

func ServerEnableOpCache(t *testing.T) {
	r := mocks.MakeMockAPIRequest(data, http.StatusOK)

	err := resources.EnableOpCache(r, 100)
	if err != nil {
		t.Error("Should not have an error")
	}
}

func ServerEnableOpCacheErrors(t *testing.T) {
	for key := range errors {
		r := mocks.MakeMockAPIRequest(data, errors[key])

		err := resources.EnableOpCache(r, 100)
		if err == nil {
			t.Errorf("Should have an error")
		}
	}
}

func ServerDisableOpCache(t *testing.T) {
	r := mocks.MakeMockAPIRequest(data, http.StatusOK)

	err := resources.DisableOpCache(r, 100)
	if err != nil {
		t.Error("Should not have an error")
	}
}

func ServerDisableOpCacheErrors(t *testing.T) {
	for key := range errors {
		r := mocks.MakeMockAPIRequest(data, errors[key])

		err := resources.DisableOpCache(r, 100)
		if err == nil {
			t.Errorf("Should have an error")
		}
	}
}

func TestServerSetup(t *testing.T) {
	data = JsonData("servers.json")
	defer data.Close()

	t.Run("should list all servers", ServerList)
	t.Run("should hanlde list all server errors", ServerListErrors)
	t.Run("should get a single server", ServerRead)
	t.Run("should handle a single server error", ServerReadErrors)
	t.Run("should enable opcache", ServerEnableOpCache)
	t.Run("should handle opcache enable errors", ServerEnableOpCacheErrors)
	t.Run("should disable opcache", ServerDisableOpCache)
	t.Run("should handle opcache disable errors", ServerDisableOpCacheErrors)
}

// JsonData returns a file pointer to data
func JsonData(dat string) *os.File {
	data, err := os.Open("../test_data/v1/" + dat)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
