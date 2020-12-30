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
func CreatePosgresRepository(db *postgres.Postgres) Repository {
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
	_, err = repo.db.SQL("SELECT * FROM product_specifications WHERE id = ? AND deleted_at = nil", id).Get(&productSpecification)
	return
}

// DeleteByID delete productSpecification by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("product_specifications").Where("id = ?", id).Delete(&entity.ProductSpecification{})
	return err
}
