package sensor

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

// Save save sensor to db
func (repo PostgresRepository) Save(sensor entity.Sensor) error {
	_, err := repo.db.Table("sensors").Insert(&sensor)
	return err
}

// GetList get list of sensor
func (repo PostgresRepository) GetList(pagination entity.Pagination) (sensors []entity.Sensor, count int, err error) {
	count, err = repo.db.
		Paginate("sensors", &sensors, pagination)
	return
}

// Update update sensor
func (repo PostgresRepository) Update(id string, changeset entity.SensorChangeSet) error {
	_, err := repo.db.Table("sensors").Where("id = ?", id).Update(&changeset)
	return err
}

// FindByID find sensor by id
func (repo PostgresRepository) FindByID(id string) (sensor entity.Sensor, err error) {
	_, err = repo.db.SQL("SELECT * FROM sensors WHERE id = ? AND deleted_at IS null", id).Get(&sensor)
	return
}

// DeleteByID delete sensor by id
func (repo PostgresRepository) DeleteByID(id string) error {
	_, err := repo.db.Table("sensors").Where("id = ?", id).Delete(&entity.Sensor{})
	return err
}
