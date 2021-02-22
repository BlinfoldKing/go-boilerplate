package assetwarehouse

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	"go-boilerplate/modules/asset"
	"go-boilerplate/modules/warehouse"
)

// Service contains business logic
type Service struct {
	repository Repository
	asset      asset.Service
	warehouse  warehouse.Service
}

// InitService create new service
func InitService(adapters adapters.Adapters) Service {
	repo := CreatePosgresRepository(adapters.Postgres)
	asset := asset.InitAssetService(adapters)
	warehouse := warehouse.InitWarehouseService(adapters)
	return CreateService(repo, asset, warehouse)
}

// CreateService init service
func CreateService(repo Repository, asset asset.Service, warehouse warehouse.Service) Service {
	return Service{repo, asset, warehouse}
}

func (service Service) mapAssetWarehouseToAssetWarehouseGroup(assetwarehouse entity.AssetWarehouse) (ag entity.AssetWarehouseGroup, err error) {
	asset, err := service.asset.GetByID(assetwarehouse.AssetID)
	if err != nil {
		return
	}

	warehouse, err := service.warehouse.GetByID(assetwarehouse.WarehouseID)
	if err != nil {
		return
	}

	ag = entity.AssetWarehouseGroup{
		AssetWarehouse: assetwarehouse,
		Asset:          asset,
		Warehouse:      warehouse.Warehouse,
	}
	return
}

// CreateAssetWarehouse create new assetWarehouse
func (service Service) CreateAssetWarehouse(assetID string, warehouseID string) (assetWarehouse entity.AssetWarehouse, err error) {
	assetWarehouse, err = entity.NewAssetWarehouse(assetID, warehouseID)
	if err != nil {
		return
	}
	err = service.repository.Save(assetWarehouse)
	return
}

// GetList get list of assetWarehouse
func (service Service) GetList(pagination entity.Pagination) (assetWarehouse []entity.AssetWarehouseGroup, count int, err error) {
	aws, count, err := service.repository.GetList(pagination)
	if err != nil {
		return
	}

	for _, aw := range aws {
		ag, err := service.mapAssetWarehouseToAssetWarehouseGroup(aw)
		if err != nil {
			return assetWarehouse, count, err
		}

		assetWarehouse = append(assetWarehouse, ag)
	}
	return
}

// Update update assetWarehouse
func (service Service) Update(id string, changeset entity.AssetWarehouseChangeSet) (assetWarehouse entity.AssetWarehouseGroup, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return
	}
	return service.GetByID(id)
}

// GetByID find assetWarehouseby id
func (service Service) GetByID(id string) (assetWarehouse entity.AssetWarehouseGroup, err error) {
	aw, err := service.repository.FindByID(id)
	if err != nil {
		return
	}

	return service.mapAssetWarehouseToAssetWarehouseGroup(aw)
}

// DeleteByID delete assetWarehouseby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
