package work_order_document

// CreateRequest request for create new work_order_document
type CreateRequest struct {
	Name string `json:"name" validate:"required"`
}

// UpdateRequest request for update work_order_document
type UpdateRequest struct {
	Name string `json:"name" validate:"required"`
}
