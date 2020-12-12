package users

// PaginationQuery paginatation params
type PaginationQuery struct {
	Limit  int `json:"limit" validate:"required"`
	Offset int `json:"offset" validate:"required"`
}
