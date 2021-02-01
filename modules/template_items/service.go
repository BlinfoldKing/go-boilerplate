package templateitems

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
)

// Service contains business logic
type Service struct {
	repository Repository
}

// InitTemplateItemsService create templates items service
func InitTemplateItemsService(adapters adapters.Adapters) Service {
	repository := CreatePosgresRepository(adapters.Postgres)
	return CreateService(repository)
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateTemplateItems create new templateItems
func (service Service) CreateTemplateItems(templateID, productID string, qty int) (templateItems entity.TemplateItems, err error) {
	templateItems, err = entity.NewTemplateItems(templateID, productID, qty)
	if err != nil {
		return
	}
	err = service.repository.Save(templateItems)
	return
}

// GetList get list of templateItems
func (service Service) GetList(pagination entity.Pagination) (templateItems []entity.TemplateItems, count int, err error) {
	templateItems, count, err = service.repository.GetList(pagination)
	return
}

// Update update templateItems
func (service Service) Update(id string, changeset entity.TemplateItemsChangeSet) (templateItems entity.TemplateItems, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.TemplateItems{}, err
	}
	return service.GetByID(id)
}

// GetByTemplateID get template items by template id
func (service Service) GetByTemplateID(templateID string) (templateItems []entity.TemplateItems, err error) {
	return service.repository.FindByTemplateID(templateID)
}

// GetByID find templateItems by id
func (service Service) GetByID(id string) (templateItems entity.TemplateItems, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete templateItems by id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
