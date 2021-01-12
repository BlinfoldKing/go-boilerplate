package entity

import (
	"time"

	"github.com/satori/uuid"
)

// Site site entity
type Site struct {
	ID          string  `json:"id" xorm:"id"`
	Name        string  `json:"name" xorm:"name"`
	Latitude    float32 `json:"latitude" xorm:"latitude"`
	Longitude   float32 `json:"longitude" xorm:"longitude"`
	Description string  `json:"description" xorm:"description"`
	Address     string  `json:"address" xorm:"address"`

	CreatedAt time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt *time.Time `json:"deleted_at" xorm:"deleted"`
}

// SiteGroup user data with role mapped
type SiteGroup struct {
	Site
	Documents []Document `json:"documents"`
	Contact   []Contact  `json:"contacts"`
	Assets    []Asset    `json:"assets"`
}

// SiteChangeSet changeset for role
type SiteChangeSet struct {
	Name        string  `json:"name" xorm:"name"`
	Latitude    float32 `json:"latitude" xorm:"latitude"`
	Longitude   float32 `json:"longitude" xorm:"longitude"`
	Description string  `json:"description" xorm:"description"`
	Address     string  `json:"address" xorm:"address"`
}

// NewSite create newsite
func NewSite(name string, latitude float32, longitude float32, description string, address string) (site Site, err error) {
	id := uuid.NewV4().String()

	site = Site{
		ID:          id,
		Name:        name,
		Latitude:    latitude,
		Longitude:   longitude,
		Description: description,
		Address:     address,
	}
	return
}
