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

type Server struct {
	app *iris.Application
}

func New() Server {
	app := iris.New()
	app.Use(middlewares.Logger)

	adapters, err := adapters.Init()
	if err != nil {
		logrus.Panic(err)
	}

	modules.Init(app, adapters)

	return Server{
		app,
	}

}

func (server Server) Listen() {
	server.app.Run(
		iris.Addr(fmt.Sprintf(":%s", config.PORT())),
		iris.WithoutBodyConsumptionOnUnmarshal,
	)
}
