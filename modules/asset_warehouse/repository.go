package assetwarehouse

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.AssetWarehouse) error
	DeleteByID(id string) error
	FindByID(id string) (entity.AssetWarehouse, error)
	Update(id string, changeset entity.AssetWarehouseChangeSet) error
	GetList(pagination entity.Pagination) (AssetWarehouses []entity.AssetWarehouse, count int, err error)
}
