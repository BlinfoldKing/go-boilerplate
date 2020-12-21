package users

// UpdateRequest  request for update
type UpdateRequest struct {
	ID    string `json:"id" validate:"required"`
	Email string `json:"email" validate:"email,required"`
}
