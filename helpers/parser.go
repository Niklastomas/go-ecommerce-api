package helpers

import (
	"encoding/json"
	"net/http"
)

func ParseRequest(w http.ResponseWriter, r *http.Request, data interface{}) error {
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}
