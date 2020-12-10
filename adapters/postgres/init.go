package postgres

import (
	"go-boilerplate/config"
	"xorm.io/xorm"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

// Init create data base driver using xorm
func Init() (db *xorm.Engine, err error) {
	return xorm.NewEngine("postgres", config.DBURL())
}
