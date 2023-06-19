package route

import (
	"net/http"

	"github.com/Prajwalad101/datekeeper/pkg/datastore"
	"github.com/Prajwalad101/datekeeper/pkg/domain/user"
)

func UserRoute() {
	mux := datastore.Mux

	listUserHandler := http.HandlerFunc(user.ListUser)
	getUserHandler := http.HandlerFunc(user.GetUser)

	mux.Handle("/users", Protected.Then(listUserHandler))
	mux.Handle("/user", Protected.Then(getUserHandler))
}
