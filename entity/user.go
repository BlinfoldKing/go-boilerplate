package entity

import (
	"github.com/satori/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User user entity
type User struct {
	ID           string `xorm:"id"`
	Email        string `xorm:"email"`
	PasswordHash string `xorm:"password_hash"`
}

// NewUser create new user
func NewUser(email, password string) (user User, err error) {
	id := uuid.NewV4().String()
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	user = User{
		ID:           id,
		Email:        email,
		PasswordHash: string(bytes),
	}

	return
}

// ComparePassword Compare current password hash and a password
func (user User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	return err == nil
}
