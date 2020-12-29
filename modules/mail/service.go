package mail

import (
	"go-boilerplate/adapters"
	"go-boilerplate/config"
)

// Service is interface for mail service
type Service interface {
	SendEmail(sender, subject, body, recipient string) (string, error)
}

// CreateService service factory
func CreateService(adapter adapters.Adapters) Service {
	if config.MAILER() == "gomail" {
		return CreateGomailService(adapter.Gomail)
	}

	return CreateMailgunService(adapter.Mailgun)
}
