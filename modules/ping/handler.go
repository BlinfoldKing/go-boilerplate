package ping

import (
	"go-boilerplate/adapters"
	"go-boilerplate/helper"

	"github.com/kataras/iris/v12"
)

type handler struct {
	adapters adapters.Adapters
}

// Ping handle ping request
func (handler handler) Ping(ctx iris.Context) {
	ctx.JSON(iris.Map{"ping": "pong"})

	ctx.Next()
}

// PingNats handle ping request
func (handler handler) PingNats(ctx iris.Context) {
	err := PublishToQueue(Message{
		Name: "world",
	})

	if err != nil {
		helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
		return
	}

	ctx.JSON(iris.Map{"ping": "pong"})

	ctx.Next()
}
