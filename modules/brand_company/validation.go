package brandcompany

// CreateRequest request for create new brand_company
type CreateRequest struct {
	BrandID   string `json:"brand_id" validate:"required"`
	CompanyID string `json:"company_id" validate:"required"`
}

// UpdateRequest request for update brand_company
type UpdateRequest struct {
	BrandID   string `json:"brand_id"`
	CompanyID string `json:"company_id"`
}
