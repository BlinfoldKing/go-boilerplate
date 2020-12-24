package asset

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.Asset) error
	DeleteByID(id string) error
	FindByID(id string) (entity.Asset, error)
	Update(id string, changeset entity.AssetChangeSet) error
	GetList(pagination entity.Pagination) (Assets []entity.Asset, count int, err error)
}
