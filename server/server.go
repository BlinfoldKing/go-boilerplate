package server

import (
	"go-boilerplate/middlewares"
	"go-boilerplate/modules/ping"

	"github.com/kataras/iris/v12"
)

type Server struct {
	Attachment Attachment
	app        *iris.Application
}

func New() Server {
	app := iris.New()
	app.Use(middlewares.Logger)

	ping.Routes(app)

	return Server{
		Attachment{},
		app,
	}
}

func (server Server) Listen() {
	server.app.Run(iris.Addr("0.0.0.0:8000"))
}
