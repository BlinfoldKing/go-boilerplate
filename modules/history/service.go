package history

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

// CreateHistory create new history
func (service Service) CreateHistory(userID, assetID, action, description string, cost float64) (history entity.History, err error) {
	history, err = entity.NewHistory(
		userID,
		assetID,
		action,
		description,
		cost,
	)
	if err != nil {
		return
	}
	err = service.repository.Save(history)
	return
}

// GetList get list of history
func (service Service) GetList(pagination entity.Pagination) (history []entity.History, count int, err error) {
	history, count, err = service.repository.GetList(pagination)
	return
}

// Update update history
func (service Service) Update(id string, changeset entity.HistoryChangeSet) (history entity.History, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.History{}, err
	}
	return service.GetByID(id)
}

// GetByID find historyby id
func (service Service) GetByID(id string) (history entity.History, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete historyby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
