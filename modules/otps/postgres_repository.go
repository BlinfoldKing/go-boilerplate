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

// FindByToken gets otp based on token
func (repo PostgresRepository) FindByToken(token string) (otp entity.OTP, err error) {
	_, err = repo.db.SQL("SELECT * FROM otps WHERE token = ? AND deleted_at IS NULL", token).Get(&otp)
	return
}

// FindByTokenAndEmail gets otp based on token and email
func (repo PostgresRepository) FindByTokenAndEmail(token, email string) (otp entity.OTP, err error) {
	_, err = repo.db.SQL("SELECT * FROM otps WHERE token = ? AND email = ? AND deleted_at IS NULL", token, email).Get(&otp)
	return
}

// DeleteByToken deletes otp based on token
func (repo PostgresRepository) DeleteByToken(token string) (err error) {
	_, err = repo.db.Table("otps").Where("token = ?", token).Delete(&entity.OTP{})
	return err
}
