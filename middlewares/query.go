package middlewares

import (
	"encoding/json"
	"fmt"
	"go-boilerplate/entity"
	"go-boilerplate/helper"

	"github.com/kataras/iris/v12"
)

// ValidatePaginationQuery get and validate url params for pagination query
func ValidatePaginationQuery(ctx iris.Context) {
	paginationType := ctx.URLParamDefault("type", "nojson")

	switch paginationType {
	case "inline_offset":
		skip := ctx.URLParamIntDefault("skip", 0)
		take := ctx.URLParamIntDefault("take", 10)
		rawGroup := ctx.URLParamDefault("group", "")
		rawSort := ctx.URLParamDefault("sort", "")
		rawFilter := ctx.URLParamDefault("filter", "")
		rawDistinct := ctx.URLParamDefault("distinct", "")
		rawSelect := ctx.URLParamDefault("select", "")

		var sorts []map[string]interface{}
		var where *map[string]interface{} = nil
		var groups *[]entity.PaginationGroup

		var selection, distinctOn *[]string = nil, nil

		if rawSelect != "" {
			var err error
			selection, err = parseList(rawSelect)
			if err != nil {
				helper.CreateErrorResponse(ctx, err).
					BadRequest().
					JSON()
				return
			}
		}

		if rawDistinct != "" {
			var err error
			distinctOn, err = parseList(rawDistinct)
			if err != nil {
				helper.CreateErrorResponse(ctx, err).
					BadRequest().
					JSON()
				return
			}
		}

		if rawGroup != "" {
			err := json.Unmarshal([]byte(rawGroup), &groups)
			if err != nil {
				helper.CreateErrorResponse(ctx, err).
					BadRequest().
					JSON()
				return
			}
		}

		if rawFilter != "" {
			var err error
			where, err = filterToWhere(rawFilter)
			if err != nil {
				helper.CreateErrorResponse(ctx, err).
					BadRequest().
					JSON()
				return
			}
		}

		if rawSort != "" {
			err := json.Unmarshal([]byte(rawSort), &sorts)
			if err != nil {
				helper.CreateErrorResponse(ctx, err).
					BadRequest().
					JSON()
				return
			}

		}

		var (
			sortBy map[string]string = make(map[string]string)
		)

		if groups != nil {
			for _, group := range *groups {
				if group.Desc != nil && *group.Desc {
					sortBy[group.Selector] = "DESC"
				} else {
					sortBy[group.Selector] = "ASC"
				}
			}
		}

		for _, sort := range sorts {
			selector := sort["selector"].(string)
			if sort["desc"] != nil && sort["desc"].(bool) {
				sortBy[selector] = "DESC"
			} else {
				sortBy[selector] = "ASC"
			}
		}

		opts := entity.OffsetPagination{
			Query: entity.Query{
				Limit:      &take,
				Sort:       &sortBy,
				GroupBy:    groups,
				Where:      where,
				DistinctOn: distinctOn,
				Selection:  selection,
			},
			Offset: &skip,
		}

		ctx.Values().Set("pagination", opts)
	case "cursor":
		query := ctx.URLParam("query")
		bquery := []byte(query)
		var opts entity.CursorPagination
		err := json.Unmarshal(bquery, &opts)

		if opts.Where == nil {
			where := make(map[string]interface{})
			where["deleted_at"] = nil

			opts.Where = &where
		}

		where := *opts.Where
		if _, ok := where["deleted_at"]; !ok {
			where["deleted_at"] = nil
		}

		opts.Where = &where

		if err != nil {
			helper.CreateErrorResponse(ctx, err).
				BadRequest().
				JSON()
			return
		}
		ctx.Values().Set("pagination", opts)
	default:
		query := ctx.URLParam("query")
		bquery := []byte(query)
		var opts entity.OffsetPagination
		err := json.Unmarshal(bquery, &opts)

		if opts.Where == nil {
			where := make(map[string]interface{})
			where["deleted_at"] = nil

			opts.Where = &where
		}

		where := *opts.Where
		if _, ok := where["deleted_at"]; !ok {
			where["deleted_at"] = nil
		}

		opts.Where = &where

		if err != nil {
			helper.CreateErrorResponse(ctx, err).
				BadRequest().
				JSON()
			return
		}
		ctx.Values().Set("pagination", opts)
	}

	ctx.Next()
}

func processFilter(f interface{}) interface{} {
	switch f.(type) {
	case []interface{}:
		filter := f.([]interface{})
		result := make(map[string]interface{})
		switch len(filter) {
		case 0:
			return nil
		case 1:
			return processFilter(filter[0])
		case 2:
			if filter[0] == "!" || filter[0] == "not" {
				op1 := processFilter(filter[1]).(map[string]interface{})

				return map[string]interface{}{
					"not": []interface{}{
						op1,
					},
				}
			}
		case 3:
			switch filter[1].(type) {
			case string:
				if filter[1] == "and" || filter[1] == "or" {
					op1 := processFilter(filter[0]).(map[string]interface{})
					op2 := processFilter(filter[2]).(map[string]interface{})

					res := map[string]interface{}{
						filter[1].(string): []map[string]interface{}{
							op1,
							op2,
						},
					}

					return res

				}

				var res interface{}
				switch filter[1].(string) {
				case "in", "nin":
					res = map[string]interface{}{
						filter[0].(string): map[string]interface{}{
							filter[1].(string): filter[2],
						},
					}
				default:
					res = map[string]interface{}{
						filter[0].(string): map[string]interface{}{
							filter[1].(string): processFilter(filter[2]),
						},
					}
				}

				return res
			}
		}
		return result

	}

	return f
}

func filterToWhere(filterStr string) (where *map[string]interface{}, err error) {
	var data struct {
		Filter []interface{} `json:"filter"`
	}

	filterStr = fmt.Sprintf(`{ "filter": %s }`, filterStr)
	err = json.Unmarshal([]byte(filterStr), &data)
	if err != nil {
		return
	}

	filter := data.Filter
	res := processFilter(filter).(map[string]interface{})
	where = &res

	return
}

func parseList(selectStr string) (s *[]string, err error) {
	var data struct {
		Select []string `json:"select"`
	}

	selectStr = fmt.Sprintf(`{ "select": %s }`, selectStr)
	err = json.Unmarshal([]byte(selectStr), &data)
	if err != nil {
		return
	}

	return &data.Select, nil
}
