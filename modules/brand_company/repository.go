package brand_company

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.BrandCompany) error
	DeleteByID(id string) error
	FindByID(id string) (entity.BrandCompany, error)
	Update(id string, changeset entity.BrandCompanyChangeSet) error
	GetList(pagination entity.Pagination) (BrandCompanys []entity.BrandCompany, count int, err error)
}
