package users

import (
	"errors"
	"go-boilerplate/entity"
)

// Service contains business logic for users
type Service struct {
	repository Repository
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateUser create new user
func (service Service) CreateUser(email, password string) (user entity.User, err error) {
	user, err = entity.NewUser(email, password, entity.UserConfig{})
	if err != nil {
		return
	}

	err = service.repository.Save(user)

	return
}

// AuthenticateUser create new user
func (service Service) AuthenticateUser(email, password string) (user entity.User, err error) {
	user, err = service.repository.FindByEmail(email)
	if err != nil {
		return
	}

	ok, err := user.ComparePassword(password, entity.UserConfig{})
	if err != nil {
		return user, err
	}

	if !ok {
		return user, errors.New("wrong password")
	}

	return
}
