package resources

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Resource interface
type Resource interface {
	Marshal() ([]byte, error)
}

// BaseResource used in composition
type BaseResource struct{}

// Marshal the struct to a slice of byte
func (b BaseResource) Marshal() ([]byte, error) {
	data, err := json.Marshal(b)
	return data, err
}

// ParseResource parses a resource from a body
func ParseResource(resp *http.Response, data interface{}) error {
	bytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, data)
}
