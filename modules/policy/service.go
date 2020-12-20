package policy

import (
	"go-boilerplate/entity"
	"go-boilerplate/modules/roles"
)

// Service policy service
type Service struct {
	repo  Repository
	roles roles.Service
}

// CreateService create new repo
func CreateService(repo Repository, roles roles.Service) Service {
	return Service{repo, roles}
}

// AddPolicy add new auth policy
func (service Service) AddPolicy(roleID, path, method string) (entity.Policy, error) {
	role, err := service.roles.GetByID(roleID)
	if err != nil {
		return entity.Policy{}, err
	}

	return service.repo.AddPolicy(entity.Policy{
		Method: method,
		Path:   path,
		Role:   role,
	})
}
