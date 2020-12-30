package productspecification

// CreateRequest request for create new productSpecification
type CreateRequest struct {
	ProductID string `json:"product_id" validate:"required"`
	Parameter string `json:"parameter" validate:"required"`
	Value     string `json:"value" validate:"required"`
}

// UpdateRequest request for update productSpecification
type UpdateRequest struct {
	Parameter string `json:"parameter" validate:"required"`
	Value     string `json:"value" validate:"required"`
}
