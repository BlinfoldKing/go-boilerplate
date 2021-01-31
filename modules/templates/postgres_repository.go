package templates

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

// Save save templates to db
func (repo PostgresRepository) Save(templates entity.Templates) error {
	_, err := repo.db.Table("templates").Insert(&templates)
	return err
}

// GetList get list of templates
func (repo PostgresRepository) GetList(pagination entity.Pagination) (templatess []entity.Templates, count int, err error) {
	count, err = repo.db.
		Paginate("templates", &templatess, pagination)
	return
}

// Update update templates
func (repo PostgresRepository) Update(id string, changeset entity.TemplatesChangeSet) error {
	_, err := repo.db.Table("templates").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find templates by id
func (repo PostgresRepository) FindByID(id string) (templates entity.Templates, err error) {
	_, err = repo.db.SQL("SELECT * FROM templates WHERE id = ? AND deleted_at IS null", id).Get(&templates)
	return
}

// DeleteByID delete templates by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("templates").Where("id = ?", id).Delete(&entity.Templates{})
	return err
}
