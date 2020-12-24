package helper

import (
	"fmt"
	"os"

	"github.com/mailgun/mailgun-go"
)

// GetMGClient will get mail gun client
func GetMGClient() mailgun.Mailgun {
	domain := os.Getenv("MG_DOMAIN")
	apiKey := os.Getenv("MG_API_KEY")

	mg := mailgun.NewMailgun(domain, apiKey)
	return mg
}

// MailgunSend will send email using mailgun
func MailgunSend(sender, subject, body, recipient string) (string, error) {
	mg := GetMGClient()

	domain := os.Getenv("MG_DOMAIN")
	sender = sender + "@" + domain

	message := mg.NewMessage(sender, subject, body, recipient)
	message.SetHtml(body)
	_, id, err := mg.Send(message)

	if err != nil {
		fmt.Printf(err.Error())
	}

	return id, err
}
