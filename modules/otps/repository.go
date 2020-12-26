package otps

import "go-boilerplate/entity"

// Repository abstraction for storage
type Repository interface {
	Save(entity.OTP) error

	FindByTokenAndEmail(token, email string) (otp entity.OTP, err error)
}
