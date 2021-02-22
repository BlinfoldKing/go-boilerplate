package entity

import (
	"github.com/satori/uuid"
	"time"
)

// AssetWarehouse assetWarehouse entity
type AssetWarehouse struct {
	ID          string     `json:"id" xorm:"id"`
	AssetID     string     `json:"asset_id" xorm:"asset_id"`
	WarehouseID string     `json:"warehouse_id" xorm:"warehouse_id"`
	CreatedAt   time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt   time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt   *time.Time `json:"deleted_at" xorm:"deleted"`
}

// AssetWarehouseChangeSet change set forassetWarehouse
type AssetWarehouseChangeSet struct {
	ID          string     `json:"id" xorm:"id"`
	AssetID     string     `json:"asset_id" xorm:"asset_id"`
	WarehouseID string     `json:"warehouse_id" xorm:"warehouse_id"`
	CreatedAt   time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt   time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt   *time.Time `json:"deleted_at" xorm:"deleted"`
}

// AssetWarehouseGroup AssetWarehouse group
type AssetWarehouseGroup struct {
	AssetWarehouse
	Asset     Asset     `json:"asset"`
	Warehouse Warehouse `json:"warehouse"`
}

// NewAssetWarehouse create newassetWarehouse
func NewAssetWarehouse(assetid string, warehouseid string) (assetWarehouse AssetWarehouse, err error) {
	assetWarehouse = AssetWarehouse{
		ID:          uuid.NewV4().String(),
		AssetID:     assetid,
		WarehouseID: warehouseid,
	}
	return
}
