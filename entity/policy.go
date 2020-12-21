package entity

// Policy Auth policy
type Policy struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Role   Role   `json:"role"`
}
