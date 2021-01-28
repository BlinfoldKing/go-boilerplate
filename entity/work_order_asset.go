package entity

import (
	"time"

	"github.com/satori/uuid"
)

// WorkOrderAsset work_order_asset entity
type WorkOrderAsset struct {
	ID          string     `json:"id" xorm:"id"`
	WorkOrderID string     `json:"work_order_id" xorm:"work_order_id"`
	AssetID     string     `json:"asset_id" xorm:"asset_id"`
	Qty         int        `json:"qty" xorm:"qty"`
	CreatedAt   time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt   time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt   *time.Time `json:"deleted_at" xorm:"deleted"`
}

// WorkOrderAssetChangeSet change set forwork_order_asset
type WorkOrderAssetChangeSet struct {
	WorkOrderID string `json:"work_order_id" xorm:"work_order_id"`
	AssetID     string `json:"asset_id" xorm:"asset_id"`
	Qty         int    `json:"qty" xorm:"qty"`
}

// NewWorkOrderAsset create newwork_order_asset
func NewWorkOrderAsset(workOrderID, assetID string, qty int) (workOrderAsset WorkOrderAsset, err error) {
	workOrderAsset = WorkOrderAsset{
		ID:          uuid.NewV4().String(),
		WorkOrderID: workOrderID,
		AssetID:     assetID,
		Qty:         qty,
	}
	return
}
