package productcategory

// CreateRequest request for create new productCategory
type CreateRequest struct {
	ParentID string `json:"parent_id" validate:"required"`
	Code     string `json:"code" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

// UpdateRequest request for update productCategory
type UpdateRequest struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
