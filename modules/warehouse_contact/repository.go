package warehousecontact

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.WarehouseContact) error
	SaveBatch(warehouseContacts []entity.WarehouseContact) error
	DeleteByID(id string) error
	DeleteByWarehouseID(warehouseID string) error
	FindByID(id string) (entity.WarehouseContact, error)
	Update(id string, changeset entity.WarehouseContactChangeSet) error
	GetList(pagination entity.Pagination) (WarehouseContacts []entity.WarehouseContact, count int, err error)
}
