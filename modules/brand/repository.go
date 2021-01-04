package brand

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.Brand) error
	DeleteByID(id string) error
	FindByID(id string) (entity.Brand, error)
	Update(id string, changeset entity.BrandChangeSet) error
	GetList(pagination entity.Pagination) (Brands []entity.Brand, count int, err error)
}
