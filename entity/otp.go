package entity

import (
	"time"

	"github.com/satori/uuid"
)

// OTP entity
type OTP struct {
	Token     string     `json:"token" xorm:"token"`
	Email     string     `json:"email" xorm:"email"`
	Purpose   int        `json:"purpose" xorm:"purpose"`
	ExpiredAt time.Time  `json:"expired_at" xorm:"expired_at"`
	CreatedAt time.Time  `json:"created_at" xorm:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" xorm:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" xorm:"deleted_at"`
}

const (
	// AccountActivation means the otp is for activation
	AccountActivation = 0
	// ResetPassword means the otp is for reset password
	ResetPassword = 1
)

// NewOTP creates new OTP
func NewOTP(email string, purpose int, duration time.Duration) OTP {
	token := uuid.NewV4().String()
	exp := time.Now().Local().Add(time.Second * duration)
	return OTP{
		Token:     token,
		Email:     email,
		Purpose:   purpose,
		ExpiredAt: exp,
	}
}
