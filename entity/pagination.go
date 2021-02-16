package entity

import (
	"fmt"
	"go-boilerplate/entity/common"
)

// Pagination pagination interace
type Pagination interface {
	GetSQL(tableName string) (sql string, args []interface{}, err error)
	GetWhere() (query string, args []interface{})
}

// PaginationGroup group pagination
type PaginationGroup struct {
	Selector   string `json:"selector"`
	Desc       *bool  `json:"desc"`
	IsExpanded *bool  `json:"isExpanded"`
}

// Query shared pagination fields
type Query struct {
	Limit      *int                    `json:"limit"`
	Sort       *map[string]string      `json:"sort"`
	Where      *map[string]interface{} `json:"where"`
	DistinctOn *[]string
	GroupBy    *[]PaginationGroup
	Selection  *[]string `json:"selection"`
}

// OffsetPagination pagination parameters
type OffsetPagination struct {
	Query
	Offset *int `json:"offset"`
}

// GetWhere get where
func (opt OffsetPagination) GetWhere() (query string, args []interface{}) {
	if opt.Where != nil {
		query, args, _ = parseWhere(*opt.Where)
	}
	return
}

// GetSQL generate sql
func (opt OffsetPagination) GetSQL(tableName string) (sql string, args []interface{}, err error) {
	query := "SELECT %s * FROM %s"
	query = fmt.Sprintf(query, parseDistinct(opt.DistinctOn), tableName)

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

	var group string
	if opt.GroupBy != nil {
		group, err = parseGroupBy(*opt.GroupBy)
		if err != nil {
			return
		}
	}

	if opt.Sort != nil && len(*opt.Sort) > 0 {
		sort, err = parseSort(*opt.Sort)
		if err != nil {
			return
		}

		sort = fmt.Sprintf("ORDER BY %s", sort)
	}

	var v []interface{}
	limit, v = parseLimit(opt.Limit, opt.Offset)
	values = append(values, v...)

	sql = fmt.Sprintf("%s %s %s %s %s", query, where, group, sort, limit)
	args = values

	return
}

// CursorPagination pagination parameters
type CursorPagination struct {
	Query
	ID   string  `json:"id"`
	Seek *string `json:"seek"`
}

// GetWhere get where
func (opt CursorPagination) GetWhere() (query string, args []interface{}) {
	if opt.Where != nil {
		query, args, _ = parseWhere(*opt.Where)
	}
	return
}

