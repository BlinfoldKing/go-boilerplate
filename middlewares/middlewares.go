package middlewares

import (
	"go-boilerplate/adapters"
	"go-boilerplate/helper"

	"github.com/kataras/iris/v12"
)

// ValidateBody read and validate body
func ValidateBody(adapters adapters.Adapters, T interface{}) func(iris.Context) {
	return func(ctx iris.Context) {
		err := ctx.ReadJSON(T)
		if err != nil {
			helper.
				CreateErrorResponse(ctx, err).
				InternalServer().
				JSON()
			return
		}

		err = adapters.Validator.Struct(T)
		if err != nil {
			helper.
				CreateErrorResponse(ctx, err).
				InternalServer().
				JSON()
			return
		}

		// return nil
		ctx.Values().Set("body", T)
		ctx.Next()
	}
}
