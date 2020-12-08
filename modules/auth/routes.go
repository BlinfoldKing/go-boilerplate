package auth

import (
	"go-boilerplate/adapters"
	"go-boilerplate/modules/users"

	"github.com/kataras/iris/v12"
)

const name = "/auth"

func Routes(app *iris.Application, adapters adapters.Adapters) {
	userRepository := users.CreatePosgresRepository(adapters.Postgres)
	userService := users.CreateService(userRepository)
	handler := handler{userService, adapters}

	auth := app.Party(name)
	auth.Post("/register", handler.Register)
}
