package companydocument

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

// Save save companyDocument to db
func (repo PostgresRepository) Save(companyDocument entity.CompanyDocument) error {
	_, err := repo.db.Table("company_documents").Insert(&companyDocument)
	return err
}

// SaveBatch inserts a batch of companyDocuments
func (repo PostgresRepository) SaveBatch(companyDocuments []entity.CompanyDocument) error {
	_, err := repo.db.Table("company_documents").Insert(&companyDocuments)
	return err
}

// GetList get list of companyDocument
func (repo PostgresRepository) GetList(pagination entity.Pagination) (companyDocuments []entity.CompanyDocument, count int, err error) {
	count, err = repo.db.
		Paginate("company_documents", &companyDocuments, pagination)
	return
}

// Update update companyDocument
func (repo PostgresRepository) Update(id string, changeset entity.CompanyDocumentChangeSet) error {
	_, err := repo.db.Table("company_documents").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find company_document by id
func (repo PostgresRepository) FindByID(id string) (companyDocument entity.CompanyDocument, err error) {
	_, err = repo.db.SQL("SELECT * FROM company_documents WHERE id = ? AND deleted_at IS NULL", id).Get(&companyDocument)
	return
}

// DeleteByID delete company_document by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("company_documents").Where("id = ?", id).Delete(&entity.Company{})
	return err
}

// DeleteByCompanyID delete company_document by company_id
func (repo PostgresRepository) DeleteByCompanyID(companyID string) error {
	_, err := repo.db.Table("company_documents").Where("company_id = ?", companyID).Delete(&entity.Company{})
	return err
}
