package entity

import (
	"time"

	"github.com/satori/uuid"
)

// Warehouse warehouse entity
type Warehouse struct {
	ID          string     `json:"id" xorm:"id"`
	Name        string     `json:"name" xorm:"name"`
	Description string     `json:"description" xorm:"description"`
	Address     string     `json:"address" xorm:"address"`
	Latitude    float64    `json:"latitude" xorm:"latitude"`
	Longitude   float64    `json:"longitude" xorm:"longitude"`
	CreatedAt   time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt   time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt   *time.Time `json:"deleted_at" xorm:"deleted"`
}

// WarehouseGroup warehouse data with mapped tables
type WarehouseGroup struct {
	Warehouse
	Contacts []Contact `json:"contacts"`
}

// WarehouseChangeSet change set forwarehouse
type WarehouseChangeSet struct {
	Name        string  `json:"name" xorm:"name"`
	Description string  `json:"description" xorm:"description"`
	Address     string  `json:"address" xorm:"address"`
	Latitude    float64 `json:"latitude" xorm:"latitude"`
	Longitude   float64 `json:"longitude" xorm:"longitude"`
}

// NewWarehouse create newwarehouse
func NewWarehouse(name, description, address string, latitude, longitude float64) (warehouse Warehouse, err error) {
	warehouse = Warehouse{
		ID:          uuid.NewV4().String(),
		Name:        name,
		Description: description,
		Address:     address,
		Latitude:    latitude,
		Longitude:   longitude,
	}
	return
}
