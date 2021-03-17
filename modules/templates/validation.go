package templates

import "go-boilerplate/entity"

// CreateRequest request for create new templates
type CreateRequest struct {
	Name          string                 `json:"name" validate:"required"`
	Description   string                 `json:"description" validate:"required"`
	Payload       string                 `json:"payload"`
	TemplateItems []entity.TemplateItems `json:"template_items"`
	InvolvedIDs   []string               `json:"involved_ids"`
}

// UpdateRequest request for update templates
type UpdateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Payload     string `json:"payload"`
}
