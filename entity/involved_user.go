package entity

import (
	"time"

	"github.com/satori/uuid"
)

// InvolvedUser involved_user entity
type InvolvedUser struct {
	ID          string     `json:"id" xorm:"id"`
	UserID      string     `json:"user_id" xorm:"user_id"`
	WorkOrderID string     `json:"work_order_id" xorm:"work_order_id"`
	CreatedAt   time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt   time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt   *time.Time `json:"deleted_at" xorm:"deleted"`
}

// InvolvedUserChangeSet change set forinvolved_user
type InvolvedUserChangeSet struct {
	UserID      string `json:"user_id" xorm:"user_id"`
	WorkOrderID string `json:"work_order_id" xorm:"work_order_id"`
}

// NewInvolvedUser create newinvolved_user
func NewInvolvedUser(userID, workOrderID string) (involvedUser InvolvedUser, err error) {
	involvedUser = InvolvedUser{
		ID:          uuid.NewV4().String(),
		UserID:      userID,
		WorkOrderID: workOrderID,
	}
	return
}
