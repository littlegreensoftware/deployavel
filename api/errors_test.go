package api_test

import (
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/littlegreensoftware/deployavel/api"
)

func TestRequestError(t *testing.T) {
	input := map[string]int{
		"bad_request":           http.StatusBadRequest,
		"unauthorized":          http.StatusUnauthorized,
		"not_found":             http.StatusNotFound,
		"unprocessable_entity":  http.StatusUnprocessableEntity,
		"too_many_requests":     http.StatusTooManyRequests,
		"internal_server_error": http.StatusInternalServerError,
		"service_unavailable":   http.StatusServiceUnavailable,
		"server_error":          1000,
	}

	for in := range input {
		err := api.RequestError(input[in])

		index := strings.Index(err.Error(), strconv.Itoa(input[in]))

		if index == -1 {
			t.Errorf("Expected to find error code: %d", input[in])
		}
	}
}
