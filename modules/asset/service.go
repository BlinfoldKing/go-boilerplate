package asset

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	"time"
)

// Service contains business logic
type Service struct {
	repository Repository
}

func InitAssetService(adapters adapters.Adapters) Service {
	repository := CreatePosgresRepository(adapters.Postgres)
	return CreateService(repository)
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
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

// DeleteByID delete assetby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
