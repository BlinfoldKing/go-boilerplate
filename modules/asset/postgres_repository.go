package asset

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

// Save save asset to db
func (repo PostgresRepository) Save(asset entity.Asset) error {
	_, err := repo.db.Table("assets").Insert(&asset)
	return err
}

// GetList get list of asset
func (repo PostgresRepository) GetList(pagination entity.Pagination) (assets []entity.Asset, count int, err error) {
	count, err = repo.db.
		Paginate("assets", &assets, pagination)
	return
}

// Update update asset
func (repo PostgresRepository) Update(id string, changeset entity.AssetChangeSet) error {
	_, err := repo.db.Table("assets").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find asset by id
func (repo PostgresRepository) FindByID(id string) (asset entity.Asset, err error) {
	_, err = repo.db.SQL("SELECT * FROM assets WHERE id = ?", id).Get(&asset)
	return
}

// FindByWorkOrderID find assets by work order id
func (repo PostgresRepository) FindByWorkOrderID(workOrderID string) (assets []entity.Asset, err error) {
	err = repo.db.
		SQL(`SELECT 
				a.*
			FROM 
				work_order_assets wa
			INNER JOIN assets a
				ON wa.work_order_id = ?
				AND wa.asset_id = a.id`,
			workOrderID).Find(&assets)
	return
}

// DeleteByID delete asset by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Exec("DELETE FROM assets WHERE id = ?", id)
	return err
}
