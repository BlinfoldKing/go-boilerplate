package brand

// CreateRequest request for create new brand
type CreateRequest struct {
	Name string `json:"name" validate:"required"`
}

// UpdateRequest request for update brand
type UpdateRequest struct {
	Name string `json:"name" validate:"required"`
}
