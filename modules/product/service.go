package product

import (
	"go-boilerplate/entity"
	"time"
)

// Service contains business logic
type Service struct {
	repository Repository
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateProduct create new product
func (service Service) CreateProduct(
	name string,
	brandID string,
	productCategoryID string,
	productType string,
	lifetime time.Time,
	maintenanceInterval int,
) (product entity.Product, err error) {
	product, err = entity.NewProduct(
		name,
		brandID,
		productCategoryID,
		productType,
		lifetime,
		maintenanceInterval,
	)
	if err != nil {
		return
	}
	err = service.repository.Save(product)
	return
}

// GetList get list of product
func (service Service) GetList(pagination entity.Pagination) (product []entity.Product, count int, err error) {
	product, count, err = service.repository.GetList(pagination)
	return
}

// Update update product
func (service Service) Update(id string, changeset entity.ProductChangeSet) (product entity.Product, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.Product{}, err
	}
	return service.GetByID(id)
}

// GetByID find productby id
func (service Service) GetByID(id string) (product entity.Product, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete productby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
