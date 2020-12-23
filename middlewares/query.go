package middlewares

import (
	"encoding/json"
	"go-boilerplate/entity"
	"go-boilerplate/helper"

	"github.com/kataras/iris/v12"
)

// ValidatePaginationQuery get and validate url params for pagination query
func ValidatePaginationQuery(ctx iris.Context) {
	query := ctx.URLParam("query")
	paginationType := ctx.URLParamDefault("type", "offset")

	switch paginationType {
	case "cursor":
		bquery := []byte(query)
		var opts entity.CursorPagination
		err := json.Unmarshal(bquery, &opts)
		if err != nil {
			helper.CreateErrorResponse(ctx, err).
				BadRequest().
				JSON()
			return
		}
		ctx.Values().Set("pagination", opts)
	default:
		bquery := []byte(query)
		var opts entity.OffsetPagination
		err := json.Unmarshal(bquery, &opts)
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
