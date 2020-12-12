package users

import (
	"go-boilerplate/entity"

	"xorm.io/xorm"
	// "gorm.io/gorm"
)

// PostgresRepository repository implementation on postgres
type PostgresRepository struct {
	db *xorm.Engine
}

// CreatePosgresRepository init PostgresRepository
func CreatePosgresRepository(db *xorm.Engine) Repository {
	return PostgresRepository{db}
}

// Save save user to db
func (repo PostgresRepository) Save(user entity.User) error {
	_, err := repo.db.Table("users").Insert(&user)
	return err
}

// FindByEmail find user by email
func (repo PostgresRepository) FindByEmail(email string) (user entity.User, err error) {
	_, err = repo.db.SQL("SELECT * FROM users WHERE email = ?", email).Get(&user)
	return
}
