package sensorlog

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.SensorLog) error
	DeleteByID(id string) error
	FindByID(id string) (entity.SensorLog, error)
	Update(id string, changeset entity.SensorLogChangeSet) error
	GetList(pagination entity.Pagination) (SensorLogs []entity.SensorLog, count int, err error)
}
