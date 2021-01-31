package warehouse

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.Warehouse) error
	DeleteByID(id string) error
	FindByID(id string) (entity.Warehouse, error)
	Update(id string, changeset entity.WarehouseChangeSet) error
	GetList(pagination entity.Pagination) (Warehouses []entity.Warehouse, count int, err error)
	GetAllWarehousebyAssetID(id string) (warehouse []entity.Warehouse, err error)
}
