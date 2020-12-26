package mail

import (
	"go-boilerplate/config"
	"go-boilerplate/helper"

	"github.com/mailgun/mailgun-go"
)

// MailgunService is implementation of mail service
type MailgunService struct {
	client *mailgun.MailgunImpl
}

// CreateMailgunService init MailgunService
func CreateMailgunService(client *mailgun.MailgunImpl) Service {
	return MailgunService{client}
}

// SendEmail will send email using mailgun
func (service MailgunService) SendEmail(sender, subject, body, recipient string) (string, error) {

	sender = sender + "@" + config.MAILGUNDOMAIN()

	message := service.client.NewMessage(sender, subject, body, recipient)
	message.SetHtml(body)
	_, id, err := service.client.Send(message)

	if err != nil {
		helper.Logger.Error(err)
	}

	return id, err
}
