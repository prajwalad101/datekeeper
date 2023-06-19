package datastore

import (
	"database/sql"
	"net/http"
)

var (
	DB  *sql.DB
	Mux *http.ServeMux
)
