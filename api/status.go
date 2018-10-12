package api

import (
	"net/http"

	"encoding/json"
)

func (v *v1) statusHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(v.status.GetStatusInformation())
	})
}
