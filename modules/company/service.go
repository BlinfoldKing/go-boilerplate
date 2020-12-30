package company

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

// CreateCompany create new company
func (service Service) CreateCompany(
	name string,
	companyType entity.CompanyType,
	address string,
	phoneNumber string,
) (company entity.Company, err error) {
	company, err = entity.NewCompany(
		name,
		companyType,
		address,
		phoneNumber,
	)
	if err != nil {
		return
	}
	err = service.repository.Save(company)
	return
}

// GetList get list of company
func (service Service) GetList(pagination entity.Pagination) (company []entity.Company, count int, err error) {
	company, count, err = service.repository.GetList(pagination)
	return
}

// Update update company
func (service Service) Update(id string, changeset entity.CompanyChangeSet) (company entity.Company, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.Company{}, err
	}
	return service.GetByID(id)
}

// GetByID find companyby id
func (service Service) GetByID(id string) (company entity.Company, err error) {
	return service.repository.FindByID(id)
}

// GetByBrandID finds companies by brandID
func (service Service) GetByBrandID(brandID string) (companies []entity.Company, err error) {
	return service.repository.FindByBrandID(brandID)
}

// DeleteByID delete companyby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
