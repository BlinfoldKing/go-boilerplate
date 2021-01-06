package documents

import (
	"go-boilerplate/adapters/postgres"
	"go-boilerplate/entity"
)

// PostgresRepository repository implementation on postgres
type PostgresRepository struct {
	db *postgres.Postgres
}

// CreatePostgresRepository init PostgresRepository
func CreatePostgresRepository(db *postgres.Postgres) StorageRepository {
	return PostgresRepository{db}
}

// Save document to db
func (repo PostgresRepository) Save(document entity.Document) error {
	_, err := repo.db.Table("documents").Insert(&document)
	return err
}

// FindByID find document by id
func (repo PostgresRepository) FindByID(id string) (document entity.Document, err error) {
	_, err = repo.db.SQL("SELECT * FROM documents WHERE id = ? AND deleted_at IS NULL", id).Get(&document)
	return
}

// FindByProductID find documents by product id
func (repo PostgresRepository) FindByProductID(productID string) (documents []entity.Document, err error) {
	err = repo.db.
		SQL(`SELECT 
				d.*
			FROM 
				product_documents pd
			INNER JOIN documents d
				ON pd.product_id = ?
				AND pd.document_id = d.id`,
			productID).Find(&documents)
	return
}

// FindByHistoryID find documents by history id
func (repo PostgresRepository) FindByHistoryID(historyID string) (documents []entity.Document, err error) {
	err = repo.db.
		SQL(`SELECT 
				d.*
			FROM 
				history_documents hd
			INNER JOIN documents d
				ON hd.history_id = ?
				AND hd.document_id = d.id`,
			historyID).Find(&documents)
	return
}

// FindByCompanyID find documents by company id
func (repo PostgresRepository) FindByCompanyID(companyID string) (documents []entity.Document, err error) {
	err = repo.db.
		SQL(`SELECT 
				d.*
			FROM 
				company_documents cd
			INNER JOIN documents d
				ON cd.company_id = ?
				AND cd.document_id = d.id`,
			companyID).Find(&documents)
	return
}

// FindByObjectBucketName finds document by objectName and bucketName
func (repo PostgresRepository) FindByObjectBucketName(objectName, bucketName string) (document entity.Document, err error) {
	_, err = repo.db.SQL("SELECT * FROM documents WHERE object_name = ? AND bucket_name = ?", objectName, bucketName).Get(&document)
	return
}
