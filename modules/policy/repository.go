package policy

import "go-boilerplate/entity"

// Repository policy repository
type Repository interface {
	AddPolicy(entity.Policy) (entity.Policy, error)
}
