package involved_user

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

// Save save involved_user to db
func (repo PostgresRepository) Save(involved_user entity.InvolvedUser) error {
	_, err := repo.db.Table("involved_users").Insert(&involved_user)
	return err
}

// GetList get list of involved_user
func (repo PostgresRepository) GetList(pagination entity.Pagination) (involved_users []entity.InvolvedUser, count int, err error) {
	count, err = repo.db.
		Paginate("involved_users", &involved_users, pagination)
	return
}

// Update update involved_user
func (repo PostgresRepository) Update(id string, changeset entity.InvolvedUserChangeSet) error {
	_, err := repo.db.Table("involved_users").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find involved_user by id
func (repo PostgresRepository) FindByID(id string) (involved_user entity.InvolvedUser, err error) {
	_, err = repo.db.SQL("SELECT * FROM involved_users WHERE id = ? AND deleted_at = null", id).Get(&involved_user)
	return
}

// DeleteByID delete involved_user by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("involved_users").Where("id = ?", id).Delete(&entity.InvolvedUser{})
	return err
}
