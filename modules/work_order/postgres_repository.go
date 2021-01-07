package workorder

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

// Save save work_order to db
func (repo PostgresRepository) Save(workOrder entity.WorkOrder) error {
	_, err := repo.db.Table("work_orders").Insert(&workOrder)
	return err
}

// GetList get list of work_order
func (repo PostgresRepository) GetList(pagination entity.Pagination) (workOrders []entity.WorkOrder, count int, err error) {
	count, err = repo.db.
		Paginate("work_orders", &workOrders, pagination)
	return
}

// Update update work_order
func (repo PostgresRepository) Update(id string, changeset entity.WorkOrderChangeSet) error {
	_, err := repo.db.Table("work_orders").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find work_order by id
func (repo PostgresRepository) FindByID(id string) (workOrder entity.WorkOrder, err error) {
	_, err = repo.db.SQL("SELECT * FROM work_orders WHERE id = ? AND deleted_at IS null", id).Get(&workOrder)
	return
}

// DeleteByID delete work_order by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("work_orders").Where("id = ?", id).Delete(&entity.WorkOrder{})
	return err
}
