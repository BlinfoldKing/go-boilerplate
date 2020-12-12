package users

import "go-boilerplate/entity"

// Repository abstraction for storage
type Repository interface {
	Save(entity.User) error
	FindByEmail(email string) (entity.User, error)

	GetList(limit, offset int) ([]entity.User, error)
	// Update(entity.User) error
	// FindById(id string) (entity.User, error)
	// DeleteById(id string) (entity.User, error)
}
