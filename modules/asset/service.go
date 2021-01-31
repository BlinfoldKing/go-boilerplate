package asset

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	"go-boilerplate/modules/company"
	"go-boilerplate/modules/product"
	"go-boilerplate/modules/warehouse"
	"time"
)

// Service contains business logic
type Service struct {
	repository Repository
	warehouse  warehouse.Service
	company    company.Service
	product    product.Service
}

// InitAssetService create new asset service
func InitAssetService(adapters adapters.Adapters) Service {
	repository := CreatePosgresRepository(adapters.Postgres)
	product := product.InitProductService(adapters)
	company := company.InitCompanyService(adapters)
	warehouse := warehouse.InitWarehouseService(adapters)

	return Service{
		repository: repository,
		product:    product,
		company:    company,
		warehouse:  warehouse,
	}
}
func (service Service) mapAssetToAssetGroup(asset entity.Asset) (ag entity.AssetGroup, err error) {
	ag.Asset = asset

	ag.Company, err = service.company.GetByID(ag.SupplierCompanyID)
	if err != nil {
		return
	}

	ag.Product, err = service.product.GetByID(ag.ProductID)
	if err != nil {
		return
	}

	ag.Warehouse, err = service.warehouse.GetAllByAssetID(asset.ID)
	if err != nil {
		return
	}

	return
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{
		repository: repo,
	}
}

// CreateAsset create new asset
func (service Service) CreateAsset(
	productID string,
	serialNumber string,
	status int,
	purchaseDate time.Time,
	purchasePrice float32,
	supplierCompanyID string,
) (asset entity.Asset, err error) {
	asset, err = entity.NewAsset(
		productID,
		serialNumber,
		status,
		purchaseDate,
		purchasePrice,
		supplierCompanyID,
	)
	if err != nil {
		return
	}
	err = service.repository.Save(asset)
	return
}

// GetList get list of asset
func (service Service) GetList(pagination entity.Pagination) (asset []entity.Asset, count int, err error) {
	asset, count, err = service.repository.GetList(pagination)
	return
}

// Update update asset
func (service Service) Update(id string, changeset entity.AssetChangeSet) (asset entity.Asset, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.Asset{}, err
	}
	return service.GetByID(id)
}

// GetByID find assetby id
func (service Service) GetByID(id string) (asset entity.Asset, err error) {
	return service.repository.FindByID(id)
}

// GetByWorkOrderID finds asset by work order ID
func (service Service) GetByWorkOrderID(workOrderID string) (assets []entity.Asset, err error) {
	return service.repository.FindByWorkOrderID(workOrderID)
}

// DeleteByID delete assetby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
