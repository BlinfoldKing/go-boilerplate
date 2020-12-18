package userroles

import (
	"go-boilerplate/entity"
)

// Repository user roles repository
type Repository interface {
	Save(entity.UserRole) error
	GetAllByUserID(id string) ([]entity.UserRole, error)
	FindByID(id string) (entity.UserRole, error)
	DeleteByID(id string) error
}
