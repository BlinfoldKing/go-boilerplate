package auth

import (
	"fmt"
	"go-boilerplate/config"
	"go-boilerplate/entity"
	"go-boilerplate/helper"
	"go-boilerplate/modules/mail"
	"go-boilerplate/modules/otps"
	"go-boilerplate/modules/users"
)

// Service service for auth
type Service struct {
	users users.Service
	otps  otps.Service
}

// CreateAuthService create new service
func CreateAuthService(
	users users.Service,
	otps otps.Service,
) Service {
	return Service{users, otps}
}

func generateLink(token entity.OTP, email string) string {
	path := "request-activation"
	if token.Purpose == entity.ResetPassword {
		path = "reset-password/request"
	}
	return fmt.Sprintf(
		"%s%s/auth/%s?token=%s&email=%s",
		config.APPURL(),
		config.PREFIX(),
		path,
		token.Token,
		email,
	)
}

// Login authenticate user
func (service Service) Login(email, password string) (entity.UserGroup, error) {
	return service.users.AuthenticateUser(email, password)
}

// Register Create new user
func (service Service) Register(email, password string) (entity.UserGroup, error) {
	return service.users.CreateUser(email, password)
}

// RequestResetPassword request password reset
func (service Service) RequestResetPassword(email string) error {
	token, err := service.otps.CreateOTP(email, entity.ResetPassword)

	data := map[string]interface{}{
		"name": email,
		"link": generateLink(token, email),
	}

	template, _ := helper.GenerateHTMLTemplate("reset_password", data)

	err = mail.PublishToQueue(mail.Message{
		Sender:    "admin",
		Recipient: email,
		Subject:   "Reset Password",
		Body:      template,
	})

	return err
}
