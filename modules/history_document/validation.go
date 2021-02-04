package historydocument

// CreateRequest request for create new history_document
type CreateRequest struct {
	HistoryID  string `json:"history_id" validate:"required"`
	DocumentID string `json:"document_id" validate:"required"`
}

// UpdateRequest request for update history_document
type UpdateRequest struct {
	HistoryID  string `json:"history_id"`
	DocumentID string `json:"document_id"`
}
