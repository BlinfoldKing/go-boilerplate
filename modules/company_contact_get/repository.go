package companycontactget

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	FindByID(id string) (entity.CompanyContact, error)
	GetList(pagination entity.Pagination) (CompanyContacts []entity.CompanyContact, count int, err error)
}
