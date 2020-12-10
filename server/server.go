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

	modules.Init(app, adapters)

	jwt, err := middlewares.CreateJWT(adapters.Redis)
	if err != nil {
		logrus.Panic(err)
	}

	app.UseGlobal(middlewares.Logger)
	app.UseGlobal(jwt.AuthenticateToken)

	app.DoneGlobal(jwt.GenerateToken)

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
