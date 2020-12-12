package server

import (
	"fmt"
	"go-boilerplate/adapters"
	"go-boilerplate/config"
	"go-boilerplate/middlewares"
	"go-boilerplate/modules"

	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
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
		logrus.Panic(err)
	}

	err = middlewares.InitJWT(adapters)
	if err != nil {
		logrus.Panic(err)
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
		iris.Addr(fmt.Sprintf(":%s", config.PORT())),
		iris.WithoutBodyConsumptionOnUnmarshal,
	)
}
