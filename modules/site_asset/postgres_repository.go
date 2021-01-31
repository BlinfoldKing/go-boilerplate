package siteasset

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

// Save save siteAsset to db
func (repo PostgresRepository) Save(siteAsset entity.SiteAsset) error {
	_, err := repo.db.Table("site_assets").Insert(&siteAsset)
	return err
}

// SaveBatch inserts a batch of siteAsset
func (repo PostgresRepository) SaveBatch(siteAssets []entity.SiteAsset) error {
	_, err := repo.db.Table("site_assets").Insert(&siteAssets)
	return err
}

// GetList get list of siteAsset
func (repo PostgresRepository) GetList(pagination entity.Pagination) (siteAssets []entity.SiteAsset, count int, err error) {
	count, err = repo.db.
		Paginate("site_assets", &siteAssets, pagination)
	return
}

// Update update siteAsset
func (repo PostgresRepository) Update(id string, changeset entity.SiteAssetChangeSet) error {
	_, err := repo.db.Table("site_assets").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find siteAsset by id
func (repo PostgresRepository) FindByID(id string) (siteAsset entity.SiteAsset, err error) {
	_, err = repo.db.SQL("SELECT * FROM site_assets WHERE id = ? AND deleted_at IS null", id).Get(&siteAsset)
	return
}

// DeleteByID delete siteAsset by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("site_assets").Where("id = ?", id).Delete(&entity.SiteAsset{})
	return err
}
