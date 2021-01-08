package companycontact

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.CompanyContact) error
	SaveBatch(companyContacts []entity.CompanyContact) error
	DeleteByID(id string) error
	DeleteByCompanyID(companyID string) error
	FindByID(id string) (entity.CompanyContact, error)
	Update(id string, changeset entity.CompanyContactChangeSet) error
	GetList(pagination entity.Pagination) (CompanyContacts []entity.CompanyContact, count int, err error)
}
