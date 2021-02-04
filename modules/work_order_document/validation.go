package workorderdocument

// CreateRequest request for create new workorder_document
type CreateRequest struct {
	WorkOrderID string `json:"workorder_id" validate:"required"`
	DocumentID  string `json:"document_id" validate:"required"`
}

// UpdateRequest request for update workorder_document
type UpdateRequest struct {
	WorkOrderID string `json:"workorder_id"`
	DocumentID  string `json:"document_id"`
}
