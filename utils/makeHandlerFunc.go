package utils

import (
	"net/http"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

// Handles the error returned by the parameter function and returns a wrapper handlerFunc function
func MakeHandlerFunc(next apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := next(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
