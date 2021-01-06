package entity

import (
	"time"

	"github.com/satori/uuid"
)

// WorkOrder work_order entity
type WorkOrder struct {
	ID          string     `json:"id" xorm:"id"`
	PICID       string     `json:"pic_id" xorm:"pic_id"`
	Name        string     `json:"name" xorm:"name"`
	Type        string     `json:"type" xorm:"type"`
	Description string     `json:"description" xorm:"description"`
	CreatedAt   time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt   time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt   *time.Time `json:"deleted_at" xorm:"deleted"`
}

// WorkOrderChangeSet change set forwork_order
type WorkOrderChangeSet struct {
	PICID       string `json:"pic_id" xorm:"pic_id"`
	Name        string `json:"name" xorm:"name"`
	Type        string `json:"type" xorm:"type"`
	Description string `json:"description" xorm:"description"`
}

// NewWorkOrder create newwork_order
func NewWorkOrder(picid, name, workOrderType, description string) (workOrder WorkOrder, err error) {
	workOrder = WorkOrder{
		ID:          uuid.NewV4().String(),
		PICID:       picid,
		Name:        name,
		Type:        workOrderType,
		Description: description,
	}
	return
}
