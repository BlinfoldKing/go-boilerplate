package documents

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

// Save document to db
func (repo PostgresRepository) Save(document entity.Document) error {
	_, err := repo.db.Table("documents").Insert(&document)
	return err
}

// FindByID find document by id
func (repo PostgresRepository) FindByID(id string) (document entity.Document, err error) {
	_, err = repo.db.SQL("SELECT * FROM documents WHERE id = ?", id).Get(&document)
	return
}
