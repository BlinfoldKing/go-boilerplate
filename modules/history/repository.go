package history

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.History) error
	DeleteByID(id string) error
	FindByID(id string) (entity.History, error)
	Update(id string, changeset entity.HistoryChangeSet) error
	GetList(pagination entity.Pagination) (Histories []entity.History, count int, err error)
}
