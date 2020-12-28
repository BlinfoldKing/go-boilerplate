package mail

// Service is interface for mail service
type Service interface {
	SendEmail(sender, subject, body, recipient string) (string, error)
}