// GetSQL generate sql
func (opt CursorPagination) GetSQL(tableName string) (sql string, args []interface{}, err error) {
	query := ""
	if opt.Seek != nil && *opt.Seek == "prev" {
		query = `SELECT %s * FROM %s WHERE "order" <= (SELECT "order" FROM %s WHERE id = ?)`
	} else {
		query = `SELECT %s * FROM %s WHERE "order" >= (SELECT "order" FROM %s WHERE id = ?)`
	}
	query = fmt.Sprintf(query, parseDistinct(opt.DistinctOn), tableName, tableName)
	query = fmt.Sprintf("SELECT * FROM (%s)", query)

	values := []interface{}{
		opt.ID,
	}

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

	var group string
	if opt.GroupBy != nil {
		group, err = parseGroupBy(*opt.GroupBy)
		if err != nil {
			return
		}
	}

	if opt.Sort != nil && len(*opt.Sort) > 0 {
		sort, err = parseSort(*opt.Sort)
		if err != nil {
			return
		}

		sort = fmt.Sprintf("ORDER BY %s", sort)

		if opt.Seek != nil && *opt.Seek == "prev" {
			sort += ", \"order\" DESC"
		}
	} else {
		if opt.Seek != nil && *opt.Seek == "prev" {
			sort = "ORDER BY \"order\" DESC"
		}

	}

	var v []interface{}
	limit, v = parseLimit(opt.Limit, nil)
	values = append(values, v...)

	sql = fmt.Sprintf("%s %s %s %s %s", query, where, group, sort, limit)
	args = values

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

func parseBoolOperator(operator string, items []map[string]interface{}) (query string, values []interface{}, err error) {
	switch operator {
	case "not":
		for _, item := range items {
			var q string
			var v []interface{}
			q, v, err = parseWhere(item)
			if err != nil {
				return
			}

			if query == "" {
				query = q
			} else {
				query = fmt.Sprintf("%s OR %s", query, q)
			}

			values = append(values, v...)
		}
		query = fmt.Sprintf("NOT (%s)", query)
	case "and":
		for _, item := range items {
			var q string
			var v []interface{}
			q, v, err = parseWhere(item)
			if err != nil {
				return
			}

			if query == "" {
				query = q
			} else {
				query = fmt.Sprintf("%s AND %s", query, q)
			}

			values = append(values, v...)
		}
	case "or":
		for _, item := range items {
			var q string
			var v []interface{}
			q, v, err = parseWhere(item)
			if err != nil {
				return
			}

			if query == "" {
				query = q
			} else {
				query = fmt.Sprintf("%s OR %s", query, q)
			}

			values = append(values, v...)
		}
	}
	return
}

func getOperation(key string, op string, value interface{}) (res string, err error) {

	switch op {
	case "lte", "<=":
		res = fmt.Sprintf("%s <= ?", key)
	case "lt", "<":
		res = fmt.Sprintf("%s < ?", key)
	case "gte", ">=":
		res = fmt.Sprintf("%s >= ?", key)
	case "gt", ">":
		res = fmt.Sprintf("%s > ?", key)
	case "eq", "=", "is":
		if value == nil {
			res = fmt.Sprintf("%s IS NULL", key)
		} else {
			res = fmt.Sprintf("%s = ?", key)
		}
	case "neq", "!=", "is not":
		if value == nil {
			res = fmt.Sprintf("%s IS NOT NULL", key)
		} else {
			res = fmt.Sprintf("%s != ?", key)
		}
	case "in":
		res = fmt.Sprintf("%s = ANY(?)", key)
	case "nin":
		res = fmt.Sprintf("NOT %s = ANY(?)", key)
	case "startWith":
		res = key + " LIKE ? || '%'"
	case "endWith":
		res = key + " LIKE '%' || ?"
	case "contains":
		res = key + " LIKE '%' || ? || '%'"
	default:
		if value == nil {
			res = fmt.Sprintf("%s IS NULL", key)
		} else {
			res = fmt.Sprintf("%s = ?", key)
		}
	}

	return
}

func parseValueOperator(attribute string, val interface{}) (query string, values []interface{}, err error) {
	switch val.(type) {
	case map[string]interface{}:
		for op, v := range val.(map[string]interface{}) {
			q, err := getOperation(attribute, op, v)
			if err != nil {
				return query, values, err
			}

			if query == "" {
				query = q
			} else {
				query = fmt.Sprintf("%s AND %s", query, q)
			}

			if v != nil {
				switch v.(type) {
				case []interface{}:
					switch v.([]interface{})[0].(type) {
					case string:
						arr := make(common.StrArr, 0)
						for _, item := range v.([]interface{}) {
							arr = append(arr, item.(string))
						}

						s, _ := arr.ToDB()
						v = string(s)
					}
				}
				values = append(values, v)
			}
		}
	default:
		q, err := getOperation(attribute, "=", val)
		if err != nil {
			return query, values, err
		}
		if query == "" {
			query = q
		} else {
			query = fmt.Sprintf("%s AND %s", query, q)
		}

		if val != nil {
			switch val.(type) {
			case []interface{}:
				switch val.([]interface{})[0].(type) {
				case string:
					arr := make(common.StrArr, 0)
					for _, item := range val.([]interface{}) {
						arr = append(arr, item.(string))
					}

					s, _ := arr.ToDB()
					val = string(s)
				}
			}
			values = append(values, val)
		}
	}

	return
}

func parseWhere(where map[string]interface{}) (query string, values []interface{}, err error) {
	query = ""
	for key, val := range where {
		switch key {
		case "and", "or", "not":
			switch val.(type) {
			case []map[string]interface{}:
				q, value, err := parseBoolOperator(key, val.([]map[string]interface{}))
				if err != nil {
					return query, values, err
				}

				if query == "" {
					query = q
				} else {
					query = fmt.Sprintf("%s AND %s", query, q)
				}

				values = append(values, value...)
			default:
				err = fmt.Errorf("invalid value for %s", key)
				return
			}
		default:
			q, value, err := parseValueOperator(key, val)
			if err != nil {
				return query, values, err
			}

			if query == "" {
				query = q
			} else {
				query = fmt.Sprintf("%s AND %s", query, q)
			}

			values = append(values, value...)

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

func parseGroupBy(groups []PaginationGroup) (query string, err error) {
	idFound := false
	for _, group := range groups {
		idFound = idFound || group.Selector == "id"
		if query == "" {
			query = fmt.Sprintf("GROUP BY %s", group.Selector)
		} else {
			query = fmt.Sprintf("%s, %s", query, group.Selector)
		}
	}

	if !idFound {
		if query == "" {
			query = fmt.Sprintf("GROUP BY %s", "id")
		} else {
			query = fmt.Sprintf("%s, %s", query, "id")
		}
	}
	return
}

func parseDistinct(distincs *[]string) string {
	if distincs == nil || len(*distincs) < 1 {
		return ""
	}
	res := ""
	for _, item := range *distincs {
		if res == "" {
			res = item
		} else {
			res = fmt.Sprintf("%s, %s", res, item)
		}
	}

	return fmt.Sprintf("DISTINCT ON (%s)", res)
}
