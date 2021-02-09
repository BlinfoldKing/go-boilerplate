package entity

import (
	"time"

	"github.com/satori/uuid"
)

// Templates templates entity
type Templates struct {
	ID          string     `json:"id" xorm:"id"`
	Name        string     `json:"name" xorm:"name"`
	Description string     `json:"description" xorm:"description"`
	CreatedAt   time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt   time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt   *time.Time `json:"deleted_at" xorm:"deleted"`
}

// TemplatesGroup templates with items mapped
type TemplatesGroup struct {
	Templates
	TemplateItems []TemplateItemsGroup `json:"items"`
	InvolvedIDs   []User               `json:"involved_ids"`
}

// TemplatesChangeSet change set fortemplates
type TemplatesChangeSet struct {
	Name        string `json:"name" xorm:"name"`
	Description string `json:"description" xorm:"description"`
}

// NewTemplates create newtemplates
func NewTemplates(name, description string) (templates Templates, err error) {
	templates = Templates{
		ID:          uuid.NewV4().String(),
		Name:        name,
		Description: description,
	}
	return
}
