package users

import "go-boilerplate/entity"

// Repository abstraction for storage
type Repository interface {
	Save(entity.User) error

	FindByEmail(email string) (entity.User, error)
}
