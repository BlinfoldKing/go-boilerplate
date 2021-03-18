package sensorlog

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

// CreateSensorLog create new sensorLog
func (service Service) CreateSensorLog(sensorID string, unit string, payload string, value string) (sensorLog entity.SensorLog, err error) {
	sensorLog, err = entity.NewSensorLog(sensorID, unit, payload, value)
	if err != nil {
		return
	}
	err = service.repository.Save(sensorLog)
	return
}

// GetList get list of sensorLog
func (service Service) GetList(pagination entity.Pagination) (sensorLog []entity.SensorLog, count int, err error) {
	sensorLog, count, err = service.repository.GetList(pagination)
	return
}

// Update update sensorLog
func (service Service) Update(id string, changeset entity.SensorLogChangeSet) (sensorLog entity.SensorLog, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.SensorLog{}, err
	}
	return service.GetByID(id)
}

// GetByID find sensorLogby id
func (service Service) GetByID(id string) (sensorLog entity.SensorLog, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete sensorLogby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
