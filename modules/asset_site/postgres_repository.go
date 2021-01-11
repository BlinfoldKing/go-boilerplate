package assetsite

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

// Save save assetSite to db
func (repo PostgresRepository) Save(assetSite entity.AssetSite) error {
	_, err := repo.db.Table("asset_site").Insert(&assetSite)
	return err
}

// GetList get list of assetSite
func (repo PostgresRepository) GetList(pagination entity.Pagination) (assetSites []entity.AssetSite, count int, err error) {
	count, err = repo.db.
		Paginate("asset_site", &assetSites, pagination)
	return
}

// Update update assetSite
func (repo PostgresRepository) Update(id string, changeset entity.AssetSiteChangeSet) error {
	_, err := repo.db.Table("asset_site").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find assetSite by id
func (repo PostgresRepository) FindByID(id string) (assetSite entity.AssetSite, err error) {
	_, err = repo.db.SQL("SELECT * FROM asset_site WHERE id = ? AND deleted_at IS null", id).Get(&assetSite)
	return
}

// DeleteByID delete assetSite by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("asset_site").Where("id = ?", id).Delete(&entity.AssetSite{})
	return err
}
