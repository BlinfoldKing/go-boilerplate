package work_order_asset

// CreateRequest request for create new work_order_asset
type CreateRequest struct {
	Name string `json:"name" validate:"required"`
}

// UpdateRequest request for update work_order_asset
type UpdateRequest struct {
	Name string `json:"name" validate:"required"`
}
