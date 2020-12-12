package users

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/users"

// Routes init users
func Routes(app *iris.Application, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}

	users := app.Party(name)

	users.Get("/", handler.GetList)
	users.Get("/{id:string}", handler.GetByID)
	users.Delete("/{id:string}", handler.DeleteByID)
	users.Put("/", middlewares.ValidateBody(&UpdateRequest{}),
		handler.Update)
}
