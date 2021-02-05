package templates

import "go-boilerplate/entity"

// CreateRequest request for create new templates
type CreateRequest struct {
	Name          string                 `json:"name" validate:"required"`
	Description   string                 `json:"description" validate:"required"`
	TemplateItems []entity.TemplateItems `json:"template_items"`
}

// UpdateRequest request for update templates
type UpdateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
