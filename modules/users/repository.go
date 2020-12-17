package users

import "go-boilerplate/entity"

// Repository abstraction for storage
type Repository interface {
	Save(entity.User) error
	FindByEmail(email string) (entity.User, error)

	Update(id string, changeset entity.UserChangeSet) error
	GetList(entity.Pagination) ([]entity.User, error)
	FindByID(id string) (entity.User, error)
	DeleteByID(id string) error
}
