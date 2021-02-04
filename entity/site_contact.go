package entity

import (
	"time"

	"github.com/satori/uuid"
)

// SiteContact site_contact entity
type SiteContact struct {
	ID        string     `json:"id" xorm:"id"`
	SiteID    string     `json:"site_id" xorm:"site_id"`
	ContactID string     `json:"contact_id" xorm:"contact_id"`
	Position  string     `json:"position" xorm:"position"`
	CreatedAt time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt *time.Time `json:"deleted_at" xorm:"deleted"`
}

// SiteContactChangeSet change set forsite_contact
type SiteContactChangeSet struct {
	SiteID    string `json:"site_id" xorm:"site_id"`
	ContactID string `json:"contact_id" xorm:"contact_id"`
	Position  string `json:"position" xorm:"position"`
}

// SiteContactIDS request for
type SiteContactIDS struct {
	ID       string `json:"id"`
	Position string `json:"position"`
}

// NewSiteContact create newsite_contact
func NewSiteContact(siteID string, contactID string, position string) (siteContact SiteContact, err error) {
	id := uuid.NewV4().String()

	siteContact = SiteContact{
		ID:        id,
		SiteID:    siteID,
		ContactID: contactID,
		Position:  position,
	}
	return
}
