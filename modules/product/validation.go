package product

// CreateRequest request for create new product
type CreateRequest struct {
	Name                string   `json:"name" validate:"required"`
	BrandID             string   `json:"brand_id" validate:"required"`
	ProductCategoryID   string   `json:"product_category_id" validate:"required"`
	Type                string   `json:"type" validate:"required"`
	Tags                []string `json:"tags" validate:"required"`
	Lifetime            int      `json:"lifetime" validate:"required"`
	MaintenanceInterval int      `json:"maintenance_interval" validate:"required"`
	SalvageValue        float32  `json:"salvage_value" validate:"required"`
	DocumentIDs         []string `json:"document_ids"`
	Specifications      []struct {
		Parameter string `json:"parameter" validate:"required"`
		Value     string `json:"value" validate:"required"`
	} `json:"specifications" validate:"required"`
}

// UpdateRequest request for update product
type UpdateRequest struct {
	Name                string   `json:"name"`
	BrandID             string   `json:"brand_id"`
	ProductCategoryID   string   `json:"product_category_id"`
	Type                string   `json:"type"`
	Tags                []string `json:"tags"`
	Lifetime            int      `json:"lifetime"`
	MaintenanceInterval int      `json:"maintenance_interval"`
	SalvageValue        float32  `json:"salvage_value"`
}
