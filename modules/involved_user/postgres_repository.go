package involveduser

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

// Save save involved_user to db
func (repo PostgresRepository) Save(involvedUser entity.InvolvedUser) error {
	_, err := repo.db.Table("involved_users").Insert(&involvedUser)
	return err
}

// SaveBatch inserts a batch of involved users
func (repo PostgresRepository) SaveBatch(involvedUsers []entity.InvolvedUser) error {
	_, err := repo.db.Table("involved_users").Insert(&involvedUsers)
	return err
}

// GetList get list of involved_user
func (repo PostgresRepository) GetList(pagination entity.Pagination) (involvedUsers []entity.InvolvedUser, count int, err error) {
	count, err = repo.db.
		Paginate("involved_users", &involvedUsers, pagination)
	return
}

// Update update involved_user
func (repo PostgresRepository) Update(id string, changeset entity.InvolvedUserChangeSet) error {
	_, err := repo.db.Table("involved_users").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find involved_user by id
func (repo PostgresRepository) FindByID(id string) (involvedUser entity.InvolvedUser, err error) {
	_, err = repo.db.SQL("SELECT * FROM involved_users WHERE id = ? AND deleted_at IS null", id).Get(&involvedUser)
	return
}

// DeleteByID delete involved_user by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("involved_users").Where("id = ?", id).Delete(&entity.InvolvedUser{})
	return err
}

// DeleteByWorkOrderID delete involved user by work order id
func (repo PostgresRepository) DeleteByWorkOrderID(workOrderID string) error {
	_, err := repo.db.Table("involved_users").Where("work_order_id = ?", workOrderID).Delete(&entity.WorkOrderDocument{})
	return err
}
