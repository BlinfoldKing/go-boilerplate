package work_order_document

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

// CreateWorkOrderDocument create new work_order_document
func (service Service) CreateWorkOrderDocument(name string) (work_order_document entity.WorkOrderDocument, err error) {
	work_order_document, err := entity.NewWorkOrderDocument(name)
	if err != nil {
		return
	}
	err = service.repository.Save(work_order_document)
	return
}

// GetList get list of work_order_document
func (service Service) GetList(pagination entity.Pagination) (work_order_document []entity.WorkOrderDocument, count int, err error) {
	work_order_document, count, err = service.repository.GetList(pagination)
	return
}

// Update update work_order_document
func (service Service) Update(id string, changeset entity.WorkOrderDocumentChangeSet) (work_order_document entity.WorkOrderDocument, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.WorkOrderDocument{}, err
	}
	return service.GetByID(id)
}

// GetByID find work_order_documentby id
func (service Service) GetByID(id string) (work_order_document entity.WorkOrderDocument, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete work_order_documentby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
