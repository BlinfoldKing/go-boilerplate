package entity

import (
	"time"

	"github.com/satori/uuid"
)

// Notification holds information of notifications
type Notification struct {
	ID        string     `xorm:"id" json:"id"`
	UserID    string     `xorm:"user_id" json:"user_id"`
	Title     string     `xorm:"title" json:"title"`
	Subtitle  string     `xorm:"subtitle" json:"subtitle"`
	URLLink   string     `xorm:"url_link" json:"url_link"`
	Body      string     `xorm:"body" json:"body"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// NewNotification used to create a new notification
func NewNotification(userID, title, subtitle, urlLink, body string) (notification Notification) {
	id := uuid.NewV4().String()

	return Notification{
		ID:       id,
		UserID:   userID,
		Title:    title,
		Subtitle: subtitle,
		URLLink:  urlLink,
		Body:     body,
	}
}
