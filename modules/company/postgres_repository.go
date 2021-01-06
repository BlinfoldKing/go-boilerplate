package company

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

// Save save company to db
func (repo PostgresRepository) Save(company entity.Company) error {
	_, err := repo.db.Table("companies").Insert(&company)
	return err
}

// GetList get list of company
func (repo PostgresRepository) GetList(pagination entity.Pagination) (companies []entity.Company, count int, err error) {
	count, err = repo.db.
		Paginate("companies", &companies, pagination)
	return
}

// FindByBrandID gets list of company that associated with specified brand
func (repo PostgresRepository) FindByBrandID(brandID string) (companies []entity.Company, err error) {
	err = repo.db.
		SQL(`SELECT 
				c.*
			FROM 
				brand_companies bc
			INNER JOIN companies c
				ON bc.brand_id = ?
				AND bc.company_id = c.id
				AND deleted_at IS NULL`,
			brandID).Find(&companies)
	return
}

// Update update company
func (repo PostgresRepository) Update(id string, changeset entity.CompanyChangeSet) error {
	_, err := repo.db.Table("companies").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find company by id
func (repo PostgresRepository) FindByID(id string) (company entity.Company, err error) {
	_, err = repo.db.SQL("SELECT * FROM companies WHERE id = ? AND deleted_at IS NULL", id).Get(&company)
	return
}

// DeleteByID delete company by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("companies").Where("id = ?", id).Delete(&entity.Company{})
	return err
}
