package assetwarehouseextend

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	SaveBatch(assetWarehouses []entity.AssetWarehouse) error
	DeleteByAssetID(assetID string) error
}
