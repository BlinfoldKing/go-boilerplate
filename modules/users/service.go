package users

import "go-boilerplate/entity"

type Service struct {
	repository Repository
}

func CreateService(repo Repository) Service {
	return Service{repo}
}

func (service Service) CreateUser(email, password string) (user entity.User, err error) {
	user, err = entity.NewUser(email, password)
	if err != nil {
		return
	}

	err = service.repository.Save(user)

	return
}
