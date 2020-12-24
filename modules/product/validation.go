package product

// CreateRequest request for create new product
type CreateRequest struct {
	Name string `json:"name" validate:"required"`
}

// UpdateRequest request for update product
type UpdateRequest struct {
	Name string `json:"name" validate:"required"`
}
