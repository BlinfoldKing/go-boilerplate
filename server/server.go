package server

import (
	"fmt"
	"go-boilerplate/adapters"
	"go-boilerplate/config"
	"go-boilerplate/helper"
	"go-boilerplate/middlewares"
	"go-boilerplate/modules"

	"github.com/kataras/iris/v12"
)

// Server wrapper
type Server struct {
	app     *iris.Application
	Adapter adapters.Adapters
}

// New create new user
func New() Server {
	app := iris.New()

	adapters, err := adapters.Init()
	if err != nil {
		helper.Logger.Warn(err)
	}

	err = middlewares.InitValidator(adapters)
	if err != nil {
		helper.Logger.Panic(err)
	}

	err = middlewares.InitJWT(adapters)
	if err != nil {
		helper.Logger.Panic(err)
	}

	app.UseGlobal(middlewares.Logger)
	app.UseGlobal(middlewares.AuthenticateToken)

	modules.Init(app, adapters)

	return Server{
		app,
		adapters,
	}

}

// Listen start server
func (server Server) Listen() {
	server.app.Run(
		iris.Addr(fmt.Sprintf("%s:%s", config.APPURL(), config.PORT())),
		iris.WithoutBodyConsumptionOnUnmarshal,
	)
}
