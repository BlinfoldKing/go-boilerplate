package entity

import (
	"time"
)

// OTP entity
type OTP struct {
	Token     string     `json:"token" xorm:"token"`
	Email     string     `json:"email" xorm:"email"`
	Purpose   int        `json:"purpose" xorm:"purpose"`
	ExpiredAt time.Time  `json:"expired_at" xorm:"expired_at"`
	CreatedAt time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt *time.Time `json:"deleted_at" xorm:"deleted"`
}

const (
	// AccountActivation means the otp is for activation
	AccountActivation = 0
	// ResetPassword means the otp is for reset password
	ResetPassword = 1
)

const (
	// OTPUUID means the otp uses UUID for activation
	OTPUUID = 0
	// OTPCode means the otps uses 6 digit code for activation
	OTPCode = 1
)

// NewOTP creates new OTP
func NewOTP(token, email string, purpose int, duration time.Duration) OTP {
	exp := time.Now().Local().Add(time.Second * duration)
	return OTP{
		Token:     token,
		Email:     email,
		Purpose:   purpose,
		ExpiredAt: exp,
	}
}
