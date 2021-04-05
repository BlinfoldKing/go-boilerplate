package workorderproducts

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

// Save save workOrderProducts to db
func (repo PostgresRepository) Save(workOrderProducts entity.WorkOrderProducts) error {
	_, err := repo.db.Table("work_order_products").Insert(&workOrderProducts)
	return err
}

// GetList get list of workOrderProducts
func (repo PostgresRepository) GetList(pagination entity.Pagination) (workorderproducts []entity.WorkOrderProducts, count int, err error) {
	count, err = repo.db.
		Paginate("work_order_products", &workorderproducts, pagination)
	return
}

// Update update workOrderProducts
func (repo PostgresRepository) Update(id string, changeset entity.WorkOrderProductsChangeSet) error {
	_, err := repo.db.Table("work_order_products").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find workOrderProducts by id
func (repo PostgresRepository) FindByID(id string) (workOrderProducts entity.WorkOrderProducts, err error) {
	_, err = repo.db.SQL("SELECT * FROM work_order_products WHERE id = ? AND deleted_at IS null", id).Get(&workOrderProducts)
	return
}

// DeleteByID delete workOrderProducts by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("work_order_products").Where("id = ?", id).Delete(&entity.WorkOrderProducts{})
	return err
}
