package users

// UpdateRequest  request for update
type UpdateRequest struct {
	ID    string   `json:"id" validate:"required"`
	Roles []string `json:"roles" validate:"required"`
}
