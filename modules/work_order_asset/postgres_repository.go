package work_order_asset

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

// Save save work_order_asset to db
func (repo PostgresRepository) Save(work_order_asset entity.WorkOrderAsset) error {
	_, err := repo.db.Table("work_order_assets").Insert(&work_order_asset)
	return err
}

// GetList get list of work_order_asset
func (repo PostgresRepository) GetList(pagination entity.Pagination) (work_order_assets []entity.WorkOrderAsset, count int, err error) {
	count, err = repo.db.
		Paginate("work_order_assets", &work_order_assets, pagination)
	return
}

// Update update work_order_asset
func (repo PostgresRepository) Update(id string, changeset entity.WorkOrderAssetChangeSet) error {
	_, err := repo.db.Table("work_order_assets").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find work_order_asset by id
func (repo PostgresRepository) FindByID(id string) (work_order_asset entity.WorkOrderAsset, err error) {
	_, err = repo.db.SQL("SELECT * FROM work_order_assets WHERE id = ? AND deleted_at = null", id).Get(&work_order_asset)
	return
}

// DeleteByID delete work_order_asset by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("work_order_assets").Where("id = ?", id).Delete(&entity.WorkOrderAsset{})
	return err
}
