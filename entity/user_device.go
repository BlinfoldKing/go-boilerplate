package entity

import (
	"github.com/satori/uuid"
	"time"
)

// UserDevice user_device entity
type UserDevice struct {
	ID          string     `json:"id" xorm:"id"`
	UserID      string     `json:"user_id" xorm:"user_id"`
	DeviceToken string     `json:"device_token" xorm:"device_token"`
	CreatedAt   time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt   time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt   *time.Time `json:"deleted_at" xorm:"deleted"`
}

// UserDeviceChangeSet change set foruser_device
type UserDeviceChangeSet struct {
	UserID      string `json:"user_id" xorm:"user_id"`
	DeviceToken string `json:"device_token" xorm:"device_token"`
}

// NewUserDevice create newuser_device
func NewUserDevice(userID string, deviceToken string) (device UserDevice, err error) {
	device = UserDevice{
		ID:          uuid.NewV4().String(),
		UserID:      userID,
		DeviceToken: deviceToken,
	}
	return
}
