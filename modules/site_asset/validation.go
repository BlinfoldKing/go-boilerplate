package siteasset

// CreateRequest request for create new asset_site
type CreateRequest struct {
	AssetID string `json:"asset_id" validate:"required"`
	SiteID  string `json:"site_id" validate:"required"`
}

// UpdateRequest request for update asset_site
type UpdateRequest struct {
	AssetID string `json:"asset_id"`
	SiteID  string `json:"site_id"`
}
