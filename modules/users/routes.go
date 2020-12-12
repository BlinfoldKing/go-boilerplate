package users

import (
	"go-boilerplate/adapters"

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
}
