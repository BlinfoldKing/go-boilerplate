package assetwarehouseextend

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
)

// Service contains business logic
type Service struct {
	repository Repository
}

// InitService create new service
func InitService(adapters adapters.Adapters) Service {
	repo := CreatePosgresRepository(adapters.Postgres)
	return CreateService(repo)
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateBatchAssetWarehouses creates a batch of new AssetWarehouses
func (service Service) CreateBatchAssetWarehouses(assetID string, warehouseIDs []string) (assetWarehouses []entity.AssetWarehouse, err error) {
	for _, warehouseID := range warehouseIDs {
		assetWarehouse, err := entity.NewAssetWarehouse(assetID, warehouseID)
		if err != nil {
			return []entity.AssetWarehouse{}, err
		}
		assetWarehouses = append(assetWarehouses, assetWarehouse)
	}
	err = service.repository.SaveBatch(assetWarehouses)
	return
}

// DeleteByAssetID delete asset warehouse by asset id
func (service Service) DeleteByAssetID(assetID string) (err error) {
	return service.repository.DeleteByAssetID(assetID)
}
