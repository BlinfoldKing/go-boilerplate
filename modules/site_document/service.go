package sitedocument

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

// CreateSiteDocument create new siteDocument
func (service Service) CreateSiteDocument(documentID string, siteID string) (siteDocument entity.SiteDocument, err error) {
	siteDocument, err = entity.NewSiteDocument(documentID, siteID)
	if err != nil {
		return
	}
	err = service.repository.Save(siteDocument)
	return
}

// GetList get list of siteDocument
func (service Service) GetList(pagination entity.Pagination) (siteDocument []entity.SiteDocument, count int, err error) {
	siteDocument, count, err = service.repository.GetList(pagination)
	return
}

// Update update siteDocument
func (service Service) Update(id string, changeset entity.SiteDocumentChangeSet) (siteDocument entity.SiteDocument, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.SiteDocument{}, err
	}
	return service.GetByID(id)
}

// GetByID find siteDocumentby id
func (service Service) GetByID(id string) (siteDocument entity.SiteDocument, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete siteDocumentby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
