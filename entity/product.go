package entity

import (
	"github.com/satori/uuid"
	"time"
)

// Product product entity
type Product struct {
	ID                  string    `json:"id" xorm:"id"`
	Name                string    `json:"name" xorm:"name"`
	BrandID             string    `json:"brand_id" xorm:"brand_id"`
	ProductCategoryID   string    `json:"product_category_id" xorm:"product_category_id"`
	Type                string    `json:"type" xorm:"type"`
	Lifetime            time.Time `json:"lifetime" xorm:"lifetime"`
	MaintenanceInterval int       `json:"maintenance_interval" xorm:"maintenance_interval"`
}

// ProductChangeSet change set forproduct
type ProductChangeSet struct {
	Name                string    `json:"name" xorm:"name"`
	BrandID             string    `json:"brand_id" xorm:"brand_id"`
	ProductCategoryID   string    `json:"product_category_id" xorm:"product_category_id"`
	Type                string    `json:"type" xorm:"type"`
	Lifetime            time.Time `json:"lifetime" xorm:"lifetime"`
	MaintenanceInterval int       `json:"maintenance_interval" xorm:"maintenance_interval"`
}

// NewProduct create newproduct
func NewProduct(
	name string,
	brandID string,
	productCategoryID string,
	productType string,
	lifetime time.Time,
	maintenanceInterval int,

) (product Product, err error) {
	product = Product{
		ID:                  uuid.NewV1().String(),
		Name:                name,
		BrandID:             brandID,
		ProductCategoryID:   productCategoryID,
		Type:                productType,
		Lifetime:            lifetime,
		MaintenanceInterval: maintenanceInterval,
	}
	return
}
