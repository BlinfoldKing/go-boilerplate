package productcategory

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
)

// Service contains business logic
type Service struct {
	repository Repository
}

func InitProductCategoryService(adapters adapters.Adapters) Service {
	repository := CreatePostgresRepository(adapters.Postgres)
	return CreateService(repository)
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateProductCategory create new productCategory
func (service Service) CreateProductCategory(parentID, code, name string) (productCategory entity.ProductCategory, err error) {
	productCategory, err = entity.NewProductCategory(parentID, code, name)
	if err != nil {
		return
	}
	err = service.repository.Save(productCategory)
	return
}

// GetList get list of productCategory
func (service Service) GetList(pagination entity.Pagination) (productCategory []entity.ProductCategory, count int, err error) {
	productCategory, count, err = service.repository.GetList(pagination)
	return
}

// Update update productCategory
func (service Service) Update(id string, changeset entity.ProductCategoryChangeSet) (productCategory entity.ProductCategory, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.ProductCategory{}, err
	}
	return service.GetByID(id)
}

// GetByID find productCategoryby id
func (service Service) GetByID(id string) (productCategory entity.ProductCategory, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete productCategoryby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
