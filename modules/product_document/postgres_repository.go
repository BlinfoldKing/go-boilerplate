package productdocument

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

// Save save product_document to db
func (repo PostgresRepository) Save(productDocument entity.ProductDocument) error {
	_, err := repo.db.Table("product_documents").Insert(&productDocument)
	return err
}

// SaveBatch inserts a batch of productDocuments
func (repo PostgresRepository) SaveBatch(productDocuments []entity.ProductDocument) error {
	_, err := repo.db.Table("product_documents").Insert(&productDocuments)
	return err
}

// GetList get list of product_document
func (repo PostgresRepository) GetList(pagination entity.Pagination) (productDocuments []entity.ProductDocument, count int, err error) {
	count, err = repo.db.
		Paginate("product_documents", &productDocuments, pagination)
	return
}

// Update update product_document
func (repo PostgresRepository) Update(id string, changeset entity.ProductDocumentChangeSet) error {
	_, err := repo.db.Table("product_documents").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find product_document by id
func (repo PostgresRepository) FindByID(id string) (productDocument entity.ProductDocument, err error) {
	_, err = repo.db.SQL("SELECT * FROM product_documents WHERE id = ? AND deleted_at IS null", id).Get(&productDocument)
	return
}

// DeleteByID delete product_document by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("product_documents").Where("id = ?", id).Delete(&entity.Product{})
	return err
}

// DeleteByProductID delete product_document by product_id
func (repo PostgresRepository) DeleteByProductID(productID string) error {
	_, err := repo.db.Table("product_documents").Where("product_id = ?", productID).Delete(&entity.Product{})
	return err
}
