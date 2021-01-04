package entity

import (
	s "github.com/gosimple/slug"
	"github.com/satori/uuid"
	"time"
)

const (
	// DefaultADMIN default admin id
	DefaultADMIN = "127805c7-9d1a-4332-8786-3b988da607e2"
	// DefaultMEMBER default member id
	DefaultMEMBER = "9d02fc35-3d92-4755-bbe2-8ba99d2b57b2"
)

// Role user entity
type Role struct {
	ID          string     `xorm:"id" json:"id"`
	Slug        string     `xorm:"slug" json:"slug"`
	Description *string    `xorm:"description" json:"description"`
	CreatedAt   time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt   time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt   *time.Time `json:"deleted_at" xorm:"deleted"`
}

// RoleChangeSet changeset for role
type RoleChangeSet struct {
	Slug        string  `xorm:"slug" json:"slug"`
	Description *string `xorm:"description" json:"description"`
}

// NewRole create new role
func NewRole(slug string, description *string) (Role, error) {
	return Role{
		ID:          uuid.NewV4().String(),
		Slug:        s.Make(slug),
		Description: description,
	}, nil
}
