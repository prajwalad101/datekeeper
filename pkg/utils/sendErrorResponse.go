package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Encodes the error response in json and writes it to w
//
// The status code sent in the header is the same as status in the error parameter
func SendErrorResponse(w http.ResponseWriter, error ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(error.Status)
	json.NewEncoder(w).Encode(error)
}
