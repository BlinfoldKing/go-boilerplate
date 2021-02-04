package userdevice

import (
	"go-boilerplate/entity"
)

// Service contains business logic
type Service struct {
	repository Repository
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateUserDevice create new user_device
func (service Service) CreateUserDevice(userid string, devicetoken string) (userdevice entity.UserDevice, err error) {
	userdevice, err = entity.NewUserDevice(userid, devicetoken)
	if err != nil {
		return
	}
	err = service.repository.Save(userdevice)
	return
}

// GetList get list of user_device
func (service Service) GetList(pagination entity.Pagination) (userdevice []entity.UserDevice, count int, err error) {
	userdevice, count, err = service.repository.GetList(pagination)
	return
}

// Update update user_device
func (service Service) Update(id string, changeset entity.UserDeviceChangeSet) (userdevice entity.UserDevice, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.UserDevice{}, err
	}
	return service.GetByID(id)
}

// GetByID find user_deviceby id
func (service Service) GetByID(id string) (userdevice entity.UserDevice, err error) {
	return service.repository.FindByID(id)
}

// GetByUserID find user_deviceby id
func (service Service) GetByUserID(id string) (userdevice entity.UserDevice, err error) {
	return service.repository.FindByUserID(id)
}

// DeleteByID delete user_deviceby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
