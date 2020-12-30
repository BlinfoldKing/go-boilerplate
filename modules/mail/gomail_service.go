package mail

import (
	"go-boilerplate/config"

	"gopkg.in/gomail.v2"
)

// GomailService is implementation of mail service
type GomailService struct {
	client *gomail.Dialer
}

// CreateGomailService init MailgunService
func CreateGomailService(client *gomail.Dialer) Service {
	return GomailService{client}
}

// SendEmail will send email using mailgun
func (service GomailService) SendEmail(sender, subject, body, recipient string) (string, error) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", config.GOMAILAUTHEMAIL())
	mailer.SetHeader("To", recipient)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	err := service.client.DialAndSend(mailer)
	return "", err
}
