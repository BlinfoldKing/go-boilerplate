package workorder

// CreateRequest request for create new work_order
type CreateRequest struct {
	NoOrder     string    `json:"no_order" validate:"required"`
	PICID       string    `json:"pic_id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Type        int       `json:"type"`
	Status      int       `json:"status"`
	Description string    `json:"description" validate:"required"`
	InvolvedIDs *[]string `json:"involved_ids"`
	Assets      *[]struct {
		ID  string `json:"id" validate:"required"`
		Qty int    `json:"qty" validate:"required"`
	} `json:"assets"`
	DocumentIDs *[]string `json:"document_ids"`
	SiteID      *string   `json:"site_id"`
}

// UpdateRequest request for update work_order
type UpdateRequest struct {
	PICID       string  `json:"pic_id"`
	Name        string  `json:"name"`
	Type        int     `json:"type"`
	Status      int     `json:"status"`
	Description string  `json:"description"`
	SiteID      *string `json:"site_id"`
}

// ApproveRequest request for aprroval
type ApproveRequest struct {
	SiteID string `json:"site_id" validate:"required"`
}

// MutationRequest request for aprroval
type MutationRequest struct {
	NextSiteID string `json:"next_site_id" validae:"required"`
}
