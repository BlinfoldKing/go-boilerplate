package site

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

// CreateSite create new site
func (service Service) CreateSite(name string, latitude float32, longitude float32, description string, address string) (site entity.Site, err error) {
	site, err = entity.NewSite(name, latitude, longitude, description, address)
	if err != nil {
		return
	}
	err = service.repository.Save(site)
	return
}

// GetList get list of site
func (service Service) GetList(pagination entity.Pagination) (site []entity.Site, count int, err error) {
	site, count, err = service.repository.GetList(pagination)
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
