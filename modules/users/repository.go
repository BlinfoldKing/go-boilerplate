package users

import "go-boilerplate/entity"

type Repository interface {
	Save(entity.User) error

	FindByEmail(email string) (entity.User, error)
}
