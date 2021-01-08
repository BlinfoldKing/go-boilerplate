package warehousecontact

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
)

// Service contains business logic
type Service struct {
	repository Repository
}

// InitWarehouseContactService inits ware house contact service
func InitWarehouseContactService(adapters adapters.Adapters) Service {
	repository := CreatePostgresRepository(adapters.Postgres)
	return CreateService(repository)
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

// CreateBatchWarehouseContacts creates a batch of new WarehouseContacts
func (service Service) CreateBatchWarehouseContacts(warehouseID string, contactIDs []string) (warehouseContacts []entity.WarehouseContact, err error) {
	for _, contactID := range contactIDs {
		warehouseContact, err := entity.NewWarehouseContact(warehouseID, contactID)
		if err != nil {
			return []entity.WarehouseContact{}, err
		}
		warehouseContacts = append(warehouseContacts, warehouseContact)
	}
	err = service.repository.SaveBatch(warehouseContacts)
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

// DeleteByWarehouseID delete warehouse contact by warehouse id
func (service Service) DeleteByWarehouseID(warehouseID string) (err error) {
	return service.repository.DeleteByWarehouseID(warehouseID)
}
