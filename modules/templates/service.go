package templates

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	templateitems "go-boilerplate/modules/template_items"
)

// Service contains business logic
type Service struct {
	repository    Repository
	templateItems templateitems.Service
}

// InitTemplateService create new template
func InitTemplateService(adapters adapters.Adapters) Service {
	repository := CreatePosgresRepository(adapters.Postgres)
	templateItemService := templateitems.InitTemplateItemsService(adapters)
	return CreateService(
		repository,
		templateItemService,
	)
}

// CreateService init service
func CreateService(repo Repository, templateItemService templateitems.Service) Service {
	return Service{repo, templateItemService}
}

func (service Service) mapTemplateItemsToTemplateGroup(template entity.Templates) (templatesGroup entity.TemplatesGroup, err error) {
	templatesGroup.Templates = template

	templatesGroup.TemplateItems, err = service.templateItems.GetByTemplateID(template.ID)
	if err != nil {
		return
	}

	return
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
func (service Service) GetList(pagination entity.Pagination) (templatesGroups []entity.TemplatesGroup, count int, err error) {
	templates, count, err := service.repository.GetList(pagination)
	if err != nil {
		return
	}
	for _, template := range templates {
		templatesGroup, _ := service.mapTemplateItemsToTemplateGroup(template)
		templatesGroups = append(templatesGroups, templatesGroup)
	}
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
