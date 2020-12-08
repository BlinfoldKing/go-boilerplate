package ping

import "github.com/kataras/iris/v12"

func Ping(ctx iris.Context) {
	ctx.JSON(iris.Map{"ping": "pong"})

	ctx.Next()
}
