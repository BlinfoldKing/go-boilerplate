package productspecification

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

// CreateProductSpecification create new productSpecification
func (service Service) CreateProductSpecification(productID, parameter, value string) (productSpecification entity.ProductSpecification, err error) {
	productSpecification, err = entity.NewProductSpecification(productID, parameter, value)
	if err != nil {
		return
	}
	err = service.repository.Save(productSpecification)
	return
}

// GetList get list of productSpecification
func (service Service) GetList(pagination entity.Pagination) (productSpecification []entity.ProductSpecification, count int, err error) {
	productSpecification, count, err = service.repository.GetList(pagination)
	return
}

// Update update productSpecification
func (service Service) Update(id string, changeset entity.ProductSpecificationChangeSet) (productSpecification entity.ProductSpecification, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.ProductSpecification{}, err
	}
	return service.GetByID(id)
}

// GetByID find productSpecificationby id
func (service Service) GetByID(id string) (productSpecification entity.ProductSpecification, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete productSpecificationby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
