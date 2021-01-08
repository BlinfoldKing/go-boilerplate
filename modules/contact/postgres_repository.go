package contact

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

// Save save contact to db
func (repo PostgresRepository) Save(contact entity.Contact) error {
	_, err := repo.db.Table("contacts").Insert(&contact)
	return err
}

// GetList get list of contact
func (repo PostgresRepository) GetList(pagination entity.Pagination) (contacts []entity.Contact, count int, err error) {
	count, err = repo.db.
		Paginate("contacts", &contacts, pagination)
	return
}

// FindByWarehouseID gets list of contact that associated with specified brand
func (repo PostgresRepository) FindByWarehouseID(warehouseID string) (contacts []entity.Contact, err error) {
	err = repo.db.
		SQL(`SELECT 
				c.*
			FROM 
				warehouse_contacts wc
			INNER JOIN contacts c
				ON wc.warehouse_id = ?
				AND wc.contact_id = c.id
				AND deleted_at IS NULL`,
			warehouseID).Find(&contacts)
	return
}

// Update update contact
func (repo PostgresRepository) Update(id string, changeset entity.ContactChangeSet) error {
	_, err := repo.db.Table("contacts").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find contact by id
func (repo PostgresRepository) FindByID(id string) (contact entity.Contact, err error) {
	_, err = repo.db.SQL("SELECT * FROM contacts WHERE id = ? AND deleted_at IS null", id).Get(&contact)
	return
}

// DeleteByID delete contact by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("contacts").Where("id = ?", id).Delete(&entity.Contact{})
	return err
}
