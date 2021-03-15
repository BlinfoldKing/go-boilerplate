package sensor

// CreateRequest request for create new sensor
type CreateRequest struct {
	SensorType  int    `json:"sensor_type"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SiteID      string `json:"site_id"`
}

// UpdateRequest request for update sensor
type UpdateRequest struct {
	SensorType  int    `json:"sensor_type"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SiteID      string `json:"site_id"`
}
