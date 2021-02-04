package userdevice

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.UserDevice) error
	DeleteByID(id string) error
	FindByID(id string) (entity.UserDevice, error)
	FindByUserID(userid string) (entity.UserDevice, error)
	Update(id string, changeset entity.UserDeviceChangeSet) error
	GetList(pagination entity.Pagination) (UserDevices []entity.UserDevice, count int, err error)
}
