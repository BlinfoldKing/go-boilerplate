package workorderasset

// CreateRequest request for create new workorder_asset
type CreateRequest struct {
	WorkOrderID string `json:"workorder_id" validate:"required"`
	AssetID     string `json:"asset_id" validate:"required"`
	Qty         int    `json:"qty" validate:"required"`
}

// UpdateRequest request for update workorder_asset
type UpdateRequest struct {
	WorkOrderID string `json:"workorder_id"`
	AssetID     string `json:"asset_id"`
	Qty         int    `json:"qty"`
}
