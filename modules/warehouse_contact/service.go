package warehousecontact

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

// CreateWarehouseContact create new warehouse_contact
func (service Service) CreateWarehouseContact(warehouseID, contactID string) (warehouseContact entity.WarehouseContact, err error) {
	warehouseContact, err = entity.NewWarehouseContact(warehouseID, contactID)
	if err != nil {
		return
	}
	err = service.repository.Save(warehouseContact)
	return
}

// GetList get list of warehouse_contact
func (service Service) GetList(pagination entity.Pagination) (warehouseContact []entity.WarehouseContact, count int, err error) {
	warehouseContact, count, err = service.repository.GetList(pagination)
	return
}

// Update update warehouse_contact
func (service Service) Update(id string, changeset entity.WarehouseContactChangeSet) (warehouseContact entity.WarehouseContact, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.WarehouseContact{}, err
	}
	return service.GetByID(id)
}

// GetByID find warehouse_contactby id
func (service Service) GetByID(id string) (warehouseContact entity.WarehouseContact, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete warehouse_contactby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
