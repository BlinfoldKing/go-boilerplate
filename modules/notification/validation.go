package notification

// CreateRequest request for create new notification
type CreateRequest struct {
	UserID   string `json:"user_id" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Subtitle string `json:"subtitle" validate:"required"`
	URLLink  string `json:"url_link" validate:"required"`
	Body     string `json:"body" validate:"required"`
}

// UpdateRequest request for update notification
type UpdateRequest struct {
	Title string `json:"title" validate:"required"`
}
