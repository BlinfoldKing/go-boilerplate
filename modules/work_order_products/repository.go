package workorderproducts

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.WorkOrderProducts) error
	DeleteByID(id string) error
	FindByID(id string) (entity.WorkOrderProducts, error)
	Update(id string, changeset entity.WorkOrderProductsChangeSet) error
	GetList(pagination entity.Pagination) (WorkOrderProductss []entity.WorkOrderProducts, count int, err error)
}
