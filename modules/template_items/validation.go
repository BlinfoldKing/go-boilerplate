package templateitems

// CreateRequest request for create new template_items
type CreateRequest struct {
	TemplateID string `json:"template_id" validate:"required"`
	ProductID  string `json:"product_id" validate:"required"`
	Qty        int    `json:"qty" validate:"required"`
}

// UpdateRequest request for update template_items
type UpdateRequest struct {
	TemplateID string `json:"template_id" validate:"required"`
	ProductID  string `json:"product_id" validate:"required"`
	Qty        int    `json:"qty" validate:"required"`
}
