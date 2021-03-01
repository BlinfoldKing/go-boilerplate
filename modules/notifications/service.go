package notifications

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

// CreateNotification create new notification
func (service Service) CreateNotification(userID, title, subtitle, urlLink, body string, notifType, status int) (notification entity.Notification, err error) {
	notification = entity.NewNotification(userID, title, subtitle, urlLink, body, notifType, status)
	err = service.repository.Save(notification)
	return
}

// GetList get list of notification
func (service Service) GetList(pagination entity.Pagination) (notification []entity.Notification, count int, err error) {
	notification, count, err = service.repository.GetList(pagination)
	return
}

// Update update notification
func (service Service) Update(id string, changeset entity.NotificationChangeSet) (notification entity.Notification, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.Notification{}, err
	}
	return service.GetByID(id)
}

// GetByID find notificationby id
func (service Service) GetByID(id string) (notification entity.Notification, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete notificationby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
