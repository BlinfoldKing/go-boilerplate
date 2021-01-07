package workorderdocument

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

// Save save work_order_document to db
func (repo PostgresRepository) Save(workOrderDocument entity.WorkOrderDocument) error {
	_, err := repo.db.Table("work_order_documents").Insert(&workOrderDocument)
	return err
}

// SaveBatch inserts a batch of workorderDocuments
func (repo PostgresRepository) SaveBatch(workorderDocuments []entity.WorkOrderDocument) error {
	_, err := repo.db.Table("work_order_documents").Insert(&workorderDocuments)
	return err
}

// GetList get list of work_order_document
func (repo PostgresRepository) GetList(pagination entity.Pagination) (workOrderDocuments []entity.WorkOrderDocument, count int, err error) {
	count, err = repo.db.
		Paginate("work_order_documents", &workOrderDocuments, pagination)
	return
}

// Update update work_order_document
func (repo PostgresRepository) Update(id string, changeset entity.WorkOrderDocumentChangeSet) error {
	_, err := repo.db.Table("work_order_documents").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find work_order_document by id
func (repo PostgresRepository) FindByID(id string) (workOrderDocument entity.WorkOrderDocument, err error) {
	_, err = repo.db.SQL("SELECT * FROM work_order_documents WHERE id = ? AND deleted_at IS null", id).Get(&workOrderDocument)
	return
}

// DeleteByID delete work_order_document by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("work_order_documents").Where("id = ?", id).Delete(&entity.WorkOrderDocument{})
	return err
}

// DeleteByWorkOrderID delete workorder_document by workorder_id
func (repo PostgresRepository) DeleteByWorkOrderID(workorderID string) error {
	_, err := repo.db.Table("work_order_documents").Where("work_order_id = ?", workorderID).Delete(&entity.WorkOrder{})
	return err
}
