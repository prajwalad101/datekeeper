package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/prajwalad101/datekeeper/handler"
	"github.com/prajwalad101/datekeeper/middleware"
	"github.com/prajwalad101/datekeeper/utils"
)

func UserRouter(router *chi.Mux) {
	r := chi.NewRouter()

	makeHandlerFunc := utils.MakeHandlerFunc

	// Public Routes
	r.Group(func(r chi.Router) {
		r.Post("/register", makeHandlerFunc(handler.HandleRegister))
		r.Post("/login", makeHandlerFunc(handler.HandleLogin))
	})

	// Private Routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.VerifyJWT)
		r.Get("/detail", makeHandlerFunc(handler.HandleGetUser))
	})

	router.Mount("/users", r)
}
