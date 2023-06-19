package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), " ")

		fmt.Println("AuthHeader", authHeader)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
