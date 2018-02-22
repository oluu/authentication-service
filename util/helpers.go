package util

import (
	"encoding/json"
	"net/http"
)

// ReadJSON takes in a request and decodes the request into the specified interface
func ReadJSON(r *http.Request, data interface{}) (interface{}, error) {
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(data)
	return data, err
}
