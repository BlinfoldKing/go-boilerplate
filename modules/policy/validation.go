package policy

// AddPolicyRequest add new policy
type AddPolicyRequest struct {
	Method string `json:"method" validate:"required"`
	Path   string `json:"path" validate:"required"`
	RoleID string `json:"role_id" validate:"required"`
}

// DeletePolicyRequest add new policy
type DeletePolicyRequest struct {
	Method string `json:"method" validate:"required"`
	Path   string `json:"path" validate:"required"`
	RoleID string `json:"role_id" validate:"required"`
}
