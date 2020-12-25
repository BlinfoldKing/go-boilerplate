package product

import "time"

// CreateRequest request for create new product
type CreateRequest struct {
	Name                string    `json:"name" validate:"required"`
	BrandID             string    `json:"brand_id" validate:"required"`
	ProductCategoryID   string    `json:"product_category_id" validate:"required"`
	Type                string    `json:"type" validate:"required"`
	Lifetime            time.Time `json:"lifetime" validate:"required"`
	MaintenanceInterval int       `json:"maintenance_interval" validate:"required"`
}

// UpdateRequest request for update product
type UpdateRequest struct {
	Name                string    `json:"name"`
	BrandID             string    `json:"brand_id"`
	ProductCategoryID   string    `json:"product_category_id"`
	Type                string    `json:"type"`
	Lifetime            time.Time `json:"lifetime"`
	MaintenanceInterval int       `json:"maintenance_interval"`
}
