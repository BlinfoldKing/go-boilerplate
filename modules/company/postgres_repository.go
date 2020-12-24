package company

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

// Save save company to db
func (repo PostgresRepository) Save(company entity.Company) error {
	_, err := repo.db.Table("companies").Insert(&company)
	return err
}

// GetList get list of company
func (repo PostgresRepository) GetList(pagination entity.Pagination) (companys []entity.Company, count int, err error) {
	count, err = repo.db.
		Paginate("companies", &companys, pagination)
	return
}

// Update update company
func (repo PostgresRepository) Update(id string, changeset entity.CompanyChangeSet) error {
	_, err := repo.db.Table("companies").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find company by id
func (repo PostgresRepository) FindByID(id string) (company entity.Company, err error) {
	_, err = repo.db.SQL("SELECT * FROM companies WHERE id = ?", id).Get(&company)
	return
}

// DeleteByID delete company by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Exec("DELETE FROM companies WHERE id = ?", id)
	return err
}
