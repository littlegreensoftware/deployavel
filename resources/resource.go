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

// ParseResource parses a resource from a body
func ParseResource(resp *http.Response, data interface{}) error {
	bytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, data)
}
