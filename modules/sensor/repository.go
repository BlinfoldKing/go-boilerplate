package sensor

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.Sensor) error
	DeleteByID(id string) error
	FindByID(id string) (entity.Sensor, error)
	Update(id string, changeset entity.SensorChangeSet) error
	GetList(pagination entity.Pagination) (Sensors []entity.Sensor, count int, err error)
}
