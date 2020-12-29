package brandcompany

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.BrandCompany) error
	SaveBatch([]entity.BrandCompany) error
	DeleteByID(id string) error
	DeleteByBrandID(brandID string) error
	FindByID(id string) (entity.BrandCompany, error)
	Update(id string, changeset entity.BrandCompanyChangeSet) error
	GetList(pagination entity.Pagination) (BrandCompanies []entity.BrandCompany, count int, err error)
}
