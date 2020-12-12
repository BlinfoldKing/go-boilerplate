package postgres

import (
	"fmt"
	"go-boilerplate/config"

	"xorm.io/xorm"
)

// Postgres extending xorm for extra feature
type Postgres struct {
	*xorm.Engine
}

// Init create data base driver using xorm
func Init() (db *Postgres, err error) {
	engine, _ := xorm.NewEngine("postgres", config.DBCONFIG())
	return &Postgres{engine}, nil
}

// PaginationOpt pagination options
type PaginationOpt struct {
	Limit  *int
	Offset *int
}

// Paginate paginate table
func (postgres Postgres) Paginate(tableName string, data interface{}, opt PaginationOpt) error {
	limit := 0
	offset := 0

	if opt.Limit != nil {
		limit = *opt.Limit
	}

	if opt.Offset != nil {
		offset = *opt.Offset
	}

	query := "SELECT * FROM %s LIMIT $1 OFFSET $2"
	query = fmt.Sprintf(query, tableName)

	err := postgres.
		SQL(query, limit, offset).
		Find(data)

	if err != nil {
		err = fmt.Errorf("query: %s, %s", query, err.Error())
	}
	return err
}
