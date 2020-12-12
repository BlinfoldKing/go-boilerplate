package users

// UpdateRequest  request for update
type UpdateRequest struct {
	ID   string `json:"id" validate:"required"`
	Role string `json:"role" validate:"required"`
}
