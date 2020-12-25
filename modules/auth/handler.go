package auth

import (
	"fmt"
	"go-boilerplate/adapters"
	"go-boilerplate/config"
	"go-boilerplate/helper"
	"go-boilerplate/modules/roles"
	"go-boilerplate/modules/users"

	"github.com/kataras/iris/v12"
)

type handler struct {
	users    users.Service
	roles    roles.Service
	adapters adapters.Adapters
}

// Register create new user
func (handler handler) Register(ctx iris.Context) {
	request := ctx.Values().Get("body").(*RegisterRequest)

	user, err := handler.users.CreateUser(request.Email, request.Password)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}

	if !config.EMAILACTIVATION() {
		ctx.Values().Set("user", user)
		ctx.Next()
		return
	}
	helper.
		CreateResponse(ctx).
		Ok().
		WithData(map[string]interface{}{"message": fmt.Sprintf("activation email has been sent to %s", request.Email)}).
		JSON()
}

// Login login with email
func (handler handler) Login(ctx iris.Context) {
	request := ctx.Values().Get("body").(*LoginRequest)

	user, err := handler.users.AuthenticateUser(request.Email, request.Password)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			Unauthorized().
			JSON()
		return
	}

	ctx.Values().Set("user", user)
	ctx.Next()
}

func (handler handler) Logout(ctx iris.Context) {
	helper.CreateResponse(ctx).Ok().WithMessage("logout success").JSON()
	ctx.Next()
}
