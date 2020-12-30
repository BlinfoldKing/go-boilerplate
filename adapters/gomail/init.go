package gomail

import (
	"go-boilerplate/config"
	"gopkg.in/gomail.v2"
)

// Init Create new gomail dialer
func Init() *gomail.Dialer {
	dialer := gomail.NewDialer(
		config.GOMAILSMTPHOST(),
		config.GOMAILSMTPPORT(),
		config.GOMAILAUTHEMAIL(),
		config.GOMAILAUTHPASSWORD(),
	)
	return dialer
}
