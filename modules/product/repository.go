package product

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.Product) error
	DeleteByID(id string) error
	FindByID(id string) (entity.Product, error)
	Update(id string, changeset entity.ProductChangeSet) error
	GetList(pagination entity.Pagination) (Products []entity.Product, count int, err error)
}
