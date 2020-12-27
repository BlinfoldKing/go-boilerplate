package notification

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.Notification) error
	DeleteByID(id string) error
	FindByID(id string) (entity.Notification, error)
	Update(id string, changeset entity.NotificationChangeSet) error
	GetList(pagination entity.Pagination) (Notifications []entity.Notification, count int, err error)
}
