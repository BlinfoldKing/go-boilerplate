package brand_company

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

// Save save brand_company to db
func (repo PostgresRepository) Save(brand_company entity.BrandCompany) error {
	_, err := repo.db.Table("brand_companys").Insert(&brand_company)
	return err
}

// GetList get list of brand_company
func (repo PostgresRepository) GetList(pagination entity.Pagination) (brand_companys []entity.BrandCompany, count int, err error) {
	count, err = repo.db.
		Paginate("brand_companys", &brand_companys, pagination)
	return
}

// Update update brand_company
func (repo PostgresRepository) Update(id string, changeset entity.BrandCompanyChangeSet) error {
	_, err := repo.db.Table("brand_companys").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find brand_company by id
func (repo PostgresRepository) FindByID(id string) (brand_company entity.BrandCompany, err error) {
	_, err = repo.db.SQL("SELECT * FROM brand_companys WHERE id = ?", id).Get(&brand_company)
	return
}

// DeleteByID delete brand_company by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Exec("DELETE FROM brand_companys WHERE id = ?", id)
	return err
}
