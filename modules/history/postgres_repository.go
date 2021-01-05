package history

import (
	"go-boilerplate/adapters/postgres"
	"go-boilerplate/entity"
)

// PostgresRepository repository implementation on postgres
type PostgresRepository struct {
	db *postgres.Postgres
}

// CreatePostgresRepository init PostgresRepository
func CreatePostgresRepository(db *postgres.Postgres) Repository {
	return PostgresRepository{db}
}

// Save save history to db
func (repo PostgresRepository) Save(history entity.History) error {
	_, err := repo.db.Table("histories").Insert(&history)
	return err
}

// GetList get list of history
func (repo PostgresRepository) GetList(pagination entity.Pagination) (histories []entity.History, count int, err error) {
	count, err = repo.db.
		Paginate("histories", &histories, pagination)
	return
}

// Update update history
func (repo PostgresRepository) Update(id string, changeset entity.HistoryChangeSet) error {
	_, err := repo.db.Table("histories").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find history by id
func (repo PostgresRepository) FindByID(id string) (history entity.History, err error) {
	_, err = repo.db.SQL("SELECT * FROM histories WHERE id = ? AND deleted_at IS null", id).Get(&history)
	return
}

// DeleteByID delete history by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("histories").Where("id = ?", id).Delete(&entity.History{})
	return err
}
