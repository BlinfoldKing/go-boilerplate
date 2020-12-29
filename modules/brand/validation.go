package brand

// CreateRequest request for create new brand
type CreateRequest struct {
	Name          string `json:"name" validate:"required"`
	OriginCountry string `json:"origin_country" validate:"origin_country"`
}

// UpdateRequest request for update brand
type UpdateRequest struct {
	Name          string `json:"name" validate:"required"`
	OriginCountry string `json:"origin_country" validate:"origin_country"`
}
