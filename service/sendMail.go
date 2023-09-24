package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mailgun/mailgun-go/v4"
	"github.com/prajwalad101/datekeeper/utils"
)

type EmailPayload struct {
	Sender    string
	Subject   string
	Body      string
	Recipient string
}

func SendMail(payload EmailPayload, templateName string, templateVariables map[string]string) {
	env := utils.Env

	mg := mailgun.NewMailgun(env.MailgunDomain, env.MailgunKey)

	message := mg.NewMessage(payload.Sender, payload.Subject, payload.Body, payload.Recipient)
	message.SetTemplate(templateName)

	for k, v := range templateVariables {
		err := message.AddTemplateVariable(
			k,
			v,
		)
		if err != nil {
			log.Println(err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// send a message with 10 second timeout
	resp, _, err := mg.Send(ctx, message)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(resp)
}
