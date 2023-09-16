package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mailgun/mailgun-go/v4"
	"github.com/prajwalad101/datekeeper/pkg/utils"
)

type EmailPayload struct {
	Sender    string
	Subject   string
	Body      string
	Recipient string
}

func SendMail(payload EmailPayload) {
	env := utils.GetEnv()

	mg := mailgun.NewMailgun(env.MailgunDomain, env.MailgunKey)

	message := mg.NewMessage(payload.Sender, payload.Subject, payload.Body, payload.Recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// send a message with 10 second timeout
	resp, id, err := mg.Send(ctx, message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
