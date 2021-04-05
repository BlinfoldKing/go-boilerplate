package workorderproducts

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
)

// Service contains business logic
type Service struct {
	repository Repository
}

// InitWorkOrderProductService :nodoc
func InitWorkOrderProductService(adapters adapters.Adapters) Service {
	repo := CreatePosgresRepository(adapters.Postgres)

	return CreateService(repo)
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateWorkOrderProducts create new workOrderProducts
func (service Service) CreateWorkOrderProducts(workOrderid string, productid string, qty int, status int) (workOrderProducts entity.WorkOrderProducts, err error) {
	workOrderProducts, err = entity.NewWorkOrderProducts(status, workOrderid, productid, qty)
	if err != nil {
		return
	}
	err = service.repository.Save(workOrderProducts)
	return
}

// GetList get list of workOrderProducts
func (service Service) GetList(pagination entity.Pagination) (workOrderProducts []entity.WorkOrderProducts, count int, err error) {
	workOrderProducts, count, err = service.repository.GetList(pagination)
	return
}

// Update update workOrderProducts
func (service Service) Update(id string, changeset entity.WorkOrderProductsChangeSet) (workOrderProducts entity.WorkOrderProducts, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.WorkOrderProducts{}, err
	}
	return service.GetByID(id)
}

// GetByID find workOrderProductsby id
func (service Service) GetByID(id string) (workOrderProducts entity.WorkOrderProducts, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete workOrderProductsby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
