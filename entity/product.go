package entity

import (
	"go-boilerplate/entity/common"
	"time"

	"github.com/satori/uuid"
)

// Product product entity
type Product struct {
	ID                  string        `json:"id" xorm:"id"`
	Name                string        `json:"name" xorm:"name"`
	BrandID             string        `json:"brand_id" xorm:"brand_id"`
	ProductCategoryID   string        `json:"product_category_id" xorm:"product_category_id"`
	Type                string        `json:"type" xorm:"type"`
	Tags                common.StrArr `json:"tags" xorm:"tags"`
	Lifetime            time.Time     `json:"lifetime" xorm:"lifetime"`
	MaintenanceInterval int           `json:"maintenance_interval" xorm:"maintenance_interval"`
	CreatedAt           time.Time     `json:"created_at" xorm:"created"`
	UpdatedAt           time.Time     `json:"updated_at" xorm:"updated"`
	DeletedAt           *time.Time    `json:"deleted_at" xorm:"deleted"`
}

// ProductGroup user data with role mapped
type ProductGroup struct {
	Product
	Brand          Brand                  `json:"brand"`
	Category       ProductCategory        `json:"category"`
	Documents      []Document             `json:"documents"`
	Specifications []ProductSpecification `json:"specifications"`
}

// ProductChangeSet change set forproduct
type ProductChangeSet struct {
	Name                string        `json:"name" xorm:"name"`
	BrandID             string        `json:"brand_id" xorm:"brand_id"`
	ProductCategoryID   string        `json:"product_category_id" xorm:"product_category_id"`
	Type                string        `json:"type" xorm:"type"`
	Tags                common.StrArr `json:"tags" xorm:"tags"`
	Lifetime            time.Time     `json:"lifetime" xorm:"lifetime"`
	MaintenanceInterval int           `json:"maintenance_interval" xorm:"maintenance_interval"`
}

// NewProduct create newproduct
func NewProduct(
	name string,
	brandID string,
	productCategoryID string,
	productType string,
	productTags []string,
	lifetime time.Time,
	maintenanceInterval int,

) (product Product, err error) {
	product = Product{
		ID:                  uuid.NewV4().String(),
		Name:                name,
		BrandID:             brandID,
		ProductCategoryID:   productCategoryID,
		Tags:                productTags,
		Type:                productType,
		Lifetime:            lifetime,
		MaintenanceInterval: maintenanceInterval,
	}
	return
}
