package modules

import (
	"go-boilerplate/adapters"
	"go-boilerplate/modules/auth"
	"go-boilerplate/modules/ping"

	"github.com/kataras/iris/v12"
)

func Init(app *iris.Application, adapters adapters.Adapters) {
	auth.Routes(app, adapters)
	ping.Routes(app, adapters)
}
