package auth

import (
	"errors"
	"fmt"
	"go-boilerplate/config"
	"go-boilerplate/entity"
	"go-boilerplate/helper"
	"go-boilerplate/modules/mail"
	"go-boilerplate/modules/otps"
	"go-boilerplate/modules/users"
	"time"
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
	path := "activation:verify"
	if token.Purpose == entity.ResetPassword {
		path = "reset-password/verify"
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
func (service Service) Register(email, password string) (userGroup entity.UserGroup, err error) {
	userGroup, err = service.users.CreateUser(email, password)
	if err != nil {
		return
	}

	if config.EMAILACTIVATION() {
		otp, err := service.otps.CreateOTP(email, entity.AccountActivation)
		if err != nil {
			return entity.UserGroup{}, err
		}
		data := map[string]interface{}{
			"name": email,
			"link": generateLink(otp, email),
		}
		template, _ := helper.GenerateHTMLTemplate("activation_email", data)

		err = mail.PublishToQueue(mail.Message{
			Sender:    "admin",
			Recipient: email,
			Subject:   "Account Activation",
			Body:      template,
		})
	}
	return
}

// RequestVerifyActivation requests activation verification
func (service Service) RequestVerifyActivation(token, email string) (err error) {
	otp, err := service.otps.GetByTokenAndEmail(token, email)
	if err != nil {
		return
	}
	now := time.Now()
	if otp.ExpiredAt.Before(now) {
		return errors.New("token is already expired")
	}
	user, err := service.users.GetByEmail(email)
	if err != nil {
		return err
	}
	userChangeSet := entity.UserChangeSet{
		Email:        email,
		ActiveStatus: entity.Active,
	}
	user, err = service.users.Update(user.User.ID, userChangeSet)
	return
}

// RequestActivateAccount requests activate account
func (service Service) RequestActivateAccount(email string) (err error) {
	_, err = service.users.GetByEmail(email)
	if err != nil {
		return
	}
	token, err := service.otps.CreateOTP(email, entity.AccountActivation)
	if err != nil {
		return
	}

	data := map[string]interface{}{
		"name": email,
		"link": generateLink(token, email),
	}
	template, _ := helper.GenerateHTMLTemplate("activation_email", data)

	err = mail.PublishToQueue(mail.Message{
		Sender:    "admin",
		Recipient: email,
		Subject:   "Account Activation",
		Body:      template,
	})

	return
}

// ResetPassword request password reset
func (service Service) ResetPassword(otp string, email string, password string) error {
	token, err := service.otps.GetByTokenAndEmail(otp, email)
	if token.Purpose != entity.ResetPassword {
		return errors.New("invalid token")
	}

	passwordHash, err := entity.GeneratePasswordHash(password, entity.ARGO2ID)
	if err != nil {
		return err
	}

	user, err := service.users.GetByEmail(email)
	if err != nil {
		return err
	}

	user, err = service.users.Update(user.ID, entity.UserChangeSet{
		PasswordHash: passwordHash,
	})

	return err
}

// RequestResetPassword request password reset
func (service Service) RequestResetPassword(email string) error {
	_, err := service.users.GetByEmail(email)
	if err != nil {
		return err
	}
	token, err := service.otps.CreateOTP(email, entity.ResetPassword)
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"name":  email,
		"token": token.Token,
	}

	template, err := helper.GenerateHTMLTemplate("reset_password", data)
	if err != nil {
		return err
	}

	err = mail.PublishToQueue(mail.Message{
		Sender:    "admin",
		Recipient: email,
		Subject:   "Reset Password",
		Body:      template,
	})
	if err != nil {
		return err
	}

	return err
}
