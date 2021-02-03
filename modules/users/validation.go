package users

// UpdateRequest  request for update
type UpdateRequest struct {
	Email            string `json:"email" validate:"email,required"`
	CompanyContactID string `json:"company_contact_id"`
}
