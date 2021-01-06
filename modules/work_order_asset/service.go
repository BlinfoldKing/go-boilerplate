package workorderasset

import (
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
func (service Service) CreateWorkOrderAsset(workOrderID, assetID string) (workOrderAsset entity.WorkOrderAsset, err error) {
	workOrderAsset, err = entity.NewWorkOrderAsset(workOrderID, assetID)
	if err != nil {
		return
	}
	err = service.repository.Save(workOrderAsset)
	return
}

// CreateBatchWorkOrderAssets creates a batch of new workorderAssets
func (service Service) CreateBatchWorkOrderAssets(workorderID string, assetIDs []string) (workorderAssets []entity.WorkOrderAsset, err error) {
	for _, assetID := range assetIDs {
		workorderAsset, err := entity.NewWorkOrderAsset(workorderID, assetID)
		if err != nil {
			return []entity.WorkOrderAsset{}, err
		}
		workorderAssets = append(workorderAssets, workorderAsset)
	}
	err = service.repository.SaveBatch(workorderAssets)
	return
}

// GetList get list of work_order_asset
func (service Service) GetList(pagination entity.Pagination) (workOrderAsset []entity.WorkOrderAsset, count int, err error) {
	workOrderAsset, count, err = service.repository.GetList(pagination)
	return
}

// Update update work_order_asset
func (service Service) Update(id string, changeset entity.WorkOrderAssetChangeSet) (workOrderAsset entity.WorkOrderAsset, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.WorkOrderAsset{}, err
	}
	return service.GetByID(id)
}

// GetByID find work_order_assetby id
func (service Service) GetByID(id string) (workOrderAsset entity.WorkOrderAsset, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete work_order_assetby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}

// DeleteByWorkOrderID delete workorder_asset by workorder id
func (service Service) DeleteByWorkOrderID(workorderID string) (err error) {
	return service.repository.DeleteByWorkOrderID(workorderID)
}
