package templatesinvolvedid

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.TemplatesInvolvedID) error
	DeleteByID(id string) error
	FindByID(id string) (entity.TemplatesInvolvedID, error)
	Update(id string, changeset entity.TemplatesInvolvedIDChangeSet) error
	GetList(pagination entity.Pagination) (TemplatesInvolvedIDs []entity.TemplatesInvolvedID, count int, err error)
}
