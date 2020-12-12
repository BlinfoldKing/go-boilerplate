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
func (repo PostgresRepository) GetList(limit, offset int) (users []entity.User, err error) {
	err = repo.db.
		Paginate("users", &users, postgres.PaginationOpt{
			Limit:  &limit,
			Offset: &offset,
		})
	return
}
