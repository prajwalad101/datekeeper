package model

import (
	"fmt"
	"log"

	"github.com/prajwalad101/datekeeper/datastore"
)

type Event struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Note string `json:"note,omitempty"`
	Date string `json:"date,omitempty"`
}

func GetEventByID(id string) (*Event, error) {
	if id == "" {
		return nil, fmt.Errorf("Id is required")
	}
	row := datastore.DB.QueryRow("SELECT * FROM events WHERE id = $1", id)

	event := new(Event)
	err := row.Scan(&event.Id, &event.Name, &event.Note, &event.Date)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func GetEvents() ([]*Event, error) {
	rows, err := datastore.DB.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
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
		return nil, err
	}
	return events, nil
}

func CreateEvent(e *Event) error {
	_, err := datastore.DB.Exec(
		"INSERT INTO events (id, name, note, date) VALUES (DEFAULT, $1, $2, $3)",
		e.Name,
		e.Note,
		e.Date,
	)
	return err
}

func UpdateEvent(id string, e *Event) error {
	if id == "" {
		return fmt.Errorf("Id is required")
	}

	_, err := datastore.DB.Exec(
		"UPDATE events SET name=$2, note=$3, date=$4 WHERE id=$1",
		id,
		e.Name,
		e.Note,
		e.Date,
	)
	return err
}

func DeleteEvent(id string) error {
	if id == "" {
		return fmt.Errorf("Id is required")
	}

	_, err := datastore.DB.Exec(
		"DELETE from events WHERE id=$1",
		id,
	)
	return err
}
