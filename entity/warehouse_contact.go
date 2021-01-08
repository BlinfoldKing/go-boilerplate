package entity

import (
	"time"

	"github.com/satori/uuid"
)

// WarehouseContact warehouse_contact entity
type WarehouseContact struct {
	ID          string     `json:"id" xorm:"id"`
	WarehouseID string     `json:"warehouse_id" xorm:"warehouse_id"`
	ContactID   string     `json:"contact_id" xorm:"contact_id"`
	CreatedAt   time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt   time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt   *time.Time `json:"deleted_at" xorm:"deleted"`
}

// WarehouseContactChangeSet change set forwarehouse_contact
type WarehouseContactChangeSet struct {
	WarehouseID string `json:"warehouse_id" xorm:"warehouse_id"`
	ContactID   string `json:"contact_id" xorm:"contact_id"`
}

// NewWarehouseContact create newwarehouse_contact
func NewWarehouseContact(warehouseID, contactID string) (warehouseContact WarehouseContact, err error) {
	warehouseContact = WarehouseContact{
		ID:          uuid.NewV4().String(),
		WarehouseID: warehouseID,
		ContactID:   contactID,
	}
	return
}
