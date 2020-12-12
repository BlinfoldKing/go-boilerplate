package auth

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"
	"go-boilerplate/modules/users"

	"github.com/kataras/iris/v12"
)

const name = "/auth"

// Routes init auth
func Routes(app *iris.Application, adapters adapters.Adapters) {
	userRepository := users.CreatePosgresRepository(adapters.Postgres)
	userService := users.CreateService(userRepository)
	handler := handler{userService, adapters}

	auth := app.Party(name)

	auth.Post("/register", middlewares.ValidateBody(adapters, &RegisterRequest{}),
		handler.Register, middlewares.GenerateToken)

	auth.Post("/login", middlewares.ValidateBody(adapters, &LoginRequest{}),
		handler.Login, middlewares.GenerateToken)

	auth.Post("/logout", middlewares.InvalidateToken, handler.Logout)

}
