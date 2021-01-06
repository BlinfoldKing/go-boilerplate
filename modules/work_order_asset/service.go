package work_order_asset

import (
	"errors"
	"go-boilerplate/entity"
)

// Service contains business logic
type Service struct {
	repository Repository
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateWorkOrderAsset create new work_order_asset
func (service Service) CreateWorkOrderAsset(name string) (work_order_asset entity.WorkOrderAsset, err error) {
	work_order_asset, err := entity.NewWorkOrderAsset(name)
	if err != nil {
		return
	}
	err = service.repository.Save(work_order_asset)
	return
}

// GetList get list of work_order_asset
func (service Service) GetList(pagination entity.Pagination) (work_order_asset []entity.WorkOrderAsset, count int, err error) {
	work_order_asset, count, err = service.repository.GetList(pagination)
	return
}

// Update update work_order_asset
func (service Service) Update(id string, changeset entity.WorkOrderAssetChangeSet) (work_order_asset entity.WorkOrderAsset, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.WorkOrderAsset{}, err
	}
	return service.GetByID(id)
}

// GetByID find work_order_assetby id
func (service Service) GetByID(id string) (work_order_asset entity.WorkOrderAsset, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete work_order_assetby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
