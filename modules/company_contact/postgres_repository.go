package companycontact

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

// Save save company_contact to db
func (repo PostgresRepository) Save(companyContact entity.CompanyContact) error {
	_, err := repo.db.Table("company_contacts").Insert(&companyContact)
	return err
}

// SaveBatch inserts a batch of companyContacts
func (repo PostgresRepository) SaveBatch(companyContacts []entity.CompanyContact) error {
	_, err := repo.db.Table("company_contacts").Insert(&companyContacts)
	return err
}

// GetList get list of company_contact
func (repo PostgresRepository) GetList(pagination entity.Pagination) (companyContacts []entity.CompanyContact, count int, err error) {
	count, err = repo.db.
		Paginate("company_contacts", &companyContacts, pagination)
	return
}

// Update update company_contact
func (repo PostgresRepository) Update(id string, changeset entity.CompanyContactChangeSet) error {
	_, err := repo.db.Table("company_contacts").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find company_contact by id
func (repo PostgresRepository) FindByID(id string) (companyContact entity.CompanyContact, err error) {
	_, err = repo.db.SQL("SELECT * FROM company_contacts WHERE id = ? AND deleted_at IS null", id).Get(&companyContact)
	return
}

// DeleteByID delete company_contact by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("company_contacts").Where("id = ?", id).Delete(&entity.CompanyContact{})
	return err
}

// DeleteByCompanyID delete company contact by company id
func (repo PostgresRepository) DeleteByCompanyID(companyID string) error {
	_, err := repo.db.Table("company_contacts").Where("company_id = ?", companyID).Delete(&entity.CompanyContact{})
	return err
}
