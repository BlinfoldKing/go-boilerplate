package mail

type Service interface {
	SendEmail(sender, subject, body, recipient string) (string, error)
}
