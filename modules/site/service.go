package site

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	"go-boilerplate/modules/asset"
	"go-boilerplate/modules/contact"
	"go-boilerplate/modules/documents"
	sitecontact "go-boilerplate/modules/site_contact"
	sitedocument "go-boilerplate/modules/site_document"
	siteasset "go-boilerplate/modules/site_asset"
)

// Service contains business logic
type Service struct {
	repository    Repository
	documents     documents.Service
	contacts      contact.Service
	assets        asset.Service
	siteContacts  sitecontact.Service
	siteDocuments sitedocument.Service
	siteAssets siteasset.Service
}

// InitSiteService contains business logic
func InitSiteService(adapters adapters.Adapters) Service {
	repository := CreatePosgresRepository(adapters.Postgres)

	contactService := contact.InitContactService(adapters)
	documentService := documents.InitDocumentsService(adapters)
	assetService := asset.InitAssetService(adapters)
	siteContactService := sitecontact.InitSiteContactService(adapters)
	siteDocumentService := sitedocument.InitSiteDocumentService(adapters)
	siteAssetService := siteasset.InitSiteAssetService(adapters)

	return CreateService(
		repository,
		documentService,
		contactService,
		assetService,
		siteContactService,
		siteDocumentService,
		siteAssetService,
	)
}

func (service Service) mapSitesToSiteGroups(sites []entity.Site) (siteGroups []entity.SiteGroup, err error) {
	for _, site := range sites {
		documents, err := service.documents.GetBySiteID(site.ID)
		if err != nil {
			return []entity.SiteGroup{}, err
		}

		contacts, err := service.contacts.GetBySiteID(site.ID)
		if err != nil {
			return []entity.SiteGroup{}, err
		}

		asset, err := service.assets.GetBySiteID(site.ID)
		if err != nil {
			return []entity.SiteGroup{}, err
		}

		siteGroup := entity.SiteGroup{
			Site:      site,
			Documents: documents,
			Contact:   contacts,
			Assets: asset,
		}

		siteGroups = append(siteGroups, siteGroup)
	}
	return
}

// CreateService init service
func CreateService(
	repo Repository,
	documentService documents.Service,
	contactService contact.Service,
	assetService asset.Service,
	siteContactService sitecontact.Service,
	siteDocumentService sitedocument.Service,
	siteAssetService siteasset.Service,
) Service {
	return Service{
		repo,
		documentService,
		contactService,
		assetService,
		siteContactService,
		siteDocumentService,
		siteAssetService,
	}
}

// CreateSite create new site
func (service Service) CreateSite(
	name string,
	latitude float32,
	longitude float32,
	description string,
	address string,
	documentIDs []string,
	contactIDs []entity.SiteContactIDS,
	assetIDs []string,
) (site entity.Site, err error) {
	site, err = entity.NewSite(name, latitude, longitude, description, address)
	if err != nil {
		return
	}
	err = service.repository.Save(site)

	_, err = service.siteDocuments.CreateBatchSiteDocument(site.ID, documentIDs)
	_, err = service.siteContacts.CreateBatchSiteContact(site.ID, contactIDs)
	_, err = service.siteAssets.CreateBatchSiteAsset(site.ID, assetIDs)
	return
}

// GetList get list of site
func (service Service) GetList(pagination entity.Pagination) (siteGroups []entity.SiteGroup, count int, err error) {
	sites, count, err := service.repository.GetList(pagination)
	if err != nil {
		return
	}
	siteGroups, err = service.mapSitesToSiteGroups(sites)
	return
}

// Update update site
func (service Service) Update(id string, changeset entity.SiteChangeSet) (site entity.Site, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.Site{}, err
	}
	return service.GetByID(id)
}

// GetByID find siteby id
func (service Service) GetByID(id string) (site entity.Site, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete siteby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
