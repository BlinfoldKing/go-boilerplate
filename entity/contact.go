package entity

import (
	"time"

	"github.com/satori/uuid"
)

// Contact contact entity
type Contact struct {
	ID          string     `json:"id" xorm:"id"`
	Name        string     `json:"name" xorm:"name"`
	PhoneNumber string     `json:"phone_number" xorm:"phone_number"`
	Email       string     `json:"email" xorm:"email"`
	Address     string     `json:"address" xorm:"address"`
	Photo       string     `json:"photo" xorm:"photo"`
	CreatedAt   time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt   time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt   *time.Time `json:"deleted_at" xorm:"deleted"`
}

// ContactChangeSet change set forcontact
type ContactChangeSet struct {
	Name        string `json:"name" xorm:"name"`
	PhoneNumber string `json:"phone_number" xorm:"phone_number"`
	Email       string `json:"email" xorm:"email"`
	Address     string `json:"address" xorm:"address"`
	Photo       string `json:"photo" xorm:"photo"`
}

// NewContact create newcontact
func NewContact(name, phoneNumber, email, address, photo string) (contact Contact, err error) {
	contact = Contact{
		ID:          uuid.NewV4().String(),
		Name:        name,
		PhoneNumber: phoneNumber,
		Email:       email,
		Address:     address,
		Photo:       photo,
	}
	return
}
