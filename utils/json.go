package utils

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ApiError struct {
	Error string `json:"error"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}
