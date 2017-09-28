package mlpapi

import (
	"encoding/json"
	"net/http"
)

// GetJSON decodes JSON into a target struct
func GetJSON(r *http.Response, target interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
