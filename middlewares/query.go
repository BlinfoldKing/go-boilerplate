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

	bquery := []byte(query)
	var opts entity.Pagination
	err := json.Unmarshal(bquery, &opts)
	if err != nil {
		helper.CreateErrorResponse(ctx, err).
			BadRequest().
			JSON()
	}

	ctx.Values().Set("pagination", &opts)
	ctx.Next()
}
