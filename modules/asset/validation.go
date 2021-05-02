package asset

import (
	"go-boilerplate/entity"
	"time"
)

// CreateRequest request for create new asset
type CreateRequest struct {
	ProductID         string    `json:"product_id" validate:"required"`
	SerialNumber      string    `json:"serial_number" validate:"required"`
	Status            int       `json:"status" validate:"min=0,max=10"`
	PurchaseDate      time.Time `json:"purchase_date" validate:"required"`
	PurchasePrice     float32   `json:"purchase_price" validate:"required"`
	SupplierCompanyID string    `json:"supplier_company_id" validate:"required"`
	SalvageValue      float32   `json:"salvage_value" validate:"required"`
	CreatedBy         *string   `json:"created_by"`
}

// UpdateRequest request for update asset
type UpdateRequest struct {
	entity.AssetChangeSet
	WarehouseIDs []string `json:"warehouse_ids"`
}
