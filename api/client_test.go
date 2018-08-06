package api_test

import (
	"testing"

	"bitbucket.org/littlegreensoftware/deployavel/api"
)

func TestGlobalHTTPClient(t *testing.T) {
	// Setup the global client
	api.NewHTTPClient()

	client1 := api.GlobalHTTPClient()

	// Setup another global client
	api.NewHTTPClient()

	client2 := api.GlobalHTTPClient()

	if client1 != client2 {
		t.Error("expected bloth clients to be the same")
	}
}
