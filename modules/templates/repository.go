package templates

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.Templates) error
	DeleteByID(id string) error
	FindByID(id string) (entity.Templates, error)
	Update(id string, changeset entity.TemplatesChangeSet) error
	GetList(pagination entity.Pagination) (Templatess []entity.Templates, count int, err error)
}
