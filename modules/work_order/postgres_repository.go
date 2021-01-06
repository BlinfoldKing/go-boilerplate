package work_order

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

// Save save work_order to db
func (repo PostgresRepository) Save(work_order entity.WorkOrder) error {
	_, err := repo.db.Table("work_orders").Insert(&work_order)
	return err
}

// GetList get list of work_order
func (repo PostgresRepository) GetList(pagination entity.Pagination) (work_orders []entity.WorkOrder, count int, err error) {
	count, err = repo.db.
		Paginate("work_orders", &work_orders, pagination)
	return
}

// Update update work_order
func (repo PostgresRepository) Update(id string, changeset entity.WorkOrderChangeSet) error {
	_, err := repo.db.Table("work_orders").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find work_order by id
func (repo PostgresRepository) FindByID(id string) (work_order entity.WorkOrder, err error) {
	_, err = repo.db.SQL("SELECT * FROM work_orders WHERE id = ? AND deleted_at = null", id).Get(&work_order)
	return
}

// DeleteByID delete work_order by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("work_orders").Where("id = ?", id).Delete(&entity.WorkOrder{})
	return err
}
