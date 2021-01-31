package sitedocument

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

// Save save site_document to db
func (repo PostgresRepository) Save(siteDocument entity.SiteDocument) error {
	_, err := repo.db.Table("site_documents").Insert(&siteDocument)
	return err
}

// GetList get list of site_document
func (repo PostgresRepository) GetList(pagination entity.Pagination) (siteDocument []entity.SiteDocument, count int, err error) {
	count, err = repo.db.
		Paginate("site_documents", &siteDocument, pagination)
	return
}

// Update update site_document
func (repo PostgresRepository) Update(id string, changeset entity.SiteDocumentChangeSet) error {
	_, err := repo.db.Table("site_documents").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find site_document by id
func (repo PostgresRepository) FindByID(id string) (siteDocument entity.SiteDocument, err error) {
	_, err = repo.db.SQL("SELECT * FROM site_documents WHERE id = ? AND deleted_at IS null", id).Get(&siteDocument)
	return
}

// DeleteByID delete site_document by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("site_documents").Where("id = ?", id).Delete(&entity.SiteDocument{})
	return err
}
