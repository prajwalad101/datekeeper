package handler

import (
	"fmt"
	"log"

	"github.com/prajwalad101/datekeeper/pkg/datastore"
	"github.com/prajwalad101/datekeeper/pkg/service"
	"github.com/robfig/cron"
)

func Schedule() {
	fmt.Println("Running Schedule")
	c := cron.New()

	// at 7 am every day
	err := c.AddFunc("0 0 7 * * *", func() { log.Println("[Job 1]Every minute job") })
	if err != nil {
		log.Println(err.Error())
	}

	c.Start()
}

// get the list of all events

// check if any date falls within a week

// if it does send a message

func ValidateEvents() error {
	rows, err := datastore.DBConnection.Query(
		"Select * FROM EVENTS WHERE event_date BETWEEN CURRENT_DATE and CURRENT_DATE + INTERVAL '1 week'",
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	events := make([]*Event, 10)

	for rows.Next() {
		event := new(Event)
		err := rows.Scan(&event.Id, &event.Name, &event.Note, &event.Date)
		if err != nil {
			return err
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	for _, event := range events {
		payload := service.EmailPayload{
			Sender:    "prajwalad101@gmail.com",
			Subject:   event.Name,
			Body:      event.Note,
			Recipient: "prajwalad101@gmail.com",
		}
		service.SendMail(payload)
	}

	return nil
}
