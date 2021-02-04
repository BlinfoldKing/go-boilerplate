package auth

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"
	"go-boilerplate/modules/otps"
	"go-boilerplate/modules/users"

	"github.com/kataras/iris/v12"
)

const name = "/auth"

// Routes init auth
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	otpsRepository := otps.CreatePostgresRepository(adapters.Postgres)
	otpsService := otps.CreateService(otpsRepository)

	userService := users.InitUserService(adapters)

	authService := CreateAuthService(userService, otpsService)

	handler := handler{authService, adapters}

	auth := prefix.Party(name)

	auth.Post("/register", middlewares.ValidateBody(&RegisterRequest{}),
		handler.Register, middlewares.GenerateToken)

	auth.Post("/login", middlewares.ValidateBody(&LoginRequest{}),
		handler.Login, middlewares.GenerateToken)

	auth.Post("/logout", middlewares.InvalidateToken, handler.Logout)

	auth.Post("/password:request", middlewares.ValidateBody(&ResetPasswordRequest{}),
		handler.ResetPasswordRequest)

	auth.Post("/password:submit", middlewares.ValidateBody(&ResetPasswordSubmit{}),
		handler.ResetPassword)

	auth.Post("/activation:request", middlewares.ValidateBody(&ActivateAccountRequest{}),
		handler.ActivateAccountRequest)
	auth.Post("/activation:verify", handler.VerifyActivationRequest)
	auth.Post("/refresh", middlewares.RefreshToken)
}
