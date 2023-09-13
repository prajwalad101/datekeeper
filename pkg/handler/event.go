package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/prajwalad101/datekeeper/pkg/datastore"
	"github.com/prajwalad101/datekeeper/pkg/utils"
)

type Event struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Note string `json:"note,omitempty"`
	Date string `json:"date,omitempty"`
}

func ListEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	rows, err := datastore.DBConnection.Query("SELECT * FROM events")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	events := make([]*Event, 0)
	for rows.Next() {
		event := new(Event)
		err := rows.Scan(&event.Id, &event.Name, &event.Note, &event.Date)
		if err != nil {
			log.Fatal(err)
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	responseData, err := json.Marshal(events)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseData)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	row := datastore.DBConnection.QueryRow("SELECT * FROM events WHERE id = $1", id)

	event := new(Event)
	err := row.Scan(&event.Id, &event.Name, &event.Note, &event.Date)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	jsonEvent, err := json.Marshal(event)
	if err != nil {
		http.Error(w, "Error marshaling event to json", 500)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonEvent)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	var e Event

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if e.Name == "" || e.Date == "" {
		errorResponse := utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Please provide required fields (name, note, date)",
		}
		utils.SendErrorResponse(w, errorResponse)
		return
	}

	result, err := datastore.DBConnection.Exec(
		"INSERT INTO events VALUES(DEFAULT, $1, $2, $3)",
		e.Name,
		e.Note,
		e.Date,
	)
	if err != nil {
		log.Printf("Error %v", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "Event %s created successfully (%d row affected)\n", e.Name, rowsAffected)
}

/* func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing middlewareOne")
		next.ServeHTTP(w, r)
		log.Print("Executing middlewareOne again")
	})
}

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing middlewareTwo")
		if r.URL.Path == "/foo" {
			return
		}

		next.ServeHTTP(w, r)
		log.Print("Executing middlewareTwo again")
	})
}

func final(w http.ResponseWriter, _ *http.Request) {
	log.Print("Executing finalHandler")
	w.Write([]byte("OK"))
} */
