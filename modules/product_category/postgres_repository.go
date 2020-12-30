package productcategory

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

// Save save productCategory to db
func (repo PostgresRepository) Save(productCategory entity.ProductCategory) error {
	_, err := repo.db.Table("product_categories").Insert(&productCategory)
	return err
}

// GetList get list of productCategory
func (repo PostgresRepository) GetList(pagination entity.Pagination) (productCategories []entity.ProductCategory, count int, err error) {
	count, err = repo.db.
		Paginate("product_categories", &productCategories, pagination)
	return
}

// Update update productCategory
func (repo PostgresRepository) Update(id string, changeset entity.ProductCategoryChangeSet) error {
	_, err := repo.db.Table("product_categories").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find productCategory by id
func (repo PostgresRepository) FindByID(id string) (productCategory entity.ProductCategory, err error) {
	_, err = repo.db.SQL("SELECT * FROM product_categories WHERE id = ?", id).Get(&productCategory)
	return
}

// DeleteByID delete productCategory by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Exec("DELETE FROM product_categories WHERE id = ?", id)
	return err
}
