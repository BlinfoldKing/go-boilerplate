package site

import "go-boilerplate/entity"

// CreateRequest request for create new site
type CreateRequest struct {
	Name        string                   `json:"name" validate:"required"`
	Latitude    float32                  `json:"latitude" validate:"required"`
	Longitude   float32                  `json:"longitude" validate:"required"`
	Description string                   `json:"description" validate:"required"`
	Address     string                   `json:"address" validate:"required"`
	DocumentIDs *[]string                `json:"document_ids"`
	ContactIDs  *[]entity.SiteContactIDS `json:"contact_ids"`
	AssetIDs    *[]string                `json:"asset_ids"`
}

// UpdateRequest request for update site
type UpdateRequest struct {
	Name        string  `json:"name"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
	Description string  `json:"description"`
	Address     string  `json:"address"`
}
