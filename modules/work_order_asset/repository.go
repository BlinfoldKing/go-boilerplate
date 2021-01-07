package workorderasset

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.WorkOrderAsset) error
	SaveBatch([]entity.WorkOrderAsset) error
	DeleteByID(id string) error
	DeleteByWorkOrderID(workOrderID string) error
	FindByID(id string) (entity.WorkOrderAsset, error)
	Update(id string, changeset entity.WorkOrderAssetChangeSet) error
	GetList(pagination entity.Pagination) (WorkOrderAssets []entity.WorkOrderAsset, count int, err error)
}
