package entity

import (
	"github.com/satori/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           uuid.UUID `gorm:"id"`
	Email        string    `gorm:"email"`
	PasswordHash string    `gorm:"password_hash"`
}

func NewUser(email, password string) (user User, err error) {
	id := uuid.NewV4()
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	user = User{
		Id:           id,
		Email:        email,
		PasswordHash: string(bytes),
	}

	return
}

func (user User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	return err == nil
}
