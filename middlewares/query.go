package middlewares

import (
	"encoding/json"
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

		var sorts []map[string]interface{}
		var groups []entity.PaginationGroup

		if rawGroup != "" {
			err := json.Unmarshal([]byte(rawGroup), &groups)
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
		for _, group := range groups {
			if group.Desc != nil && *group.Desc {
				sortBy[group.Selector] = "DESC"
			} else {
				sortBy[group.Selector] = "ASC"
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
				Limit:   &take,
				Sort:    &sortBy,
				GroupBy: &groups,
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
