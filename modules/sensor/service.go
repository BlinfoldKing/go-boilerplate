package sensor

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

// CreateSensor create new sensor
func (service Service) CreateSensor(code string, name string, description string, siteID string, sensorType int) (sensor entity.Sensor, err error) {
	sensor, err = entity.NewSensor(name, description, siteID, sensorType, code)
	if err != nil {
		return
	}
	err = service.repository.Save(sensor)
	return
}

// GetList get list of sensor
func (service Service) GetList(pagination entity.Pagination) (sensor []entity.Sensor, count int, err error) {
	sensor, count, err = service.repository.GetList(pagination)
	return
}

// Update update sensor
func (service Service) Update(id string, changeset entity.SensorChangeSet) (sensor entity.Sensor, err error) {
	err = service.repository.Update(id, changeset)
	if err != nil {
		return entity.Sensor{}, err
	}
	return service.GetByID(id)
}

// GetByID find sensorby id
func (service Service) GetByID(id string) (sensor entity.Sensor, err error) {
	return service.repository.FindByID(id)
}

// DeleteByID delete sensorby id
func (service Service) DeleteByID(id string) (err error) {
	return service.repository.DeleteByID(id)
}
