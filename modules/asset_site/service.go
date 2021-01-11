package assetsite

import (
	"go-boilerplate/entity"
)

// Service contains business logic
type Service struct {
	repository Repository
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateAssetSite create new assetSite
func (service Service) CreateAssetSite(assetID string, siteID string) (assetSite entity.AssetSite, err error) {
	assetSite, err = entity.NewAssetSite(assetID, siteID)
	if err != nil {
		return
	}
	err = service.repository.Save(assetSite)
	return
}

// GetList get list of assetSite
func (service Service) GetList(pagination entity.Pagination) (assetSite []entity.AssetSite, count int, err error) {
	assetSite, count, err = service.repository.GetList(pagination)
	return
}

// Update update assetSite
func (service Service) Update(id string, changeset entity.AssetSiteChangeSet) (assetSite entity.AssetSite, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.AssetSite{}, err
	}
	return service.GetByID(id)
}

// GetByID find assetSiteby id
func (service Service) GetByID(id string) (assetSite entity.AssetSite, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete assetSiteby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
