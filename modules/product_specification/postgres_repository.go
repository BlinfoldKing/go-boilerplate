package productspecification

import (
	"go-boilerplate/adapters/postgres"
	"go-boilerplate/entity"
)

// PostgresRepository repository implementation on postgres
type PostgresRepository struct {
	db *postgres.Postgres
}

// CreatePosgresRepository init PostgresRepository
func CreatePostgresRepository(db *postgres.Postgres) Repository {
	return PostgresRepository{db}
}

// Save save productSpecification to db
func (repo PostgresRepository) Save(productSpecification entity.ProductSpecification) error {
	_, err := repo.db.Table("product_specifications").Insert(&productSpecification)
	return err
}

// GetList get list of productSpecification
func (repo PostgresRepository) GetList(pagination entity.Pagination) (productSpecifications []entity.ProductSpecification, count int, err error) {
	count, err = repo.db.
		Paginate("product_specifications", &productSpecifications, pagination)
	return
}

// Update update productSpecification
func (repo PostgresRepository) Update(id string, changeset entity.ProductSpecificationChangeSet) error {
	_, err := repo.db.Table("product_specifications").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find productSpecification by id
func (repo PostgresRepository) FindByID(id string) (productSpecification entity.ProductSpecification, err error) {
	_, err = repo.db.SQL("SELECT * FROM product_specifications WHERE id = ? AND deleted_at = null", id).Get(&productSpecification)
	return
}

// FindByProductID find productSpecification by product_id
func (repo PostgresRepository) FindByProductID(productID string) (productSpecifications []entity.ProductSpecification, err error) {
	err = repo.db.SQL("SELECT * FROM product_specifications WHERE product_id = ? AND deleted_at = null", productID).Find(&productSpecifications)
	return
}

// DeleteByID delete productSpecification by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("product_specifications").Where("id = ?", id).Delete(&entity.ProductSpecification{})
	return err
}

// DeleteByProductID delete product_specification by product_id
func (repo PostgresRepository) DeleteByProductID(productID string) error {
	_, err := repo.db.Table("product_specifications").Where("product_id = ?", productID).Delete(&entity.Product{})
	return err
}
