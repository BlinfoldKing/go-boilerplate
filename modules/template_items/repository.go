package templateitems

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.TemplateItems) error
	DeleteByID(id string) error
	FindByID(id string) (entity.TemplateItems, error)
	Update(id string, changeset entity.TemplateItemsChangeSet) error
	GetList(pagination entity.Pagination) (TemplateItemss []entity.TemplateItems, count int, err error)
}
