package service

import (
	"fmt"
	"log"

	"github.com/prajwalad101/datekeeper/pkg/datastore"
	"github.com/prajwalad101/datekeeper/pkg/handler"
	"github.com/robfig/cron"
)

const (
	week string = "1 week"
	day  string = "1 day"
)

// Runs a list of all defined schedules
func Schedule() {
	c := cron.New()

	// at 7 am every day
	err := c.AddFunc("0 20 11 * * *", func() {
		log.Println("----------Running daily event schedule----------")
		EventNotify(week)
		EventNotify(day)
	})
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("Starting all schedules")
	c.Start()
}

// Filters events within the provided interval and sends a notification to the user
//
// The interval should be a valid postgres interval
func EventNotify(interval string) error {
	rows, err := datastore.DBConnection.Query(
		"Select * FROM EVENTS WHERE date BETWEEN CURRENT_DATE and CURRENT_DATE + INTERVAL '$1'",
		interval,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	events := make([]*handler.Event, 10)

	for rows.Next() {
		event := new(handler.Event)
		err := rows.Scan(&event.Id, &event.Name, &event.Note, &event.Date)
		if err != nil {
			return err
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	fmt.Println("Events", events)

	for _, event := range events {
		payload := EmailPayload{
			Sender:    "prajwalad101@gmail.com",
			Subject:   event.Name,
			Body:      event.Note,
			Recipient: "prajwalad101@gmail.com",
		}
		SendMail(payload)
	}

	return nil
}
