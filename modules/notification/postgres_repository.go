package notification

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

// Save save notification to db
func (repo PostgresRepository) Save(notification entity.Notification) error {
	_, err := repo.db.Table("notifications").Insert(&notification)
	return err
}

// GetList get list of notification
func (repo PostgresRepository) GetList(pagination entity.Pagination) (notifications []entity.Notification, count int, err error) {
	count, err = repo.db.
		Paginate("notifications", &notifications, pagination)
	return
}

// Update update notification
func (repo PostgresRepository) Update(id string, changeset entity.NotificationChangeSet) error {
	_, err := repo.db.Table("notifications").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find notification by id
func (repo PostgresRepository) FindByID(id string) (notification entity.Notification, err error) {
	_, err = repo.db.SQL("SELECT * FROM notifications WHERE id = ?", id).Get(&notification)
	return
}

// DeleteByID delete notification by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Exec("DELETE FROM notifications WHERE id = ?", id)
	return err
}
