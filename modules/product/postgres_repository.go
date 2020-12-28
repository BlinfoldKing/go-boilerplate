package product

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

// Save save product to db
func (repo PostgresRepository) Save(product entity.Product) error {
	_, err := repo.db.Table("products").Insert(&product)
	return err
}

// GetList get list of product
func (repo PostgresRepository) GetList(pagination entity.Pagination) (products []entity.Product, count int, err error) {
	count, err = repo.db.
		Paginate("products", &products, pagination)
	return
}

// Update update product
func (repo PostgresRepository) Update(id string, changeset entity.ProductChangeSet) error {
	_, err := repo.db.Table("products").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find product by id
func (repo PostgresRepository) FindByID(id string) (product entity.Product, err error) {
	_, err = repo.db.SQL("SELECT * FROM products WHERE id = ?", id).Get(&product)
	return
}

// DeleteByID delete product by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}
