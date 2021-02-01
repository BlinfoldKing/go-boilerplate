package workorder

// CreateRequest request for create new work_order
type CreateRequest struct {
	PICID       string   `json:"pic_id" validate:"required"`
	Name        string   `json:"name" validate:"required"`
	Type        int      `json:"type"`
	Status      int      `json:"status"`
	Description string   `json:"description" validate:"required"`
	InvolvedIDs []string `json:"involved_ids"`
	Assets      []struct {
		ID  string `json:"id" validate:"required"`
		Qty int    `json:"qty" validate:"required"`
	} `json:"assets"`
	DocumentIDs []string `json:"document_ids"`
}

// UpdateRequest request for update work_order
type UpdateRequest struct {
	PICID       string `json:"pic_id"`
	Name        string `json:"name"`
	Type        int    `json:"type"`
	Status      int    `json:"status"`
	Description string `json:"description"`
}
