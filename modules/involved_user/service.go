package involveduser

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
)

func InitInvolvedUserService(adapters adapters.Adapters) Service {
	repository := CreatePostgresRepository(adapters.Postgres)
	return CreateService(repository)
}

// Service contains business logic
type Service struct {
	repository Repository
}

// CreateService init service
func CreateService(repo Repository) Service {
	return Service{repo}
}

// CreateInvolvedUser create new involved_user
func (service Service) CreateInvolvedUser(userID, workOrderID string) (involvedUser entity.InvolvedUser, err error) {
	involvedUser, err = entity.NewInvolvedUser(userID, workOrderID)
	if err != nil {
		return
	}
	err = service.repository.Save(involvedUser)
	return
}

// CreateBatchInvolvedUsers creates a batch of new InvolvedUsers
func (service Service) CreateBatchInvolvedUsers(workOrderID string, userIDs []string) (involvedUsers []entity.InvolvedUser, err error) {
	for _, userID := range userIDs {
		involvedUser, err := entity.NewInvolvedUser(userID, workOrderID)
		if err != nil {
			return []entity.InvolvedUser{}, err
		}
		involvedUsers = append(involvedUsers, involvedUser)
	}
	err = service.repository.SaveBatch(involvedUsers)
	return
}

// GetList get list of involved_user
func (service Service) GetList(pagination entity.Pagination) (involvedUser []entity.InvolvedUser, count int, err error) {
	involvedUser, count, err = service.repository.GetList(pagination)
	return
}

// Update update involved_user
func (service Service) Update(id string, changeset entity.InvolvedUserChangeSet) (involvedUser entity.InvolvedUser, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.InvolvedUser{}, err
	}
	return service.GetByID(id)
}

// GetByID find involved_userby id
func (service Service) GetByID(id string) (involvedUser entity.InvolvedUser, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete involved_userby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}

// DeleteByWorkOrderID delete involved_user by WorkOrder id
func (service Service) DeleteByWorkOrderID(workOrderID string) (err error) {
	return service.repository.DeleteByWorkOrderID(workOrderID)
}
