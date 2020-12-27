package auth

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"
	"go-boilerplate/modules/otps"
	"go-boilerplate/modules/roles"
	userroles "go-boilerplate/modules/user_roles"
	"go-boilerplate/modules/users"

	"github.com/kataras/iris/v12"
)

const name = "/auth"

// Routes init auth
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	userRepository := users.CreatePosgresRepository(adapters.Postgres)

	roleRepository := roles.CreatePosgresRepository(adapters.Postgres)
	roleService := roles.CreateService(roleRepository)

	userRoleRepository := userroles.CreatePosgresRepository(adapters.Postgres)
	userRoleService := userroles.CreateService(userRoleRepository)

	otpsRepository := otps.CreatePostgresRepository(adapters.Postgres)
	otpsService := otps.CreateService(otpsRepository)

	userService := users.CreateService(userRepository, roleService, userRoleService)

	authService := CreateAuthService(userService, otpsService)

	handler := handler{authService, adapters}

	auth := prefix.Party(name)

	auth.Post("/register", middlewares.ValidateBody(&RegisterRequest{}),
		handler.Register, middlewares.GenerateToken)

	auth.Post("/login", middlewares.ValidateBody(&LoginRequest{}),
		handler.Login, middlewares.GenerateToken)

	auth.Post("/logout", middlewares.InvalidateToken, handler.Logout)

	auth.Post("/reset-password/request", middlewares.ValidateBody(&ResetPasswordRequest{}),
		handler.ResetPasswordRequest)

	auth.Post("/activation:request", middlewares.ValidateBody(&ActivateAccountRequest{}),
		handler.ActivateAccountRequest)
	auth.Post("/activation:verify", handler.VerifyActivationRequest)
}
