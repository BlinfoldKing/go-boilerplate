package middlewares

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func Logger(ctx iris.Context) {
	body, _ := ctx.GetBody()
	log := fmt.Sprintf(
		"path: %s, method: %s, body: %s",
		ctx.Path(), ctx.Request().Method, body)
	ctx.Application().Logger().Info(log)

	ctx.Next()
}
