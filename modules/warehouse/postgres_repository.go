package warehouse

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

// Save save warehouse to db
func (repo PostgresRepository) Save(warehouse entity.Warehouse) error {
	_, err := repo.db.Table("warehouses").Insert(&warehouse)
	return err
}

// GetList get list of warehouse
func (repo PostgresRepository) GetList(pagination entity.Pagination) (warehouses []entity.Warehouse, count int, err error) {
	count, err = repo.db.
		Paginate("warehouses", &warehouses, pagination)
	return
}

// Update update warehouse
func (repo PostgresRepository) Update(id string, changeset entity.WarehouseChangeSet) error {
	_, err := repo.db.Table("warehouses").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find warehouse by id
func (repo PostgresRepository) FindByID(id string) (warehouse entity.Warehouse, err error) {
	_, err = repo.db.SQL("SELECT * FROM warehouses WHERE id = ? AND deleted_at IS null", id).Get(&warehouse)
	return
}

// DeleteByID delete warehouse by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("warehouses").Where("id = ?", id).Delete(&entity.Warehouse{})
	return err
}

// GetAllWarehousebyAssetID get all warehouse
func (repo PostgresRepository) GetAllWarehousebyAssetID(id string) (warehouse []entity.Warehouse, err error) {
	err = repo.db.
		SQL(`SELECT
				w.*
			FROM
				warehouse w
			INNER JOIN asset_warehouses aw
				ON aw.asset_id = ?
			`, id).Find(&warehouse)

	return
}
