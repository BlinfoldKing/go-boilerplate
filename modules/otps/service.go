package otps

import (
	"go-boilerplate/config"
	"go-boilerplate/entity"
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
	otp = entity.NewOTP(email, purpose, config.OTPDURATION())
	err = service.repository.Save(otp)
	return
}
