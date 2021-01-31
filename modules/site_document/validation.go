package sitedocument

// CreateRequest request for create new site_document
type CreateRequest struct {
	DocumentID string `json:"document_id" validate:"required"`
	SiteID     string `json:"site_id" validate:"required"`
}

// UpdateRequest request for update site_document
type UpdateRequest struct {
	DocumentID string `json:"document_id"`
	SiteID     string `json:"site_id"`
}
