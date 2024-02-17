package utils

import (
	"encoding/json"
	"net/http"
)

// parses json form the body of a request.
func ParseBody(r *http.Request, x interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(x); err != nil {
		return err
	}
	return nil
}
