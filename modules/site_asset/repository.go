package siteasset

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.SiteAsset) error
	SaveBatch([]entity.SiteAsset) error
	DeleteByID(id string) error
	FindByID(id string) (entity.SiteAsset, error)
	FindByAssetID(id string) (entity.SiteAsset, error)
	Update(id string, changeset entity.SiteAssetChangeSet) error
	GetList(pagination entity.Pagination) (AssetSites []entity.SiteAsset, count int, err error)
}
