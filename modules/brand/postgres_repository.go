package brand

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

// Save save brand to db
func (repo PostgresRepository) Save(brand entity.Brand) error {
	_, err := repo.db.Table("brands").Insert(&brand)
	return err
}

// GetList get list of brand
func (repo PostgresRepository) GetList(pagination entity.Pagination) (brands []entity.Brand, count int, err error) {
	count, err = repo.db.
		Paginate("brands", &brands, pagination)
	return
}

// Update update brand
func (repo PostgresRepository) Update(id string, changeset entity.BrandChangeSet) error {
	_, err := repo.db.Table("brands").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find brand by id
func (repo PostgresRepository) FindByID(id string) (brand entity.Brand, err error) {
	_, err = repo.db.SQL("SELECT * FROM brands WHERE id = ? AND deleted_at is null", id).Get(&brand)
	return
}

// DeleteByID delete brand by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("brands").Where("id = ?", id).Delete(&entity.Brand{})
	return err
}
