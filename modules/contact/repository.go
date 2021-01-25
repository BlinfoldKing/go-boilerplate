package contact

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.Contact) error
	DeleteByID(id string) error
	FindByID(id string) (entity.Contact, error)
	FindByCompanyID(companyID string) (contacts []entity.Contact, err error)
	FindByWarehouseID(warehouseID string) (contacts []entity.Contact, err error)
	FindBySiteID(siteID string) (contacts []entity.Contact, err error)
	Update(id string, changeset entity.ContactChangeSet) error
	GetList(pagination entity.Pagination) (Contacts []entity.Contact, count int, err error)
}
