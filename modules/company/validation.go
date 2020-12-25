package company

// CreateRequest request for create new company
type CreateRequest struct {
	Name string `json:"name" validate:"required"`
}

// UpdateRequest request for update company
type UpdateRequest struct {
	Name string `json:"name" validate:"required"`
}
