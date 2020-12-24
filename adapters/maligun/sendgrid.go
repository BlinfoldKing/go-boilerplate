package helper

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendgridSend will send email using mailgun
func SendgridSend(senderEmail, senderName, recipientName, recipientEmail, subject, plainTextContent, htmlContent string) error {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	client := sendgrid.NewSendClient(apiKey)

	domain := os.Getenv("SENDGRID_DOMAIN")
	senderEmail = senderEmail + "@" + domain

	from := mail.NewEmail(senderName, senderEmail)
	to := mail.NewEmail(recipientName, recipientEmail)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	response, err := client.Send(message)
	if err != nil {
		fmt.Println(err)
		fmt.Println(response)
		return err
	}
	return nil

}
