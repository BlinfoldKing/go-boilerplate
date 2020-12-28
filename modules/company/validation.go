package company

// CreateRequest request for create new company
type CreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Type        int    `json:"type" validate:"required"`
	Address     string `json:"address" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

// UpdateRequest request for update company
type UpdateRequest struct {
	Name        string `json:"name"`
	Type        int    `json:"type"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}
