package assetsite

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.AssetSite) error
	DeleteByID(id string) error
	FindByID(id string) (entity.AssetSite, error)
	Update(id string, changeset entity.AssetSiteChangeSet) error
	GetList(pagination entity.Pagination) (AssetSites []entity.AssetSite, count int, err error)
}
