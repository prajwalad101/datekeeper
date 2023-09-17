package service

import (
	"fmt"
	"log"
	"time"

	"github.com/prajwalad101/datekeeper/pkg/datastore"
	"github.com/prajwalad101/datekeeper/pkg/handler"
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
	err := c.AddFunc("0 45 18 * * *", func() {
		log.Println("----------Running daily event schedule----------")
		err := EventNotify(week)
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
	rows, err := datastore.DBConnection.Query(
		"Select * FROM EVENTS WHERE date BETWEEN CURRENT_DATE and (CURRENT_DATE + $1::INTERVAL)",
		interval,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	events := make([]*handler.Event, 0)

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

	for _, event := range events {
		layout := "2006-01-02T15:04:05Z"
		parsedTime, err := time.Parse(layout, event.Date)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(parsedTime)
		date := parsedTime.Format("January 02")
		fmt.Println(date)

		subject := fmt.Sprintf("Notification for %v", date)
		content := fmt.Sprintf("You have an event for %v. Note: %s", date, event.Note)

		payload := EmailPayload{
			Sender:    utils.GetEnv().MailSender,
			Subject:   subject,
			Body:      content,
			Recipient: "prajwalad101@gmail.com",
		}
		SendMail(payload)
	}

	return nil
}
