package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/prajwalad101/datekeeper/pkg/datastore"
	"github.com/prajwalad101/datekeeper/pkg/db"
	"github.com/prajwalad101/datekeeper/pkg/handler"
	"github.com/robfig/cron"
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
	fmt.Println("Running Schedule")
	c := cron.New()
	mux := http.NewServeMux()
	c.Start()

	mux.HandleFunc("/events/list", handler.ListEvents)
	mux.HandleFunc("/events/show", handler.GetEvent)
	mux.HandleFunc("/events/create", handler.CreateEvent)

	handler.Schedule()

	log.Print("Listening on :3000")
	http.ListenAndServe(":3000", mux)
}
