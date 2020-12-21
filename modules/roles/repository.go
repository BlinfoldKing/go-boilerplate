package roles

import "go-boilerplate/entity"

// Repository abstraction for storage
type Repository interface {
	Save(entity.Role) error
	FindBySlug(slug string) (entity.Role, error)

	Update(id string, changeset entity.RoleChangeSet) error
	GetList(entity.Pagination) ([]entity.Role, error)
	FindByID(id string) (entity.Role, error)
	DeleteByID(id string) error
}
