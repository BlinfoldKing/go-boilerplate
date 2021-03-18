package sensorlog

// CreateRequest request for create new sensorLog
type CreateRequest struct {
	Unit     string `json:"unit"`
	Payload  string `json:"payload"`
	SensorID string `json:"sensor_id"`
	Value    string `json:"value"`
}

// UpdateRequest request for update sensorLog
type UpdateRequest struct {
	Unit     string `json:"unit"`
	Payload  string `json:"payload"`
	SensorID string `json:"sensor_id"`
	Value    string `json:"value"`
}
