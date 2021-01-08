package history

// CreateRequest request for create new history
type CreateRequest struct {
	UserID      string   `json:"user_id" validate:"required"`
	AssetID     string   `json:"asset_id" validate:"required"`
	Action      string   `json:"action" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Cost        float64  `json:"cost" validate:"required"`
	DocumentIDs []string `json:"document_ids"`
}

// UpdateRequest request for update history
type UpdateRequest struct {
	UserID      string  `json:"user_id"`
	AssetID     string  `json:"asset_id"`
	Action      string  `json:"action"`
	Description string  `json:"description"`
	Cost        float64 `json:"cost"`
}
