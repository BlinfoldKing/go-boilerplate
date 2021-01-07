package work_order

import (
	"errors"
	"go-boilerplate/entity"
	"go-boilerplate/modules/asset"
	"go-boilerplate/modules/documents"
	involveduser "go-boilerplate/modules/involved_user"
	"go-boilerplate/modules/users"
	work_order_asset "go-boilerplate/modules/work_order_asset"
	work_order_document "go-boilerplate/modules/work_order_document"
)

// Service contains business logic
type Service struct {
	repository         Repository
	assets             asset.Service
	documents          documents.Service
	involvedUsers      involveduser.Service
	users              users.Service
	workOrderAssets    workorderasset.Service
	workOrderDocuments workorderdocument.Service
}

func InitWorkOrderService(adapters adapters.Adapters) Service {
	repository := CreatePostgresRepository(adapters.Postgres)

	assetService := asset.InitAssetService(adapters)
	documentService := documents.InitDocumentsService(adapters)
	involvedUserService := involveduser.InitInvolvedUserService(adapters)
	userService := users.InitUserService(adapters)
	workOrderAssetService := workorderasset.InitWorkOrderAssetService(adapters)
	workOrderDocumentService := workorderdocument.InitWorkOrderDocumentService(adapters)
	return CreateService(
		repository,
		assetService,
		documentService,
		historyDocumentService,
		userService,
	)
}

// CreateService init service
func CreateService(
	repo Repository,
	assets asset.Service,
	documents documents.Service,
	involvedUsers involveduser.Service,
	users users.Service,
) Service {
	return Service{
		repo,
		assets,
		documents,
		involvedUsers,
		users,
	}
}

// CreateWorkOrder create new work_order
func (service Service) CreateWorkOrder(
	picID,
	name,
	description string,
	workOrderType entity.WorkOrderType,
	involvedIDs,
	assetIDs,
	documentIDs []string,
) (workOrder entity.WorkOrder, err error) {
	workOrder, err := entity.NewWorkOrder(
		picID,
		name,
		description,
		workOrderType,
	)
	if err != nil {
		return
	}
	err = service.repository.Save(work_order)
	if err != nil {
		return
	}

	_, err = service.workOrderAssets.CreateBatchWorkOrderAssets(workOrder.ID, assetIDs)
	if err != nil {
		return
	}

	_, err = service.workOrderDocuments.CreateBatchWorkOrderDocuments(workOrder.ID, documentIDs)
	if err != nil {
		return
	}

	_, err = service.involvedUsers.CreateBatchInvolvedUserAssets(workOrder.ID, assetIDs)
	if err != nil {
		return
	}
	return
}

// GetList get list of work_order
func (service Service) GetList(pagination entity.Pagination) (work_order []entity.WorkOrder, count int, err error) {
	work_order, count, err = service.repository.GetList(pagination)
	return
}

// Update update work_order
func (service Service) Update(id string, changeset entity.WorkOrderChangeSet) (work_order entity.WorkOrder, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.WorkOrder{}, err
	}
	return service.GetByID(id)
}

// GetByID find work_orderby id
func (service Service) GetByID(id string) (work_order entity.WorkOrder, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete work_orderby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
