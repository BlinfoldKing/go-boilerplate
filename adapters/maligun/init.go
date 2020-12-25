package mailgun

import (
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
