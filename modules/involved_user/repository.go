package involved_user

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.InvolvedUser) error
	DeleteByID(id string) error
	FindByID(id string) (entity.InvolvedUser, error)
	Update(id string, changeset entity.InvolvedUserChangeSet) error
	GetList(pagination entity.Pagination) (InvolvedUsers []entity.InvolvedUser, count int, err error)
}
