package api

import (
	"fmt"
	"net/http"
)

// RequestError determines the HTTP error
func RequestError(code int) error {
	switch code {
	case http.StatusBadRequest:
		return fmt.Errorf("%d: valid data was given but the request has failed", code)
	case http.StatusUnauthorized:
		return fmt.Errorf("%d: no valid API Key was given", code)
	case http.StatusNotFound:
		return fmt.Errorf("%d: the requested resource could not be found", code)
	case http.StatusUnprocessableEntity:
		return fmt.Errorf("%d: the payload has missing required parameters or invalid data was given", code)
	case http.StatusTooManyRequests:
		return fmt.Errorf("%d: too many attempts", code)
	case http.StatusInternalServerError:
		return fmt.Errorf("%d: request failed due to an internal error", code)
	case http.StatusServiceUnavailable:
		return fmt.Errorf("%d: Server is offline for maintenance", code)
	default:
		return fmt.Errorf("%d: Server error", code)
	}
}
