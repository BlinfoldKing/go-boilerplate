package siteasset

import (
	"go-boilerplate/adapters"
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

// InitSiteAssetService contains business logic
func InitSiteAssetService(adapters adapters.Adapters) Service {
	repository := CreatePosgresRepository(adapters.Postgres)

	return CreateService(
		repository,
	)
}

// CreateAssetSite create new siteAsset
func (service Service) CreateAssetSite(assetID string, siteID string) (siteAsset entity.SiteAsset, err error) {
	siteAsset, err = entity.NewSiteAsset(assetID, siteID)
	if err != nil {
		return
	}
	err = service.repository.Save(siteAsset)
	return
}

// CreateBatchSiteAsset creates a batch of new siteAssets
func (service Service) CreateBatchSiteAsset(siteID string, assetIDs []string) (siteAssets []entity.SiteAsset, err error) {
	for _, assetID := range assetIDs {
		siteAsset, err := entity.NewSiteAsset(siteID, assetID)
		if err != nil {
			return []entity.SiteAsset{}, err
		}
		siteAssets = append(siteAssets, siteAsset)
	}
	err = service.repository.SaveBatch(siteAssets)
	return
}

// GetList get list of siteAsset
func (service Service) GetList(pagination entity.Pagination) (siteAsset []entity.SiteAsset, count int, err error) {
	siteAsset, count, err = service.repository.GetList(pagination)
	return
}

// Update update siteAsset
func (service Service) Update(id string, changeset entity.SiteAssetChangeSet) (siteAsset entity.SiteAsset, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.SiteAsset{}, err
	}
	return service.GetByID(id)
}

// GetByID find siteAssetby id
func (service Service) GetByID(id string) (siteAsset entity.SiteAsset, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete siteAssetby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
