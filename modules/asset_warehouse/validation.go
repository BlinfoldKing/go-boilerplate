package assetwarehouse

// CreateRequest request for create new assetWarehouse
type CreateRequest struct {
	AssetID     string `json:"asset_id"`
	WarehouseID string `json:"warehouse_id"`
}

// UpdateRequest request for update assetWarehouse
type UpdateRequest struct {
	AssetID     string `json:"asset_id"`
	WarehouseID string `json:"warehouse_id"`
}
