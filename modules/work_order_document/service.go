package workorderdocument

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
)

// Service contains business logic
type Service struct {
	repository Repository
}

func InitWorkOrderDocumentService(adapters adapters.Adapters) Service {
	repository := CreatePostgresRepository(adapters.Postgres)
	return CreateService(repository)
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateWorkOrderDocument create new work_order_document
func (service Service) CreateWorkOrderDocument(workOrderID, documentID string) (workOrderDocument entity.WorkOrderDocument, err error) {
	workOrderDocument, err = entity.NewWorkOrderDocument(workOrderID, documentID)
	if err != nil {
		return
	}
	err = service.repository.Save(workOrderDocument)
	return
}

// GetList get list of work_order_document
func (service Service) GetList(pagination entity.Pagination) (workOrderDocument []entity.WorkOrderDocument, count int, err error) {
	workOrderDocument, count, err = service.repository.GetList(pagination)
	return
}

// CreateBatchWorkOrderDocuments creates a batch of new workorderDocuments
func (service Service) CreateBatchWorkOrderDocuments(workorderID string, documentIDs []string) (workorderDocuments []entity.WorkOrderDocument, err error) {
	for _, documentID := range documentIDs {
		workorderDocument, err := entity.NewWorkOrderDocument(workorderID, documentID)
		if err != nil {
			return []entity.WorkOrderDocument{}, err
		}
		workorderDocuments = append(workorderDocuments, workorderDocument)
	}
	err = service.repository.SaveBatch(workorderDocuments)
	return
}

// Update update work_order_document
func (service Service) Update(id string, changeset entity.WorkOrderDocumentChangeSet) (workOrderDocument entity.WorkOrderDocument, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.WorkOrderDocument{}, err
	}
	return service.GetByID(id)
}

// GetByID find work_order_documentby id
func (service Service) GetByID(id string) (workOrderDocument entity.WorkOrderDocument, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete work_order_documentby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}

// DeleteByWorkOrderID delete workorder_document by workorder id
func (service Service) DeleteByWorkOrderID(workorderID string) (err error) {
	return service.repository.DeleteByWorkOrderID(workorderID)
}
