package warehouse

// CreateRequest request for create new warehouse
type CreateRequest struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Address     string   `json:"address" validate:"required"`
	Latitude    float64  `json:"latitude" validate:"required"`
	Longitude   float64  `json:"longitude" validate:"required"`
	ContactIDs  []string `json:"contact_ids" validate:"required"`
}

// UpdateRequest request for update warehouse
type UpdateRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Address     string  `json:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}
