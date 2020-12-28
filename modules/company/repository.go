package company

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.Company) error
	DeleteByID(id string) error
	FindByID(id string) (entity.Company, error)
	Update(id string, changeset entity.CompanyChangeSet) error
	GetList(pagination entity.Pagination) (Companys []entity.Company, count int, err error)
}
