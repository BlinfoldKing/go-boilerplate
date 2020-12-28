package brand_company

import (
	"errors"
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

// CreateBrandCompany create new brand_company
func (service Service) CreateBrandCompany(name string) (brand_company entity.BrandCompany, err error) {
	brand_company, err := entity.NewBrandCompany(name)
	if err != nil {
		return
	}
	err = service.repository.Save(brand_company)
	return
}

// GetList get list of brand_company
func (service Service) GetList(pagination entity.Pagination) (brand_company []entity.BrandCompany, count int, err error) {
	brand_company, count, err = service.repository.GetList(pagination)
	return
}

// Update update brand_company
func (service Service) Update(id string, changeset entity.BrandCompanyChangeSet) (brand_company entity.BrandCompany, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.BrandCompany{}, err
	}
	return service.GetByID(id)
}

// GetByID find brand_companyby id
func (service Service) GetByID(id string) (brand_company entity.BrandCompany, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete brand_companyby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
