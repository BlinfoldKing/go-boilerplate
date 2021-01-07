package warehouse

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

// CreateWarehouse create new warehouse
func (service Service) CreateWarehouse(name, description, address string, latitude, longitude float64) (warehouse entity.Warehouse, err error) {
	warehouse, err = entity.NewWarehouse(name, description, address, latitude, longitude)
	if err != nil {
		return
	}
	err = service.repository.Save(warehouse)
	return
}

// GetList get list of warehouse
func (service Service) GetList(pagination entity.Pagination) (warehouse []entity.Warehouse, count int, err error) {
	warehouse, count, err = service.repository.GetList(pagination)
	return
}

// Update update warehouse
func (service Service) Update(id string, changeset entity.WarehouseChangeSet) (warehouse entity.Warehouse, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.Warehouse{}, err
	}
	return service.GetByID(id)
}

// GetByID find warehouseby id
func (service Service) GetByID(id string) (warehouse entity.Warehouse, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete warehouseby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
