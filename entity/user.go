package entity

import (
	"github.com/alexedwards/argon2id"
	"github.com/satori/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User user entity
type User struct {
	ID           string `xorm:"id" json:"id"`
	Email        string `xorm:"email" json:"email"`
	PasswordHash string `xorm:"password_hash" json:"-"`
	ActiveStatus int    `xorm:"active_status" json:"-"`
}

// UserGroup user data with role mapped
type UserGroup struct {
	User
	Roles []Role `json:"roles"`
}

// UserChangeSet changeset for user
type UserChangeSet struct {
	Email string `xorm:"email" json:"email"`
}

// HashType specifiy hashing for password
type HashType = string

const (
	// ARGO2ID using argo2id
	ARGO2ID HashType = "argo2id"
	// BCRYPT using bcrypt
	BCRYPT HashType = "bcrypt"
)

const (
	// Inactive means account is inactive
	Inactive = 0
	// Active means account is active
	Active = 1
)

// UserConfig specify optional config
type UserConfig struct {
	// HashAlgo specify hash algorithm
	HashAlgo HashType
	Role     string
}

// NewUser create new user
func NewUser(email, password string, config UserConfig) (user User, err error) {
	id := uuid.NewV4().String()
	var bytes []byte
	var hash string

	if config.HashAlgo == BCRYPT {
		bytes, err = bcrypt.GenerateFromPassword([]byte(password), 14)
	} else {
		hash, err = argon2id.CreateHash(password, argon2id.DefaultParams)
		bytes = []byte(hash)
	}

	user = User{
		ID:           id,
		Email:        email,
		PasswordHash: string(bytes),
	}

	return
}

// ComparePassword Compare current password hash and a password
func (user User) ComparePassword(password string, config UserConfig) (bool, error) {
	var err error

	if config.HashAlgo == BCRYPT {
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
		return err == nil, err
	}

	match, err := argon2id.ComparePasswordAndHash(password, user.PasswordHash)
	if err != nil {
		return false, err
	}

	return match, nil

}
