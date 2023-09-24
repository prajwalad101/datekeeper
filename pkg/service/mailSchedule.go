package service

import (
	"fmt"
	"log"
	"time"

	"github.com/prajwalad101/datekeeper/datastore"
	"github.com/prajwalad101/datekeeper/model"
	"github.com/prajwalad101/datekeeper/pkg/utils"
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
	err := c.AddFunc("0 0 7 * * *", func() {
		log.Println("----------Running daily event schedule----------")
		err := EventNotify(week)
		err = EventNotify(day)

		if err != nil {
			log.Println(err.Error())
		}
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
	rows, err := datastore.DB.Query(
		"Select * FROM EVENTS WHERE EXTRACT(MONTH FROM date) = EXTRACT(MONTH FROM NOW() + $1::INTERVAL) AND EXTRACT(DAY FROM date) = EXTRACT(DAY FROM NOW() + $1::INTERVAL)",
		interval,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	events := make([]*model.Event, 0)

	for rows.Next() {
		event := new(model.Event)
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
		layout := "2006-01-02T15:04:05Z"
		parsedTime, err := time.Parse(layout, event.Date)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(parsedTime)
		date := parsedTime.Format("January 02")

		subject := fmt.Sprintf("Notification for '%s'", event.Name)

		templateVariables := map[string]string{
			"title":        event.Name,
			"receiverName": "Prajwal",
			"note":         event.Note,
			"interval":     interval,
			"date":         date,
		}

		env := utils.Env

		payload := EmailPayload{
			Sender:    env.MailSender,
			Subject:   subject,
			Body:      "",
			Recipient: "prajwalad101@gmail.com",
		}

		SendMail(payload, "event notification (datekeeper)", templateVariables)
	}

	return nil
}
