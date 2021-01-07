package contact

// CreateRequest request for create new contact
type CreateRequest struct {
	Name        string `json:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Photo       string `json:"photo" validate:"required"`
}

// UpdateRequest request for update contact
type UpdateRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Photo       string `json:"photo"`
}
