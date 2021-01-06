package involved_user

import (
	"errors"
	"go-boilerplate/entity"
)

// Service contains business logic
type Service struct {
	repository Repository
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateInvolvedUser create new involved_user
func (service Service) CreateInvolvedUser(name string) (involved_user entity.InvolvedUser, err error) {
	involved_user, err := entity.NewInvolvedUser(name)
	if err != nil {
		return
	}
	err = service.repository.Save(involved_user)
	return
}

// GetList get list of involved_user
func (service Service) GetList(pagination entity.Pagination) (involved_user []entity.InvolvedUser, count int, err error) {
	involved_user, count, err = service.repository.GetList(pagination)
	return
}

// Update update involved_user
func (service Service) Update(id string, changeset entity.InvolvedUserChangeSet) (involved_user entity.InvolvedUser, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.InvolvedUser{}, err
	}
	return service.GetByID(id)
}

// GetByID find involved_userby id
func (service Service) GetByID(id string) (involved_user entity.InvolvedUser, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete involved_userby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
