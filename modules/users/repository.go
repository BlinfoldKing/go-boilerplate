package users

//go:generate mockgen -package users -source=repository.go -destination repository_mock.go

import "go-boilerplate/entity"

// Repository abstraction for storage
type Repository interface {
	Save(entity.User) error
	FindByEmail(email string) (entity.User, error)

	Update(id string, changeset entity.UserChangeSet) error
	GetList(entity.Pagination) (users []entity.User, count int, err error)
	FindByID(id string) (entity.User, error)
	FindByWorkOrderID(workOrderID string) (users []entity.User, err error)
	FindByTemplatesID(templatesID string) (users []entity.User, err error)
	DeleteByID(id string) error
}
