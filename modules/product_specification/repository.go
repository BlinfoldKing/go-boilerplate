package productspecification

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.ProductSpecification) error
	DeleteByID(id string) error
	DeleteByProductID(productID string) error
	FindByID(id string) (entity.ProductSpecification, error)
	FindByProductID(productID string) (productSpecifications []entity.ProductSpecification, err error)
	Update(id string, changeset entity.ProductSpecificationChangeSet) error
	GetList(pagination entity.Pagination) (ProductSpecifications []entity.ProductSpecification, count int, err error)
}
