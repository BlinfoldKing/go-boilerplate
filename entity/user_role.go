package entity

import "github.com/satori/uuid"

// UserRole role assigned to a user
type UserRole struct {
	ID     string `xorm:"id" json:"id"`
	UserID string `xorm:"user_id" json:"user_id"`
	RoleID string `xorm:"role_id" json:"role_id"`
}

// NewUserRole create new user role
func NewUserRole(userID, roleID string) (UserRole, error) {
	return UserRole{
		uuid.NewV4().String(),
		userID,
		roleID,
	}, nil
}
