package sitedocument

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
)

// Service contains business logic
type Service struct {
	repository Repository
}

// InitSiteDocumentService contains business logic
func InitSiteDocumentService(adapters adapters.Adapters) Service {
	repository := CreatePosgresRepository(adapters.Postgres)

	return CreateService(
		repository,
	)
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

// CreateBatchSiteDocument creates a batch of new siteDocuments
func (service Service) CreateBatchSiteDocument(siteID string, documentIDs []string) (siteDocuments []entity.SiteDocument, err error) {
	for _, documentID := range documentIDs {
		siteDocument, err := entity.NewSiteDocument(documentID, siteID)
		if err != nil {
			return []entity.SiteDocument{}, err
		}
		siteDocuments = append(siteDocuments, siteDocument)
	}
	err = service.repository.SaveBatch(siteDocuments)
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
