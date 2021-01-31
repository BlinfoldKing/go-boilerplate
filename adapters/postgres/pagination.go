package postgres

import (
	"go-boilerplate/entity"
)

type withID struct {
	ID string
}

// Paginate paginate table
func (postgres Postgres) Paginate(tableName string, data interface{}, opt entity.Pagination) (count int, err error) {
	query, values, err := opt.GetSQL(tableName)
	if err != nil {
		return
	}

	err = postgres.SQL(query, values...).Find(data)
	if err != nil {
		return
	}

	query, values = opt.GetWhere()
	c, err := postgres.Table(tableName).Where(query, values...).Count()
	count = int(c)
	return
}
