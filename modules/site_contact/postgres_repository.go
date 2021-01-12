package sitecontact

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

// Save save siteContact to db
func (repo PostgresRepository) Save(siteContact entity.SiteContact) error {
	_, err := repo.db.Table("site_contacts").Insert(&siteContact)
	return err
}

// SaveBatch inserts a batch of siteContact
func (repo PostgresRepository) SaveBatch(siteContacts []entity.SiteContact) error {
	_, err := repo.db.Table("site_contacts").Insert(&siteContacts)
	return err
}

// GetList get list of siteContact
func (repo PostgresRepository) GetList(pagination entity.Pagination) (siteContacts []entity.SiteContact, count int, err error) {
	count, err = repo.db.
		Paginate("site_contacts", &siteContacts, pagination)
	return
}

// Update update siteContact
func (repo PostgresRepository) Update(id string, changeset entity.SiteContactChangeSet) error {
	_, err := repo.db.Table("site_contacts").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find siteContact by id
func (repo PostgresRepository) FindByID(id string) (siteContact entity.SiteContact, err error) {
	_, err = repo.db.SQL("SELECT * FROM site_contacts WHERE id = ? AND deleted_at IS null", id).Get(&siteContact)
	return
}

// DeleteByID delete siteContact by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("site_contacts").Where("id = ?", id).Delete(&entity.SiteContact{})
	return err
}
