package work_order

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.WorkOrder) error
	DeleteByID(id string) error
	FindByID(id string) (entity.WorkOrder, error)
	Update(id string, changeset entity.WorkOrderChangeSet) error
	GetList(pagination entity.Pagination) (WorkOrders []entity.WorkOrder, count int, err error)
}
