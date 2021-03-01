package assetwarehouseextend

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

// SaveBatch inserts a batch of assetWarehouses
func (repo PostgresRepository) SaveBatch(assetWarehouses []entity.AssetWarehouse) error {
	_, err := repo.db.Table("asset_warehouses").Insert(&assetWarehouses)
	return err
}

// DeleteByAssetID delete asset warehouse by asset id
func (repo PostgresRepository) DeleteByAssetID(assetID string) error {
	_, err := repo.db.Table("asset_warehouses").Where("asset_id = ?", assetID).Delete(&entity.AssetWarehouse{})
	return err
}
