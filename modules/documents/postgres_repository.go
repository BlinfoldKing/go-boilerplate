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

// GetList get list of document
func (repo PostgresRepository) GetList(pagination entity.Pagination) (documents []entity.Document, count int, err error) {
	count, err = repo.db.
		Paginate("documents", &documents, pagination)
	return
}

// Update update document
func (repo PostgresRepository) Update(id string, changeset entity.DocumentChangeSet) error {
	_, err := repo.db.Table("documents").Where("id = ?", id).Update(&changeset)
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
				AND pd.document_id = d.id
				AND pd.deleted_at IS NULL`,
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
				AND hd.document_id = d.id
				AND hd.deleted_at IS NULL`,
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
				AND cd.document_id = d.id
				AND cd.deleted_at IS NULL`,
			companyID).Find(&documents)
	return
}

// FindByWorkOrderID find assets by work order id
func (repo PostgresRepository) FindByWorkOrderID(workOrderID string) (documents []entity.Document, err error) {
	err = repo.db.
		SQL(`SELECT 
				d.*
			FROM 
				work_order_documents wd
			INNER JOIN documents d
				ON wd.work_order_id = ?
				AND wd.document_id = d.id
				AND wd.deleted_at IS NULL`,
			workOrderID).Find(&documents)
	return
}

// FindBySiteID find documents by site id
func (repo PostgresRepository) FindBySiteID(siteID string) (documents []entity.Document, err error) {
	err = repo.db.
		SQL(`SELECT 
				d.*
			FROM 
				site_documents pd
			INNER JOIN documents d
				ON pd.site_id = ?
				AND pd.document_id = d.id
				AND pd.deleted_at IS NULL`,
			siteID).Find(&documents)
	return
}

// FindByObjectBucketName finds document by objectName and bucketName
func (repo PostgresRepository) FindByObjectBucketName(objectName, bucketName string) (document entity.Document, err error) {
	_, err = repo.db.SQL("SELECT * FROM documents WHERE object_name = ? AND bucket_name = ?", objectName, bucketName).Get(&document)
	return
}

// DeleteByID delete document by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("documents").Where("id = ?", id).Delete(&entity.Document{})
	return err
}
