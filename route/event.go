package route

import (
	"net/http"

	"github.com/Prajwalad101/datekeeper/middleware"
	"github.com/Prajwalad101/datekeeper/pkg/datastore"
	"github.com/Prajwalad101/datekeeper/pkg/domain/event"
)

func EventRoute() {
	mux := datastore.Mux

	createEventHandler := http.HandlerFunc(event.CreateEvent)

	mux.Handle("/event", Protected.Append(middleware.EnforceJSON).
		Then(createEventHandler),
	)
}
