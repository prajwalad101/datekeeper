package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/prajwalad101/datekeeper/handler"
	"github.com/prajwalad101/datekeeper/utils"
)

func EventRouter(router *chi.Mux) {
	r := chi.NewRouter()

	makeHandlerFunc := utils.MakeHandlerFunc

	r.Get("/", makeHandlerFunc(handler.HandleListEvents))
	r.Get("/{id}", makeHandlerFunc(handler.HandleGetEvent))
	r.Post("/", makeHandlerFunc(handler.HandleCreateEvent))
	r.Put("/{id}", makeHandlerFunc(handler.HandleUpdateEvent))
	r.Delete("/{id}", makeHandlerFunc(handler.HandleDeleteEvent))

	router.Mount("/events", r)
}
