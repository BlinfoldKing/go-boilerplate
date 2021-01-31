package asset

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.Asset) error
	DeleteByID(id string) error
	FindByID(id string) (entity.Asset, error)
	FindByWorkOrderID(workOrderID string) (assets []entity.Asset, err error)
	FindBySiteID(siteID string) (assets []entity.Asset, err error)
	Update(id string, changeset entity.AssetChangeSet) error
	GetList(pagination entity.Pagination) (Assets []entity.Asset, count int, err error)
}
