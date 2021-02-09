package templatesinvolvedid

// CreateRequest request for create new templatesInvolvedId
type CreateRequest struct {
	UserID      string `json:"user_id" validate:"required"`
	TemplatesID string `json:"templates_id" validate:"required"`
}

// UpdateRequest request for update templatesInvolvedId
type UpdateRequest struct {
	UserID      string `json:"user_id"`
	TemplatesID string `json:"templates_id"`
}
