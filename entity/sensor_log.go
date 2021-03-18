package entity

import (
	"github.com/satori/uuid"
	"time"
)

// SensorLog sensorLog entity
type SensorLog struct {
	ID        string     `json:"id" xorm:"id"`
	SensorID  string     `json:"sensor_id" xorm:"sensor_id"`
	Unit      string     `json:"unit" xorm:"unit"`
	Payload   string     `json:"payload" xorm:"payload"`
	Value     string     `json:"value" xorm:"value"`
	CreatedAt time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt *time.Time `json:"deleted_at" xorm:"deleted"`
}

// SensorLogChangeSet change set forsensorLog
type SensorLogChangeSet struct {
	ID        string     `json:"id" xorm:"id"`
	SensorID  string     `json:"sensor_id" xorm:"sensor_id"`
	Unit      string     `json:"unit" xorm:"unit"`
	Payload   string     `json:"payload" xorm:"payload"`
	Value     string     `json:"value" xorm:"value"`
	CreatedAt time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt *time.Time `json:"deleted_at" xorm:"deleted"`
}

// NewSensorLog create newsensorLog
func NewSensorLog(sensorID string, unit string, payload string, value string) (sensorLog SensorLog, err error) {
	sensorLog = SensorLog{
		ID:       uuid.NewV4().String(),
		SensorID: sensorID,
		Unit:     unit,
		Payload:  payload,
		Value:    value,
	}
	return
}
