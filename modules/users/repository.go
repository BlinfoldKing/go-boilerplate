package users
//go:generate mockgen -package users -source=repository.go -destination repository_mock.go

import "go-boilerplate/entity"

// Repository abstraction for storage
type Repository interface {
	Save(entity.User) error
	FindByEmail(email string) (entity.User, error)

	Update(id string, changeset entity.UserChangeSet) error
	GetList(limit, offset int) ([]entity.User, error)
	FindByID(id string) (entity.User, error)
	DeleteByID(id string) error
}
