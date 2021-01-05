package policy

import "go-boilerplate/entity"

// Repository policy repository
type Repository interface {
	AddPolicy(entity.Policy) (entity.Policy, error)
	ListPolicy() ([]entity.Policy, error)
	DeletePolicy(roleSlug, path, method string) error
}
