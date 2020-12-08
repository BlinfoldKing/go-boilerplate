package postgres

import (
	"go-boilerplate/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (db *gorm.DB, err error) {
	return gorm.Open(postgres.Open(config.DB_URL()), &gorm.Config{})
}
