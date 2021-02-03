package userdevice

// CreateRequest request for create new user_device
type CreateRequest struct {
	UserID      string `json:"user_id" validate:"required"`
	DeviceToken string `json:"device_token" validate:"required"`
}

// UpdateRequest request for update user_device
type UpdateRequest struct {
	UserID      string `json:"user_id" validate:"required"`
	DeviceToken string `json:"device_token" validate:"required"`
}
