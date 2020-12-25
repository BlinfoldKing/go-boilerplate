package mailgun

import (
	"go-boilerplate/config"
	"go-boilerplate/helper"

	"github.com/mailgun/mailgun-go"
)

// SendEmail will send email using mailgun
func SendEmail(client *mailgun.MailgunImpl, sender, subject, body, recipient string) (string, error) {
	sender = sender + "@" + config.MAILGUNDOMAIN()

	message := client.NewMessage(sender, subject, body, recipient)
	message.SetHtml(body)
	_, id, err := client.Send(message)

	if err != nil {
		helper.Logger.Error(err)
	}

	return id, err
}
