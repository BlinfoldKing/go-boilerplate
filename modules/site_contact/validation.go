package sitecontact

// CreateRequest request for create new site_contact
type CreateRequest struct {
	SiteID    string `json:"site_id" validate:"required"`
	ContactID string `json:"contact_id" validate:"required"`
	Position  string `json:"position" validate:"required"`
}

// UpdateRequest request for update site_contact
type UpdateRequest struct {
	SiteID    string `json:"site_id"`
	ContactID string `json:"contact_id"`
	Position  string `json:"position"`
}
