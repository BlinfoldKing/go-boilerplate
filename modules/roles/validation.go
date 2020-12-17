package roles

// CreateRequest create new role
type CreateRequest struct {
	Slug        string  `json:"slug" validate:"slug"`
	Description *string `json:"description" validate:"description"`
}

// UpdateRequest update new role
type UpdateRequest struct {
	ID          string  `json:"id" validate:"required"`
	Slug        string  `json:"slug" validate:"required"`
	Description *string `json:"description" validate:"required"`
}
