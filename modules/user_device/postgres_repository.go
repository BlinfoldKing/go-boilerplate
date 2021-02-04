package userdevice

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

// Save save user_device to db
func (repo PostgresRepository) Save(userdevice entity.UserDevice) error {
	_, err := repo.db.Table("user_devices").Insert(&userdevice)
	return err
}

// GetList get list of user_device
func (repo PostgresRepository) GetList(pagination entity.Pagination) (userdevices []entity.UserDevice, count int, err error) {
	count, err = repo.db.
		Paginate("user_devices", &userdevices, pagination)
	return
}

// Update update user_device
func (repo PostgresRepository) Update(id string, changeset entity.UserDeviceChangeSet) error {
	_, err := repo.db.Table("user_devices").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find user_device by id
func (repo PostgresRepository) FindByID(id string) (userdevice entity.UserDevice, err error) {
	_, err = repo.db.SQL("SELECT * FROM user_devices WHERE id = ? AND deleted_at IS null", id).Get(&userdevice)
	return
}

// FindByUserID find user_device by id
func (repo PostgresRepository) FindByUserID(userid string) (userdevice []entity.UserDevice, err error) {
	_, err = repo.db.SQL("SELECT * FROM user_devices WHERE user_id = ? AND deleted_at IS null", userid).Get(&userdevice)
	return
}

// DeleteByID delete user_device by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("user_devices").Where("id = ?", id).Delete(&entity.UserDevice{})
	return err
}
