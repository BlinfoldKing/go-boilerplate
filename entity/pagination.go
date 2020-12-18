package entity

// Pagination pagination parameters
type Pagination struct {
	Type   *string
	Offset *int
	Limit  *int
	Sort   *map[string]string
	Where  *map[string]interface{}
}
