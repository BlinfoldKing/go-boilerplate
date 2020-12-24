package mailgun

import (
	"fmt"

	"go-boilerplate/config"

	"github.com/mailgun/mailgun-go"
)

type Mailgun struct {
	client *mailgun.MailgunImpl
}

// Init creates mailgun client
func Init() *Mailgun {
	domain := config.MAILGUNDOMAIN()
	apiKey := config.MAILGUNAPIKEY()

	client := mailgun.NewMailgun(domain, apiKey)
	return &Mailgun{client}
}

// MailgunSend will send email using mailgun
func (mg Mailgun) MailgunSend(sender, subject, body, recipient string) (string, error) {
	sender = sender + "@" + config.MAILGUNDOMAIN()

	message := mg.client.NewMessage(sender, subject, body, recipient)
	message.SetHtml(body)
	_, id, err := mg.client.Send(message)

	if err != nil {
		fmt.Printf(err.Error())
	}

	return id, err
}
