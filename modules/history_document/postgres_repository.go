package historydocument

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

// Save save history_document to db
func (repo PostgresRepository) Save(historyDocument entity.HistoryDocument) error {
	_, err := repo.db.Table("history_documents").Insert(&historyDocument)
	return err
}

// SaveBatch inserts a batch of historyDocuments
func (repo PostgresRepository) SaveBatch(historyDocuments []entity.HistoryDocument) error {
	_, err := repo.db.Table("history_documents").Insert(&historyDocuments)
	return err
}

// GetList get list of history_document
func (repo PostgresRepository) GetList(pagination entity.Pagination) (historyDocuments []entity.HistoryDocument, count int, err error) {
	count, err = repo.db.
		Paginate("history_documents", &historyDocuments, pagination)
	return
}

// Update update history_document
func (repo PostgresRepository) Update(id string, changeset entity.HistoryDocumentChangeSet) error {
	_, err := repo.db.Table("history_documents").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find history_document by id
func (repo PostgresRepository) FindByID(id string) (historyDocument entity.HistoryDocument, err error) {
	_, err = repo.db.SQL("SELECT * FROM history_documents WHERE id = ? AND deleted_at IS NULL", id).Get(&historyDocument)
	return
}

// DeleteByID delete history_document by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("history_documents").Where("id = ?", id).Delete(&entity.HistoryDocument{})
	return err
}

// DeleteByHistoryID delete history_document by history_id
func (repo PostgresRepository) DeleteByHistoryID(historyID string) error {
	_, err := repo.db.Table("history_documents").Where("history_id = ?", historyID).Delete(&entity.HistoryDocument{})
	return err
}
