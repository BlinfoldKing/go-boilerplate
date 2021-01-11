package sitecontact

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

// CreateSiteContact create new siteContact
func (service Service) CreateSiteContact(siteID string, contactID string, position string) (siteContact entity.SiteContact, err error) {
	siteContact, err = entity.NewSiteContact(siteID, contactID, position)
	if err != nil {
		return
	}
	err = service.repository.Save(siteContact)
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
