package userroles

// CreateRequest create request
type CreateRequest struct {
	UserID string `json:"user_id" validate:"required"`
	RoleID string `json:"role_id" validate:"required"`
}
