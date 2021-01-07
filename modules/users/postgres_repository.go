package users

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

// Save save user to db
func (repo PostgresRepository) Save(user entity.User) error {
	_, err := repo.db.Table("users").Insert(&user)
	return err
}

// FindByEmail find user by email
func (repo PostgresRepository) FindByEmail(email string) (user entity.User, err error) {
	_, err = repo.db.SQL("SELECT * FROM users WHERE email = ?", email).Get(&user)
	return
}

// GetList get list of users
func (repo PostgresRepository) GetList(pagination entity.Pagination) (users []entity.User, count int, err error) {
	count, err = repo.db.
		Paginate("users", &users, pagination)
	return
}

// Update update user
func (repo PostgresRepository) Update(id string, changeset entity.UserChangeSet) error {
	_, err := repo.db.Table("users").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find user by id
func (repo PostgresRepository) FindByID(id string) (user entity.User, err error) {
	_, err = repo.db.SQL("SELECT * FROM users WHERE id = ?", id).Get(&user)
	return
}

// FindByWorkOrderID find users by work order id
func (repo PostgresRepository) FindByWorkOrderID(workOrderID string) (users []entity.User, err error) {
	err = repo.db.
		SQL(`SELECT 
				u.*
			FROM 
				involved_users iu
			INNER JOIN users u
				ON iu.work_order_id = ?
				AND iu.user_id = u.id`,
			workOrderID).Find(&users)
	return
}

// DeleteByID delete user by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
