package templates

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

// CreateTemplates create new templates
func (service Service) CreateTemplates(name, description string) (templates entity.Templates, err error) {
	templates, err = entity.NewTemplates(name, description)
	if err != nil {
		return
	}
	err = service.repository.Save(templates)
	return
}

// GetList get list of templates
func (service Service) GetList(pagination entity.Pagination) (templates []entity.Templates, count int, err error) {
	templates, count, err = service.repository.GetList(pagination)
	return
}

// Update update templates
func (service Service) Update(id string, changeset entity.TemplatesChangeSet) (templates entity.Templates, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.Templates{}, err
	}
	return service.GetByID(id)
}

// GetByID find templates by id
func (service Service) GetByID(id string) (templates entity.Templates, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete templates by id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
