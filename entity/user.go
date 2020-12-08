package entity

import "golang.org/x/crypto/bcrypt"

type User struct {
	Email        string
	PasswordHash string
}

func NewUser(email, password string) (user User, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	user = User{
		Email:        email,
		PasswordHash: string(bytes),
	}

	return
}

func (user User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	return err == nil
}
