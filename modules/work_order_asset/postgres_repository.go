package workorderasset

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

// Save save work_order_asset to db
func (repo PostgresRepository) Save(workOrderAsset entity.WorkOrderAsset) error {
	_, err := repo.db.Table("work_order_assets").Insert(&workOrderAsset)
	return err
}

// SaveBatch inserts a batch of workOrderAssets
func (repo PostgresRepository) SaveBatch(workOrderAssets []entity.WorkOrderAsset) error {
	_, err := repo.db.Table("work_order_assets").Insert(&workOrderAssets)
	return err
}

// GetList get list of work_order_asset
func (repo PostgresRepository) GetList(pagination entity.Pagination) (workOrderAssets []entity.WorkOrderAsset, count int, err error) {
	count, err = repo.db.
		Paginate("work_order_assets", &workOrderAssets, pagination)
	return
}

// GetAllByWorkorderID get list of work_order_asset
func (repo PostgresRepository) GetAllByWorkorderID(id string) (workOrderAssets []entity.WorkOrderAsset, err error) {
	err = repo.db.SQL("SELECT * FROM work_order_assets WHERE work_order_id = ? AND deleted_at IS null", id).Find(&workOrderAssets)
	return
}

// Update update work_order_asset
func (repo PostgresRepository) Update(id string, changeset entity.WorkOrderAssetChangeSet) error {
	_, err := repo.db.Table("work_order_assets").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find work_order_asset by id
func (repo PostgresRepository) FindByID(id string) (workOrderAsset entity.WorkOrderAsset, err error) {
	_, err = repo.db.SQL("SELECT * FROM work_order_assets WHERE id = ? AND deleted_at IS null", id).Get(&workOrderAsset)
	return
}

// DeleteByID delete work_order_asset by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("work_order_assets").Where("id = ?", id).Delete(&entity.WorkOrderAsset{})
	return err
}

// DeleteByWorkOrderID delete work_order_id by work_order_id
func (repo PostgresRepository) DeleteByWorkOrderID(workOrderID string) error {
	_, err := repo.db.Table("work_order_assets").Where("work_order_id = ?", workOrderID).Delete(&entity.WorkOrderAsset{})
	return err
}
