package roles

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

// Save save role to db
func (repo PostgresRepository) Save(role entity.Role) error {
	_, err := repo.db.Table("roles").Insert(&role)
	return err
}

// FindBySlug find role by slug
func (repo PostgresRepository) FindBySlug(slug string) (role entity.Role, err error) {
	_, err = repo.db.SQL("SELECT * FROM roles WHERE slug = ?", slug).Get(&role)
	return
}

// GetList get list of users
func (repo PostgresRepository) GetList(pagination entity.Pagination) (roles []entity.Role, count int, err error) {
	count, err = repo.db.
		Paginate("roles", &roles, pagination)
	return
}

// Update update user
func (repo PostgresRepository) Update(id string, changeset entity.RoleChangeSet) error {
	_, err := repo.db.Table("roles").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find user by id
func (repo PostgresRepository) FindByID(id string) (role entity.Role, err error) {
	_, err = repo.db.SQL("SELECT * FROM roles WHERE id = ?", id).Get(&role)
	return
}

// DeleteByID delete user by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Exec("DELETE FROM roles WHERE id = ?", id)
	return err
}
