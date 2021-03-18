package templates

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	templateitems "go-boilerplate/modules/template_items"
	templatesinvolvedid "go-boilerplate/modules/templates_involved_id"
	"go-boilerplate/modules/users"
)

// Service contains business logic
type Service struct {
	repository    Repository
	templateItems templateitems.Service
	involvedIDs   templatesinvolvedid.Service
	users         users.Service
}

// InitTemplateService create new template
func InitTemplateService(adapters adapters.Adapters) Service {
	repository := CreatePosgresRepository(adapters.Postgres)
	templateItemService := templateitems.InitTemplateItemsService(adapters)
	involvedIDService := templatesinvolvedid.InitTemplatesInvolvedIDService(adapters)
	users := users.InitUserService(adapters)
	return CreateService(
		repository,
		templateItemService,
		involvedIDService,
		users,
	)
}

// CreateService init service
func CreateService(
	repo Repository,
	templateItemService templateitems.Service,
	involvedIDService templatesinvolvedid.Service,
	users users.Service,
) Service {
	return Service{repo, templateItemService, involvedIDService, users}
}

func (service Service) mapTemplateItemsToTemplateGroup(template entity.Templates) (templatesGroup entity.TemplatesGroup, err error) {
	templatesGroup.Templates = template

	templatesGroup.TemplateItems, err = service.templateItems.GetByTemplateID(template.ID)
	if err != nil {
		return
	}
	templatesGroup.InvolvedIDs, err = service.users.GetByTemplatesID(template.ID)
	return
}

// CreateTemplates create new templates
func (service Service) CreateTemplates(name, description, payload string, templateItems []entity.TemplateItems, involvedIDs []string) (templates entity.Templates, err error) {
	templates, err = entity.NewTemplates(name, description, payload)
	if err != nil {
		return
	}
	err = service.repository.Save(templates)

	for _, templateItem := range templateItems {
		_, err = service.templateItems.CreateTemplateItems(
			templates.ID,
			templateItem.ProductID,
			templateItem.Qty,
		)
	}

	for _, involvedID := range involvedIDs {
		_, err = service.involvedIDs.CreateTemplatesInvolvedID(
			involvedID,
			templates.ID,
		)
	}
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
func (service Service) Update(id string, changeset entity.TemplatesChangeSet) (templates entity.TemplatesGroup, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.TemplatesGroup{}, err
	}
	return service.GetByID(id)
}

// GetByID find templates by id
func (service Service) GetByID(id string) (templateGroup entity.TemplatesGroup, err error) {
	template, err := service.repository.FindByID(id)
	if err != nil {
		return
	}

	templateGroup, _ = service.mapTemplateItemsToTemplateGroup(template)
	return
}

// DeleteByID delete templates by id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
