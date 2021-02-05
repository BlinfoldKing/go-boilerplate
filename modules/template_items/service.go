package templateitems

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	"go-boilerplate/modules/product"
)

// Service contains business logic
type Service struct {
	repository Repository
	products   product.Service
}

// InitTemplateItemsService create templates items service
func InitTemplateItemsService(adapters adapters.Adapters) Service {
	repository := CreatePosgresRepository(adapters.Postgres)
	products := product.InitProductService(adapters)
	return CreateService(repository, products)
}

// CreateService init service
func CreateService(repo Repository, products product.Service) Service {
	return Service{repo, products}
}

func (service Service) mapTemplateItemsToTemplateItemsGroup(templateItems entity.TemplateItems) (templateItemsGroup entity.TemplateItemsGroup, err error) {
	templateItemsGroup.TemplateItems = templateItems

	productGroup, err := service.products.GetByID(templateItems.ProductID)
	templateItemsGroup.Product = productGroup.Product
	return
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
func (service Service) GetList(pagination entity.Pagination) (templateItemsGroups []entity.TemplateItemsGroup, count int, err error) {
	templateItems, count, err := service.repository.GetList(pagination)
	if err != nil {
		return
	}
	for _, templateItem := range templateItems {
		templateItemsGroup, _ := service.mapTemplateItemsToTemplateItemsGroup(templateItem)
		templateItemsGroups = append(templateItemsGroups, templateItemsGroup)
	}
	return
}

// Update update templateItems
func (service Service) Update(id string, changeset entity.TemplateItemsChangeSet) (templateItems entity.TemplateItemsGroup, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.TemplateItemsGroup{}, err
	}
	return service.GetByID(id)
}

// GetByTemplateID get template items by template id
func (service Service) GetByTemplateID(templateID string) (templateItemsGroups []entity.TemplateItemsGroup, err error) {
	templateItems, err := service.repository.FindByTemplateID(templateID)
	if err != nil {
		return
	}
	for _, templateItem := range templateItems {
		templateItemsGroup, _ := service.mapTemplateItemsToTemplateItemsGroup(templateItem)
		templateItemsGroups = append(templateItemsGroups, templateItemsGroup)
	}
	return
}

// GetByID find templateItems by id
func (service Service) GetByID(id string) (templateItemsGroup entity.TemplateItemsGroup, err error) {
	templateItem, err := service.repository.FindByID(id)
	if err != nil {
		return
	}

	templateItemsGroup, _ = service.mapTemplateItemsToTemplateItemsGroup(templateItem)
	return
}

// DeleteByID delete templateItems by id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
