package companycontactget

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

// GetList get list of company_contact
func (repo PostgresRepository) GetList(pagination entity.Pagination) (companyContacts []entity.CompanyContact, count int, err error) {
	count, err = repo.db.
		Paginate("company_contacts", &companyContacts, pagination)
	return
}

// FindByID find company_contact by id
func (repo PostgresRepository) FindByID(id string) (companyContact entity.CompanyContact, err error) {
	_, err = repo.db.SQL("SELECT * FROM company_contacts WHERE id = ? AND deleted_at IS null", id).Get(&companyContact)
	return
}
