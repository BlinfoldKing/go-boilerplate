package site

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

// Save save site to db
func (repo PostgresRepository) Save(site entity.Site) error {
	_, err := repo.db.Table("sites").Insert(&site)
	return err
}

// GetList get list of site
func (repo PostgresRepository) GetList(pagination entity.Pagination) (sites []entity.Site, count int, err error) {
	count, err = repo.db.
		Paginate("sites", &sites, pagination)
	return
}

// Update update site
func (repo PostgresRepository) Update(id string, changeset entity.SiteChangeSet) error {
	_, err := repo.db.Table("sites").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find site by id
func (repo PostgresRepository) FindByID(id string) (site entity.Site, err error) {
	_, err = repo.db.SQL("SELECT * FROM sites WHERE id = ? AND deleted_at IS null", id).Get(&site)

	return
}

// DeleteByID delete site by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("sites").Where("id = ?", id).Delete(&entity.Site{})
	return err
}
