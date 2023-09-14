package handler

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron"
)

func Schedule() {
	fmt.Println("Running Schedule")
	c := cron.New()

	// Define your task as a cron job
	err := c.AddFunc("*/3 * * * * *", func() {
		// This function will be executed every minute
		fmt.Println("Task executed at", time.Now())
	})
	if err != nil {
		log.Println(err.Error())
	}

	c.Start()
}
