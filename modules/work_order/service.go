package work_order

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

// CreateWorkOrder create new work_order
func (service Service) CreateWorkOrder(name string) (work_order entity.WorkOrder, err error) {
	work_order, err := entity.NewWorkOrder(name)
	if err != nil {
		return
	}
	err = service.repository.Save(work_order)
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
