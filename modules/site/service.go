package site

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	"go-boilerplate/modules/asset"
	"go-boilerplate/modules/contact"
	"go-boilerplate/modules/documents"
	siteasset "go-boilerplate/modules/site_asset"
	sitecontact "go-boilerplate/modules/site_contact"
	sitedocument "go-boilerplate/modules/site_document"
)

// Service contains business logic
type Service struct {
	repository    Repository
	documents     documents.Service
	contacts      contact.Service
	assets        asset.Service
	siteContacts  sitecontact.Service
	siteDocuments sitedocument.Service
	siteAssets    siteasset.Service
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

func (service Service) mapSiteToSiteGroup(site entity.Site) (siteGroup entity.SiteGroup, err error) {
	siteGroup.Site = site
	siteGroup.Documents, err = service.documents.GetBySiteID(site.ID)
	if err != nil {
		return
	}

	siteGroup.Contact, err = service.contacts.GetBySiteID(site.ID)
	if err != nil {
		return
	}

	siteGroup.Assets, err = service.assets.GetBySiteID(site.ID)
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
	documentIDs *[]string,
	contactIDs *[]entity.SiteContactIDS,
	assetIDs *[]string,
) (site entity.Site, err error) {
	site, err = entity.NewSite(name, latitude, longitude, description, address)
	if err != nil {
		return
	}
	err = service.repository.Save(site)

	if documentIDs != nil {
		_, err = service.siteDocuments.CreateBatchSiteDocument(site.ID, *documentIDs)
	}

	if contactIDs != nil {
		_, err = service.siteContacts.CreateBatchSiteContact(site.ID, *contactIDs)
	}

	if assetIDs != nil {
		_, err = service.siteAssets.CreateBatchSiteAsset(site.ID, *assetIDs)
	}

	return
}

// GetList get list of site
func (service Service) GetList(pagination entity.Pagination) (siteGroups []entity.SiteGroup, count int, err error) {
	sites, count, err := service.repository.GetList(pagination)
	if err != nil {
		return
	}
	for _, site := range sites {
		siteGroup, _ := service.mapSiteToSiteGroup(site)
		siteGroups = append(siteGroups, siteGroup)
	}
	return
}

// Update update site
func (service Service) Update(id string, changeset entity.SiteChangeSet) (site entity.SiteGroup, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.SiteGroup{}, err
	}
	return service.GetByID(id)
}

// GetByID find siteby id
func (service Service) GetByID(id string) (siteGroup entity.SiteGroup, err error) {
	site, err := service.repository.FindByID(id)
	if err != nil {
		return
	}
	siteGroup, err = service.mapSiteToSiteGroup(site)
	return
}

// DeleteByID delete siteby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
