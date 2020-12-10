package users

import "go-boilerplate/entity"

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
