package middleware

import (
	"context"
	"net/http"

	"github.com/prajwalad101/datekeeper/model"
	"github.com/prajwalad101/datekeeper/utils"
)

func VerifyJWT(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized)
			return
		}

		bearer := "Bearer "

		// extract the token from the header
		token := authHeader[len(bearer):]

		claims, err := utils.Authorize(token, utils.Env.JWTSecret)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized)
			return
		}

		_, err = model.GetUserByID(claims.UserID)
		if err != nil {
			utils.WriteJSON(w, http.StatusNotFound, utils.ApiError{Error: "User doesnot exist"})
			return
		}

		ctx := context.WithValue(r.Context(), "userID", claims.UserID)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
