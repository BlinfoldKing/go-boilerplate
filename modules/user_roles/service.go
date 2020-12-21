package userroles

import (
	"go-boilerplate/entity"
)

// Service contains business logic for roles
type Service struct {
	repository Repository
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateRole create new role
func (service Service) CreateRole(userID, roleID string) (role entity.UserRole, err error) {
	role, err = entity.NewUserRole(userID, roleID)
	if err != nil {
		return
	}

	err = service.repository.Save(role)

	return
}

// GetAllByUserID create new role
func (service Service) GetAllByUserID(id string) (role []entity.UserRole, err error) {
	role, err = service.repository.GetAllByUserID(id)
	return
}

// GetByID find role by id
func (service Service) GetByID(id string) (role entity.UserRole, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete role by id
func (service Service) DeleteByID(id string) error {
	return service.repository.DeleteByID(id)
}
