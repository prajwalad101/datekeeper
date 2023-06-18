package middleware

import (
	"mime"
	"net/http"
)

// Middleware that enforces Content-Type of application/json on each request
func EnforceJSONHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		if contentType != "" {
			mt, _, err := mime.ParseMediaType(contentType)
			if err != nil {
				msg := "Malformed Content-Type header"
				http.Error(w, msg, http.StatusBadRequest)
				return
			}

			if mt != "application/json" {
				msg := "Content-Type header must be application/json"
				http.Error(w, msg, http.StatusUnsupportedMediaType)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
