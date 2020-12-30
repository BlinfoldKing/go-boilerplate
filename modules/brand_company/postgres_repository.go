package brandcompany

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

// Save save brand_company to db
func (repo PostgresRepository) Save(brandCompany entity.BrandCompany) error {
	_, err := repo.db.Table("brand_companies").Insert(&brandCompany)
	return err
}

// SaveBatch inserts a batch of brand companies
func (repo PostgresRepository) SaveBatch(brandCompanies []entity.BrandCompany) error {
	_, err := repo.db.Table("brand_companies").Insert(&brandCompanies)
	return err
}

// GetList get list of brand_company
func (repo PostgresRepository) GetList(pagination entity.Pagination) (brandCompanies []entity.BrandCompany, count int, err error) {
	count, err = repo.db.
		Paginate("brand_companies", &brandCompanies, pagination)
	return
}

// Update update brand_company
func (repo PostgresRepository) Update(id string, changeset entity.BrandCompanyChangeSet) error {
	_, err := repo.db.Table("brand_companies").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find brand_company by id
func (repo PostgresRepository) FindByID(id string) (brandCompany entity.BrandCompany, err error) {
	_, err = repo.db.SQL("SELECT * FROM brand_companies WHERE id = ? AND deleted_at = null", id).Get(&brandCompany)
	return
}

// DeleteByID delete brand_company by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("brand_companies").Where("id = ?", id).Delete(&entity.BrandCompany{})
	return err
}

// DeleteByBrandID deletes all brandcompany with brand ID
func (repo PostgresRepository) DeleteByBrandID(brandID string) error {
	_, err := repo.db.Table("brand_companies").Where("brand_id = ?", brandID).Delete(&entity.BrandCompany{})
	return err
}
