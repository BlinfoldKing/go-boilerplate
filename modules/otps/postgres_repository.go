package otps

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

// Save save otp to db
func (repo PostgresRepository) Save(otp entity.OTP) error {
	_, err := repo.db.Table("otps").Insert(&otp)
	return err
}
