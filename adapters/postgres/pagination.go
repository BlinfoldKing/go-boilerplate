package postgres

import (
	"go-boilerplate/entity"
)

// Paginate paginate table
func (postgres Postgres) Paginate(tableName string, data interface{}, opt entity.Pagination) (err error) {
	query, values, err := opt.GetSQL(tableName)
	if err != nil {
		return err
	}
	err = postgres.SQL(query, values...).Find(data)

	return
}
