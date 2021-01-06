package work_order

// CreateRequest request for create new work_order
type CreateRequest struct {
	Name string `json:"name" validate:"required"`
}

// UpdateRequest request for update work_order
type UpdateRequest struct {
	Name string `json:"name" validate:"required"`
}
