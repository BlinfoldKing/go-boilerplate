package mailgun

import (
	"go-boilerplate/config"

	"github.com/mailgun/mailgun-go"
)

// Init creates mailgun client
func Init() *mailgun.MailgunImpl {
	domain := config.MAILGUNDOMAIN()
	apiKey := config.MAILGUNAPIKEY()

	client := mailgun.NewMailgun(domain, apiKey)
	return client
}
