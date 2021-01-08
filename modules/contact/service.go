package contact

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
)

// Service contains business logic
type Service struct {
	repository Repository
}

// InitContactService inits contact service
func InitContactService(adapters adapters.Adapters) Service {
	repository := CreatePostgresRepository(adapters.Postgres)
	return CreateService(repository)
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateContact create new contact
func (service Service) CreateContact(name, phoneNumber, email, address, photo string) (contact entity.Contact, err error) {
	contact, err = entity.NewContact(name, phoneNumber, email, address, photo)
	if err != nil {
		return
	}
	err = service.repository.Save(contact)
	return
}

// GetList get list of contact
func (service Service) GetList(pagination entity.Pagination) (contact []entity.Contact, count int, err error) {
	contact, count, err = service.repository.GetList(pagination)
	return
}

// Update update contact
func (service Service) Update(id string, changeset entity.ContactChangeSet) (contact entity.Contact, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.Contact{}, err
	}
	return service.GetByID(id)
}

// GetByID find contactby id
func (service Service) GetByID(id string) (contact entity.Contact, err error) {
	return service.repository.FindByID(id)
}

// GetByCompanyID finds contacts by companyID
func (service Service) GetByCompanyID(companyID string) (contacts []entity.Contact, err error) {
	return service.repository.FindByCompanyID(companyID)
}

// GetByWarehouseID finds contacts by warehouseID
func (service Service) GetByWarehouseID(warehouseID string) (contacts []entity.Contact, err error) {
	return service.repository.FindByWarehouseID(warehouseID)
}

// DeleteByID delete contactby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
