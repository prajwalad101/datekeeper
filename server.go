package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/prajwalad101/datekeeper/pkg/datastore"
	"github.com/prajwalad101/datekeeper/pkg/db"
	"github.com/prajwalad101/datekeeper/pkg/handler"
	"github.com/prajwalad101/datekeeper/pkg/service"
	"github.com/prajwalad101/datekeeper/pkg/utils"
)

// Load Environment variables
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Initialize database connection
func init() {
	db := db.InitDB()
	datastore.DBConnection = db
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/events/list", handler.ListEvents)
	mux.HandleFunc("/events/show", handler.GetEvent)
	mux.HandleFunc("/events/create", handler.CreateEvent)
	mux.HandleFunc("/events/update", handler.UpdateEvent)
	mux.HandleFunc("/events/delete", handler.DeleteEvent)

	service.Schedule()

	env := utils.GetEnv()
	log.Print("Listening on ", env.Port)
	err := http.ListenAndServe(env.Port, mux)
	log.Fatal(err)
}
