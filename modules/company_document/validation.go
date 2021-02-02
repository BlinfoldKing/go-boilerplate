package companydocument

// CreateRequest request for create new company_document
type CreateRequest struct {
	CompanyID  string `json:"company_id" validate:"required"`
	DocumentID string `json:"document_id" validate:"required"`
}

// UpdateRequest request for update company_document
type UpdateRequest struct {
	CompanyID  string `json:"company_id"`
	DocumentID string `json:"document_id"`
}
