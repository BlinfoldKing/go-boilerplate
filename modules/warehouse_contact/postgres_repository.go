package warehousecontact

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

// Save save warehouse_contact to db
func (repo PostgresRepository) Save(warehouseContact entity.WarehouseContact) error {
	_, err := repo.db.Table("warehouse_contacts").Insert(&warehouseContact)
	return err
}

// SaveBatch inserts a batch of warehouseContacts
func (repo PostgresRepository) SaveBatch(warehouseContacts []entity.WarehouseContact) error {
	_, err := repo.db.Table("warehouse_contacts").Insert(&warehouseContacts)
	return err
}

// GetList get list of warehouse_contact
func (repo PostgresRepository) GetList(pagination entity.Pagination) (warehouseContacts []entity.WarehouseContact, count int, err error) {
	count, err = repo.db.
		Paginate("warehouse_contacts", &warehouseContacts, pagination)
	return
}

// Update update warehouse_contact
func (repo PostgresRepository) Update(id string, changeset entity.WarehouseContactChangeSet) error {
	_, err := repo.db.Table("warehouse_contacts").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find warehouse_contact by id
func (repo PostgresRepository) FindByID(id string) (warehouseContact entity.WarehouseContact, err error) {
	_, err = repo.db.SQL("SELECT * FROM warehouse_contacts WHERE id = ? AND deleted_at IS null", id).Get(&warehouseContact)
	return
}

// DeleteByID delete warehouse_contact by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("warehouse_contacts").Where("id = ?", id).Delete(&entity.WarehouseContact{})
	return err
}

// DeleteByWarehouseID delete warehouse contact by warehouse id
func (repo PostgresRepository) DeleteByWarehouseID(warehouseID string) error {
	_, err := repo.db.Table("warehouse_contacts").Where("warehouse_id = ?", warehouseID).Delete(&entity.WarehouseContact{})
	return err
}
