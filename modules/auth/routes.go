package auth

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"
	"go-boilerplate/modules/roles"
	"go-boilerplate/modules/users"

	"github.com/kataras/iris/v12"
)

const name = "/auth"

// Routes init auth
func Routes(app *iris.Application, adapters adapters.Adapters) {
	userRepository := users.CreatePosgresRepository(adapters.Postgres)
	userService := users.CreateService(userRepository)

	roleRepository := roles.CreatePosgresRepository(adapters.Postgres)
	roleService := roles.CreateService(roleRepository)

	handler := handler{userService, roleService, adapters}

	auth := app.Party(name)

	auth.Post("/register", middlewares.ValidateBody(&RegisterRequest{}),
		handler.Register, middlewares.GenerateToken)

	auth.Post("/login", middlewares.ValidateBody(&LoginRequest{}),
		handler.Login, middlewares.GenerateToken)

	auth.Post("/logout", middlewares.InvalidateToken, handler.Logout)

}
