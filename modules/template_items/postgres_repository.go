package templateitems

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

// Save save templateItems to db
func (repo PostgresRepository) Save(templateItems entity.TemplateItems) error {
	_, err := repo.db.Table("template_items").Insert(&templateItems)
	return err
}

// GetList get list of templateItems
func (repo PostgresRepository) GetList(pagination entity.Pagination) (templateItemss []entity.TemplateItems, count int, err error) {
	count, err = repo.db.
		Paginate("template_items", &templateItemss, pagination)
	return
}

// Update update templateItems
func (repo PostgresRepository) Update(id string, changeset entity.TemplateItemsChangeSet) error {
	_, err := repo.db.Table("template_items").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find templateItems by id
func (repo PostgresRepository) FindByID(id string) (templateItems entity.TemplateItems, err error) {
	_, err = repo.db.SQL("SELECT * FROM template_items WHERE id = ? AND deleted_at IS null", id).Get(&templateItems)
	return
}

// FindByTemplateID find templateItems by id
func (repo PostgresRepository) FindByTemplateID(templateID string) (templateItems []entity.TemplateItems, err error) {
	err = repo.db.SQL("SELECT * FROM template_items WHERE template_id = ? AND deleted_at IS null", templateID).Find(&templateItems)
	return
}

// DeleteByID delete templateItems by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("template_items").Where("id = ?", id).Delete(&entity.TemplateItems{})
	return err
}
