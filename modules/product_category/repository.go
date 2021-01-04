package productcategory

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.ProductCategory) error
	DeleteByID(id string) error
	FindByID(id string) (entity.ProductCategory, error)
	Update(id string, changeset entity.ProductCategoryChangeSet) error
	GetList(pagination entity.Pagination) (ProductCategories []entity.ProductCategory, count int, err error)
}
