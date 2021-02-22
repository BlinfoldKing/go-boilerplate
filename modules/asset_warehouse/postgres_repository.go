package assetwarehouse

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

// Save save assetWarehouse to db
func (repo PostgresRepository) Save(assetWarehouse entity.AssetWarehouse) error {
	_, err := repo.db.Table("asset_warehouses").Insert(&assetWarehouse)
	return err
}

// GetList get list of assetWarehouse
func (repo PostgresRepository) GetList(pagination entity.Pagination) (assetwarehouses []entity.AssetWarehouse, count int, err error) {
	count, err = repo.db.
		Paginate("asset_warehouses", &assetwarehouses, pagination)
	return
}

// Update update assetWarehouse
func (repo PostgresRepository) Update(id string, changeset entity.AssetWarehouseChangeSet) error {
	_, err := repo.db.Table("asset_warehouses").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find assetWarehouse by id
func (repo PostgresRepository) FindByID(id string) (assetWarehouse entity.AssetWarehouse, err error) {
	_, err = repo.db.SQL("SELECT * FROM asset_warehouses WHERE id = ? AND deleted_at IS null", id).Get(&assetWarehouse)
	return
}

// DeleteByID delete assetWarehouse by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("asset_warehouses").Where("id = ?", id).Delete(&entity.AssetWarehouse{})
	return err
}
