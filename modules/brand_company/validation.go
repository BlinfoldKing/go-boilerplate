package brandcompany

// CreateRequest request for create new brand_company
type CreateRequest struct {
	Name string `json:"name" validate:"required"`
}

// UpdateRequest request for update brand_company
type UpdateRequest struct {
	Name string `json:"name" validate:"required"`
}
