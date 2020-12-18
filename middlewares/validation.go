package middlewares

import (
	"github.com/kataras/iris/v12"
	"go-boilerplate/adapters"
	"go-boilerplate/helper"
)

// ValidateBody read and validate body
var ValidateBody func(T interface{}) func(iris.Context)

// InitValidator ini middleware
func InitValidator(adapters adapters.Adapters) error {
	ValidateBody = func(T interface{}) func(iris.Context) {
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

	return nil
}
