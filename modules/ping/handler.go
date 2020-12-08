package ping

import (
	"go-boilerplate/adapters"

	"github.com/kataras/iris/v12"
)

type handler struct {
	adapters adapters.Adapters
}

func (handler handler) Ping(ctx iris.Context) {
	ctx.JSON(iris.Map{"ping": "pong"})

	ctx.Next()
}
