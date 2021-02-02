package productdocument

// CreateRequest request for create new product_document
type CreateRequest struct {
	ProductID  string `json:"product_id" validate:"required"`
	DocumentID string `json:"document_id" validate:"required"`
}

// UpdateRequest request for update product_document
type UpdateRequest struct {
	ProductID  string `json:"product_id"`
	DocumentID string `json:"document_id"`
}
