package workorderasset

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
)

// Service contains business logic
type Service struct {
	repository Repository
}

// InitWorkOrderAssetService create new workorderasset
func InitWorkOrderAssetService(adapters adapters.Adapters) Service {
	repository := CreatePostgresRepository(adapters.Postgres)
	return CreateService(repository)
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateWorkOrderAsset create new work_order_asset
func (service Service) CreateWorkOrderAsset(workOrderID, assetID string, qty int) (workOrderAsset entity.WorkOrderAsset, err error) {
	workOrderAsset, err = entity.NewWorkOrderAsset(workOrderID, assetID, qty)
	if err != nil {
		return
	}
	err = service.repository.Save(workOrderAsset)
	return
}

// CreateBatchWorkOrderAssets creates a batch of new workorderAssets
func (service Service) CreateBatchWorkOrderAssets(workorderID string, assets []struct {
	ID  string `json:"id" validate:"required"`
	Qty int    `json:"qty" validate:"required"`
}) (workorderAssets []entity.WorkOrderAsset, err error) {
	for _, asset := range assets {
		workorderAsset, err := entity.NewWorkOrderAsset(workorderID, asset.ID, asset.Qty)
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

// GetAllByWorkorderID delete workorder_asset by workorder id
func (service Service) GetAllByWorkorderID(workorderID string) ([]entity.WorkOrderAsset, error) {
	return service.repository.GetAllByWorkorderID(workorderID)
}
