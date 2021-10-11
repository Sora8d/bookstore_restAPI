package http_utils

import (
	"encoding/json"
	"net/http"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("ContentType", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}
