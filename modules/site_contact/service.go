package sitecontact

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
)

// Service contains business logic
type Service struct {
	repository Repository
}

// InitSiteContactService contains business logic
func InitSiteContactService(adapters adapters.Adapters) Service {
	repository := CreatePosgresRepository(adapters.Postgres)

	return CreateService(
		repository,
	)
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateSiteContact create new siteContact
func (service Service) CreateSiteContact(siteID string, contactID string, position string) (siteContact entity.SiteContact, err error) {
	siteContact, err = entity.NewSiteContact(siteID, contactID, position)
	if err != nil {
		return
	}
	err = service.repository.Save(siteContact)
	return
}

// CreateBatchSiteContact creates a batch of new siteContacts
func (service Service) CreateBatchSiteContact(siteID string, contactIDs []entity.SiteContactIDS) (siteContacts []entity.SiteContact, err error) {
	for _, contactID := range contactIDs {
		siteContact, err := entity.NewSiteContact(siteID, contactID.ID, contactID.Position)
		if err != nil {
			return []entity.SiteContact{}, err
		}
		siteContacts = append(siteContacts, siteContact)
	}
	err = service.repository.SaveBatch(siteContacts)
	return
}

// GetList get list of siteContact
func (service Service) GetList(pagination entity.Pagination) (siteContact []entity.SiteContact, count int, err error) {
	siteContact, count, err = service.repository.GetList(pagination)
	return
}

// Update update siteContact
func (service Service) Update(id string, changeset entity.SiteContactChangeSet) (siteContact entity.SiteContact, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.SiteContact{}, err
	}
	return service.GetByID(id)
}

// GetByID find siteContactby id
func (service Service) GetByID(id string) (siteContact entity.SiteContact, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete siteContactby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
