package postgres

import (
	"fmt"

	"go-boilerplate/entity"
)

// Paginate paginate table
func (postgres Postgres) Paginate(tableName string, data interface{}, opt entity.Pagination) (err error) {
	query := "SELECT * FROM %s"
	query = fmt.Sprintf(query, tableName)

	values := make([]interface{}, 0)
	var where, sort, limit string = "", "", ""
	if opt.Where != nil {
		var v []interface{}
		where, v, err = parseWhere(*opt.Where)
		if err != nil {
			return
		}

		values = append(values, v...)
		where = fmt.Sprintf("WHERE %s", where)
	}

	if opt.Sort != nil {
		sort, err = parseSort(*opt.Sort)
		if err != nil {
			return
		}

		sort = fmt.Sprintf("ORDER BY %s", sort)
	}

	if opt.Type != nil && *opt.Type == "cursor" {

	} else {
		var v []interface{}
		limit, v = parseLimit(opt.Limit, opt.Offset)
		values = append(values, v...)
	}

	query = fmt.Sprintf("%s %s %s %s", query, where, sort, limit)
	err = postgres.SQL(query, values...).Find(data)

	return
}

func parseLimit(limit *int, offset *int) (query string, val []interface{}) {
	l := 10
	o := 0

	if limit != nil {
		l = *limit
	}

	if offset != nil {
		o = *offset
	}

	return "LIMIT ? OFFSET ?", []interface{}{l, o}
}

func getOperation(key string, op string) (res string, err error) {
	switch op {
	case "lte":
		res = fmt.Sprintf("%s <= ?", key)
	case "lt":
		res = fmt.Sprintf("%s < ?", key)
	case "gte":
		res = fmt.Sprintf("%s >= ?", key)
	case "gt":
		res = fmt.Sprintf("%s > ?", key)
	case "eq":
		res = fmt.Sprintf("%s = ?", key)
	case "neq":
		res = fmt.Sprintf("%s != ?", key)
	case "in":
		res = fmt.Sprintf("%s IN ?", key)
	case "nin":
		res = fmt.Sprintf("%s NOT IN ?", key)
	default:
		err = fmt.Errorf("invalid operator: %s", op)
	}

	return
}

func parseWhere(where map[string]interface{}) (query string, values []interface{}, err error) {
	for key, val := range where {
		switch val.(type) {
		case map[string]interface{}:
			switch key {
			case "or":
				q := ""
				vs := make([]interface{}, 0)
				for k, v := range val.(map[string]interface{}) {
					switch v.(type) {
					case map[string]interface{}:
						for field, v1 := range v.(map[string]interface{}) {
							var where string
							where, err = getOperation(k, field)
							if err != nil {
								return
							}

							if q != "" {
								q = fmt.Sprintf("%s OR %s", q, where)
								vs = append(values, v1)
							} else {
								q = fmt.Sprintf("%s = ?", where)
								vs = append(values, v1)
							}

						}
					default:
						if q != "" {
							q = fmt.Sprintf("%s OR %s = ?", q, k)
							vs = append(vs, v)
						} else {
							q = fmt.Sprintf("%s = ?", q)
							vs = append(vs, v)
						}
					}

				}
				q = fmt.Sprintf("(%s)", q)

				if query != "" {
					query = fmt.Sprintf("%s AND %s = ?", query, q)
					values = append(values, vs...)
				} else {
					query = fmt.Sprintf("%s = ?", q)
				}
			default:
				for field, v := range val.(map[string]interface{}) {
					var where string
					where, err = getOperation(key, field)
					if err != nil {
						return
					}

					if query != "" {
						query = fmt.Sprintf("%s AND %s", query, where)
						values = append(values, v)
					} else {
						query = fmt.Sprintf("%s = ?", where)
						values = append(values, v)
					}

				}
			}
		default:
			if query != "" {
				query = fmt.Sprintf("%s AND %s = ?", query, key)
				values = append(values, val)
			} else {
				query = fmt.Sprintf("%s = ?", key)
				values = append(values, val)
			}
		}
	}

	return
}

func parseSort(sorts map[string]string) (query string, err error) {
	for key, val := range sorts {
		if query != "" {
			query = fmt.Sprintf("%s, %s %s", query, key, val)
		} else {
			query = fmt.Sprintf("%s %s", key, val)
		}
	}
	return
}
