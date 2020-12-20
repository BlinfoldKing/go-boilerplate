package userroles

import (
	"go-boilerplate/adapters/postgres"
	"go-boilerplate/entity"
)

// PostgresRepository repository implementation on postgres
type PostgresRepository struct {
	db *postgres.Postgres
}

// CreatePosgresRepository init PostgresRepository
func CreatePosgresRepository(db *postgres.Postgres) Repository {
	return PostgresRepository{db}
}

// Save save role to db
func (repo PostgresRepository) Save(role entity.UserRole) error {
	_, err := repo.db.Table("user_roles").Insert(&role)
	return err
}

// GetAllByUserID find role by slug
func (repo PostgresRepository) GetAllByUserID(id string) (role []entity.UserRole, err error) {
	err = repo.db.SQL("SELECT * FROM user_roles WHERE user_id = ?", id).Find(&role)
	return
}

// FindByID find user by id
func (repo PostgresRepository) FindByID(id string) (role entity.UserRole, err error) {
	_, err = repo.db.SQL("SELECT * FROM user_roles WHERE id = ?", id).Get(&role)
	return
}

// DeleteByID delete user by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Exec("DELETE FROM user_roles WHERE id = ?", id)
	return err
}
