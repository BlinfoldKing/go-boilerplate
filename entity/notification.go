package entity

import (
	"time"

	"github.com/satori/uuid"
)

// NotificationType notification type
type NotificationType = int

// NotificationStatus status type
type NotificationStatus = int

const (
	// AssetMutation :nodoc
	AssetMutation NotificationType = iota
	// AssessorReport :nodoc
	AssessorReport
	// TroubleshootingNotification :nodoc
	TroubleshootingNotification
	// AuditorReport :nodoc
	AuditorReport
	// DocumentSite :nodoc
	DocumentSite
	// MaintenanceNotification :nodoc
	MaintenanceNotification
)

const (
	// Waiting :nodoc
	Waiting NotificationStatus = iota + 1
	// Approve :nodoc
	Approve
	// Ignored :nodoc
	Ignored
)

// Notification holds information of notifications
type Notification struct {
	ID        string             `xorm:"id" json:"id"`
	UserID    string             `xorm:"user_id" json:"user_id"`
	Title     string             `xorm:"title" json:"title"`
	Subtitle  string             `xorm:"subtitle" json:"subtitle"`
	URLLink   string             `xorm:"url_link" json:"url_link"`
	Body      string             `xorm:"body" json:"body"`
	Type      NotificationType   `xorm:"type" json:"type"`
	Status    NotificationStatus `xorm:"status" json:"status"`
	CreatedAt time.Time          `json:"created_at" xorm:"created"`
	UpdatedAt time.Time          `json:"updated_at" xorm:"updated"`
	DeletedAt *time.Time         `json:"deleted_at" xorm:"deleted"`
}

// NotificationChangeSet change set fornotification
type NotificationChangeSet struct {
	Title    string             `json:"title" xorm:"title"`
	Subtitle string             `xorm:"subtitle" json:"subtitle"`
	URLLink  string             `xorm:"url_link" json:"url_link"`
	Body     string             `xorm:"body" json:"body"`
	Type     NotificationType   `xorm:"type" json:"type"`
	Status   NotificationStatus `xorm:"status" json:"status"`
}

// NewNotification used to create a new notification
func NewNotification(userID, title, subtitle, urlLink, body string, notifType NotificationType, status NotificationStatus) (notification Notification) {
	id := uuid.NewV4().String()

	return Notification{
		ID:       id,
		UserID:   userID,
		Title:    title,
		Subtitle: subtitle,
		URLLink:  urlLink,
		Body:     body,
		Type:     notifType,
		Status:   status,
	}
}
