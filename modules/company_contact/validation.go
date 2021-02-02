package companycontact

// CreateRequest request for create new company_contact
type CreateRequest struct {
	CompanyID string `json:"company_id" validate:"required"`
	ContactID string `json:"contact_id" validate:"required"`
}

// UpdateRequest request for update company_contact
type UpdateRequest struct {
	CompanyID string `json:"company_id"`
	ContactID string `json:"contact_id"`
}
