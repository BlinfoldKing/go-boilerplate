package entity

import (
	"time"

	"github.com/satori/uuid"
)

// TemplatesInvolvedID templates_involved_id entity
type TemplatesInvolvedID struct {
	ID          string     `json:"id" xorm:"id"`
	UserID      string     `json:"user_id" xorm:"user_id"`
	TemplatesID string     `json:"templates_id" xorm:"templates_id"`
	CreatedAt   time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt   time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt   *time.Time `json:"deleted_at" xorm:"deleted"`
}

// TemplatesInvolvedIDChangeSet change set fortemplates_involved_id
type TemplatesInvolvedIDChangeSet struct {
	UserID      string `json:"user_id" xorm:"user_id"`
	TemplatesID string `json:"templates_id" xorm:"templates_id"`
}

// NewTemplatesInvolvedID create newtemplates_involved_id
func NewTemplatesInvolvedID(userID, templatesID string) (templatesInvolvedID TemplatesInvolvedID, err error) {
	templatesInvolvedID = TemplatesInvolvedID{
		ID:          uuid.NewV4().String(),
		UserID:      userID,
		TemplatesID: templatesID,
	}
	return
}
