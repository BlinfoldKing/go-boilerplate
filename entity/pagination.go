package entity

// Pagination pagination parameters
type Pagination struct {
	Offset *int
	Limit  *int
	Sort   *map[string]string
	Where  *map[string]interface{}
}
