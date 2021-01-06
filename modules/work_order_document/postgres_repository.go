package work_order_document

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

// Save save work_order_document to db
func (repo PostgresRepository) Save(work_order_document entity.WorkOrderDocument) error {
	_, err := repo.db.Table("work_order_documents").Insert(&work_order_document)
	return err
}

// GetList get list of work_order_document
func (repo PostgresRepository) GetList(pagination entity.Pagination) (work_order_documents []entity.WorkOrderDocument, count int, err error) {
	count, err = repo.db.
		Paginate("work_order_documents", &work_order_documents, pagination)
	return
}

// Update update work_order_document
func (repo PostgresRepository) Update(id string, changeset entity.WorkOrderDocumentChangeSet) error {
	_, err := repo.db.Table("work_order_documents").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find work_order_document by id
func (repo PostgresRepository) FindByID(id string) (work_order_document entity.WorkOrderDocument, err error) {
	_, err = repo.db.SQL("SELECT * FROM work_order_documents WHERE id = ? AND deleted_at = null", id).Get(&work_order_document)
	return
}

// DeleteByID delete work_order_document by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("work_order_documents").Where("id = ?", id).Delete(&entity.WorkOrderDocument{})
	return err
}
