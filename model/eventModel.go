package model

import (
	"fmt"
	"log"

	"github.com/prajwalad101/datekeeper/datastore"
)

type Event struct {
	Id     string `json:"id"`
	UserID string `json:"userID"`
	Name   string `json:"name"`
	Note   string `json:"note,omitempty"`
	Date   string `json:"date,omitempty"`
}

func GetEventByID(id string) (*Event, error) {
	if id == "" {
		return nil, fmt.Errorf("Id is required")
	}
	row := datastore.DB.QueryRow("SELECT * FROM events WHERE id = $1", id)

	event := new(Event)
	err := row.Scan(&event.Id, &event.UserID, &event.Name, &event.Note, &event.Date)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func GetEvents(userID string) ([]*Event, error) {
	if userID == "" {
		return nil, fmt.Errorf("User id is required")
	}

	rows, err := datastore.DB.Query("SELECT * FROM events WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := make([]*Event, 0)
	for rows.Next() {
		event := new(Event)
		err := rows.Scan(&event.Id, &event.UserID, &event.Name, &event.Note, &event.Date)
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

func CreateEvent(e *Event, userID int) error {
	if userID == 0 {
		return fmt.Errorf("User id is required")
	}

	_, err := datastore.DB.Exec(
		"INSERT INTO events (id, name, note, date, user_id) VALUES (DEFAULT, $1, $2, $3, $4)",
		e.Name,
		e.Note,
		e.Date,
		userID,
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
