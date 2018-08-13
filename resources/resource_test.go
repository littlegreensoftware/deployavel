package resources_test

import (
	"net/http"
	"testing"

	"github.com/littlegreensoftware/deployavel/resources"
)

func TestParseResource(t *testing.T) {
	var server *resources.Server

	jason := JsonData("servers.json")
	defer jason.Close()

	resp := &http.Response{
		StatusCode: 200,
		Body:       jason,
	}

	if err := resources.ParseResource(resp, &server); err != nil {
		t.Errorf("should not have an error: %v", err)
	}

	if server == nil {
		t.Error("should have a parsed server")
	}
}
