package response

import (
	"encoding/json"
	"net/http"
)

func HttpResponse(w http.ResponseWriter, status int, message string, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	webResponse := struct {
		StatusCode int           `json:"StatusCode"`
		Message    string        `json:"message,omitempty"`
		Data       []interface{} `json:"data,omitempty"`
	}{
		StatusCode: status,
		Message:    message,
		Data:       []interface{}{data},
	}

	json.NewEncoder(w).Encode(&webResponse)
}
