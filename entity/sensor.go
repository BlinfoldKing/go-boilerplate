package entity

import (
	"github.com/satori/uuid"
	"time"
)

// SensorType sensor type
type SensorType = int

const (
	//SensorTemperature :nodoc
	SensorTemperature SensorType = iota + 1
	// SensorHumidty :nodoc
	SensorHumidty
	// SensorDoor :nodoc
	SensorDoor
	// SensorGate :nodoc
	SensorGate
	// SensorFence :nodoc
	SensorFence
	// SensorGyro :nodoc
	SensorGyro
	// SensorCam :nodoc
	SensorCam
	// SensorRouter :nodoc
	SensorRouter
	// SensorModem :nodoc
	SensorModem
	// SensorSwitch :nodoc
	SensorSwitch
)

// Sensor sensor entity
type Sensor struct {
	ID          string     `json:"id" xorm:"id"`
	Name        string     `json:"name" xorm:"name"`
	Description string     `json:"description" xorm:"description"`
	SiteID      string     `json:"site_id" xorm:"site_id"`
	SensorType  SensorType `json:"sensor_type" xorm:"sensor_type"`
	Code        string     `json:"code" xorm:"code"`
	CreatedAt   time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt   time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt   *time.Time `json:"deleted_at" xorm:"deleted"`
}

// SensorChangeSet change set forsensor
type SensorChangeSet struct {
	ID          string     `json:"id" xorm:"id"`
	Name        string     `json:"name" xorm:"name"`
	Description string     `json:"description" xorm:"description"`
	SiteID      string     `json:"site_id" xorm:"site_id"`
	SensorType  SensorType `json:"sensor_type" xorm:"sensor_type"`
	Code        string     `json:"code" xorm:"code"`
	CreatedAt   time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt   time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt   *time.Time `json:"deleted_at" xorm:"deleted"`
}

// NewSensor create newsensor
func NewSensor(name string, description string, siteID string, sensorType int, code string) (sensor Sensor, err error) {
	sensor = Sensor{
		ID:          uuid.NewV4().String(),
		Name:        name,
		Description: description,
		SiteID:      siteID,
		SensorType:  sensorType,
		Code:        code,
	}
	return
}
