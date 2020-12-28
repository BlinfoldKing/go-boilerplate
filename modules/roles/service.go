package roles

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
func (service Service) CreateRole(slug string, description *string) (role entity.Role, err error) {
	role, err = entity.NewRole(slug, description)
	if err != nil {
		return
	}

	err = service.repository.Save(role)

	return
}

// FindBySlug create new role
func (service Service) FindBySlug(slug string) (role entity.Role, err error) {
	role, err = service.repository.FindBySlug(slug)
	return
}

// GetList get list of roles
func (service Service) GetList(pagination entity.Pagination) (roles []entity.Role, count int, err error) {
	roles, count, err = service.repository.GetList(pagination)
	return
}

// Update update role
func (service Service) Update(id string, changeset entity.RoleChangeSet) (entity.Role, error) {
	err := service.repository.Update(id, changeset)
	if err != nil {
		return entity.Role{}, err
	}
	return service.GetByID(id)
}

// GetByID find role by id
func (service Service) GetByID(id string) (role entity.Role, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete role by id
func (service Service) DeleteByID(id string) error {
	return service.repository.DeleteByID(id)
}
