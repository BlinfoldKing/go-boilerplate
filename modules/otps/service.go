package otps

import (
	"go-boilerplate/config"
	"go-boilerplate/entity"
	"go-boilerplate/helper"

	"github.com/satori/uuid"
)

// Service contains business logic for otps
type Service struct {
	repository Repository
}

// CreateService constructor for service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateOTP create new otp
func (service Service) CreateOTP(email string, purpose int) (otp entity.OTP, err error) {
	var token string

	if config.OTPMETHOD() == "UUID" {
		token = uuid.NewV4().String()
	} else {
		token, err = service.generateSixDigitCode()
		if err != nil {
			return entity.OTP{}, err
		}
	}

	otp = entity.NewOTP(token, email, purpose, config.OTPDURATION())
	err = service.repository.Save(otp)
	return
}

// GetByToken gets otp by token
func (service Service) GetByToken(token string) (entity.OTP, error) {
	return service.repository.FindByToken(token)
}

// GetByTokenAndEmail gets otp by token and email
func (service Service) GetByTokenAndEmail(token, email string) (entity.OTP, error) {
	return service.repository.FindByTokenAndEmail(token, email)
}

// DeleteByToken delete OTP by token
func (service Service) DeleteByToken(token string) (err error) {
	return service.repository.DeleteByToken(token)
}

func (service Service) generateSixDigitCode() (token string, err error) {
	for {
		token, err = helper.GenerateOTP(6)
		if err != nil {
			return "", err
		}
		otp, err := service.repository.FindByToken(token)
		if otp.Token == "" {
			break
		} else if err != nil {
			return "", err
		}
	}
	return token, nil
}
