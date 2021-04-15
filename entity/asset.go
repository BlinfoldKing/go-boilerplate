package entity

import (
	"time"

	"github.com/satori/uuid"
)

// Asset asset entity
type Asset struct {
	ID                string     `json:"id" xorm:"id"`
	ProductID         string     `json:"product_id" xorm:"product_id"`
	SerialNumber      string     `json:"serial_number" xorm:"serial_number"`
	Status            int        `json:"status" xorm:"status"`
	PurchaseDate      time.Time  `json:"purchase_date" xorm:"purchase_date"`
	PurchasePrice     float32    `json:"purchase_price" xorm:"purchase_price"`
	SupplierCompanyID string     `json:"supplier_company_id" xorm:"supplier_company_id"`
	SalvageValue      float32    `json:"salvage_value" xorm:"salvage_value"`
	CreatedBy         *string    `json:"created_by" xorm:"created_by"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at"`
}

// AssetGroup asset with mapped data
type AssetGroup struct {
	Asset
	CurrentValuation float32      `json:"current_valuation"`
	Product          ProductGroup `json:"product"`
	Warehouse        []Warehouse  `json:"warehouses"`
	Company          CompanyGroup `json:"company"`
	CreatedByUser    *User        `json:"created_by_user"`
}

// AssetChangeSet change set forasset
type AssetChangeSet struct {
	ProductID         string    `json:"-" xorm:"product_id"`
	SerialNumber      string    `json:"serial_number" xorm:"serial_number"`
	Status            int       `json:"status" xorm:"status"`
	PurchaseDate      time.Time `json:"purchase_date" xorm:"purchase_date"`
	PurchasePrice     float32   `json:"purchase_price" xorm:"purchase_price"`
	SupplierCompanyID string    `json:"-" xorm:"supplier_company_id"`
	SalvageValue      float32   `json:"salvage_value" xorm:"salvage_value"`
}

// NewAsset create newasset
func NewAsset(
	productID string,
	serialNumber string,
	status int,
	purchaseDate time.Time,
	purchasePrice float32,
	supplierCompanyID string,
	salvageValue float32,
	createdBy *string,
) (asset Asset, err error) {
	asset = Asset{
		ID:                uuid.NewV4().String(),
		ProductID:         productID,
		SerialNumber:      serialNumber,
		Status:            status,
		PurchaseDate:      purchaseDate,
		PurchasePrice:     purchasePrice,
		SupplierCompanyID: supplierCompanyID,
		SalvageValue:      salvageValue,
		CreatedBy:         createdBy,
	}
	return
}
