package handler

import (
	"github.com/gvillela7/ratelimit/internal/data/response"
	"net/http"
)

func OKHandler(w http.ResponseWriter, r *http.Request) {
	response.HttpResponse(
		w, http.StatusOK,
		"Success", map[string]interface{}{
			"message": "Ok",
		})
	return
}
