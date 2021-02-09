package templatesinvolvedid

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
)

// Service contains business logic
type Service struct {
	repository Repository
}

// InitTemplatesInvolvedIDService create new template involved ID
func InitTemplatesInvolvedIDService(adapters adapters.Adapters) Service {
	repository := CreatePosgresRepository(adapters.Postgres)
	return CreateService(
		repository,
	)
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateTemplatesInvolvedID create new templatesInvolvedID
func (service Service) CreateTemplatesInvolvedID(userID, templatesID string) (templatesInvolvedID entity.TemplatesInvolvedID, err error) {
	templatesInvolvedID, err = entity.NewTemplatesInvolvedID(userID, templatesID)
	if err != nil {
		return
	}
	err = service.repository.Save(templatesInvolvedID)
	return
}

// GetList get list of templatesInvolvedID
func (service Service) GetList(pagination entity.Pagination) (templatesInvolvedID []entity.TemplatesInvolvedID, count int, err error) {
	templatesInvolvedID, count, err = service.repository.GetList(pagination)
	return
}

// Update update templatesInvolvedID
func (service Service) Update(id string, changeset entity.TemplatesInvolvedIDChangeSet) (templatesInvolvedID entity.TemplatesInvolvedID, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.TemplatesInvolvedID{}, err
	}
	return service.GetByID(id)
}

// GetByID find templatesInvolvedIDby id
func (service Service) GetByID(id string) (templatesInvolvedID entity.TemplatesInvolvedID, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete templatesInvolvedIDby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
