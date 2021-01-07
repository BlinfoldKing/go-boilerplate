package involveduser

// CreateRequest request for create new involved_user
type CreateRequest struct {
	UserID      string `json:"user_id" validate:"required"`
	WorkOrderID string `json:"work_order_id" validate:"required"`
}

// UpdateRequest request for update involved_user
type UpdateRequest struct {
	UserID      string `json:"user_id"`
	WorkOrderID string `json:"work_order_id"`
}
