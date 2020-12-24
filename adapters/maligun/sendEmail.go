package mailgun

import (
	"strings"
)

// SendEmail will send email
func SendEmail(senderEmail, senderName, recipientName, recipientEmail, subject, body string) {
	splitEmail := strings.Split(recipientEmail, "@")

	if splitEmail[1] == "hotmail.com" || splitEmail[1] == "outlook.com" || splitEmail[1] == "gmail.com" {
		go MailgunSend(senderEmail, subject, body, recipientEmail)
		return
	}

	go SendgridSend(senderEmail, senderName, recipientName, recipientEmail, subject, body, body)

}
