package involved_user

// CreateRequest request for create new involved_user
type CreateRequest struct {
	Name string `json:"name" validate:"required"`
}

// UpdateRequest request for update involved_user
type UpdateRequest struct {
	Name string `json:"name" validate:"required"`
}
