package asset

import "time"

// CreateRequest request for create new asset
type CreateRequest struct {
	ProductID         string    `json:"product_id" validate:"required"`
	SerialNumber      string    `json:"serial_number" validate:"required"`
	Status            int       `json:"status" validate:"required"`
	PurchaseDate      time.Time `json:"purchase_date" validate:"required"`
	PurchasePrice     float32   `json:"purchase_price" validate:"required"`
	SupplierCompanyID string    `json:"supplier_company_id" validate:"required"`
}

// UpdateRequest request for update asset
type UpdateRequest struct {
	ProductID         string    `json:"product_id"`
	SerialNumber      string    `json:"serial_number"`
	Status            int       `json:"status"`
	PurchaseDate      time.Time `json:"purchase_date"`
	PurchasePrice     float32   `json:"purchase_price"`
	SupplierCompanyID string    `json:"supplier_company_id"`
}
