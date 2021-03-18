package sensorlog

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

// Save save sensorLog to db
func (repo PostgresRepository) Save(sensorLog entity.SensorLog) error {
	_, err := repo.db.Table("sensor_logs").Insert(&sensorLog)
	return err
}

// GetList get list of sensorLog
func (repo PostgresRepository) GetList(pagination entity.Pagination) (sensorlogs []entity.SensorLog, count int, err error) {
	count, err = repo.db.
		Paginate("sensor_logs", &sensorlogs, pagination)
	return
}

// Update update sensorLog
func (repo PostgresRepository) Update(id string, changeset entity.SensorLogChangeSet) error {
	_, err := repo.db.Table("sensor_logs").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find sensorLog by id
func (repo PostgresRepository) FindByID(id string) (sensorLog entity.SensorLog, err error) {
	_, err = repo.db.SQL("SELECT * FROM sensor_logs WHERE id = ? AND deleted_at IS null", id).Get(&sensorLog)
	return
}

// DeleteByID delete sensorLog by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("sensor_logs").Where("id = ?", id).Delete(&entity.SensorLog{})
	return err
}
