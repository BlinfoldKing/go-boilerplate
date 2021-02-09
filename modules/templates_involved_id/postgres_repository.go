package templatesinvolvedid

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

// Save save templatesInvolvedID to db
func (repo PostgresRepository) Save(templatesInvolvedID entity.TemplatesInvolvedID) error {
	_, err := repo.db.Table("templates_involved_ids").Insert(&templatesInvolvedID)
	return err
}

// GetList get list of templatesInvolvedID
func (repo PostgresRepository) GetList(pagination entity.Pagination) (templatesInvolvedIDs []entity.TemplatesInvolvedID, count int, err error) {
	count, err = repo.db.
		Paginate("templates_involved_ids", &templatesInvolvedIDs, pagination)
	return
}

// Update update templatesInvolvedID
func (repo PostgresRepository) Update(id string, changeset entity.TemplatesInvolvedIDChangeSet) error {
	_, err := repo.db.Table("templates_involved_ids").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find templatesInvolvedID by id
func (repo PostgresRepository) FindByID(id string) (templatesInvolvedID entity.TemplatesInvolvedID, err error) {
	_, err = repo.db.SQL("SELECT * FROM templates_involved_ids WHERE id = ? AND deleted_at IS null", id).Get(&templatesInvolvedID)
	return
}

// DeleteByID delete templatesInvolvedID by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("templates_involved_ids").Where("id = ?", id).Delete(&entity.TemplatesInvolvedID{})
	return err
}
