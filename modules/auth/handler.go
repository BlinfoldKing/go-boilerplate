package auth

import (
	"fmt"
	"go-boilerplate/adapters"
	"go-boilerplate/config"
	"go-boilerplate/helper"

	"github.com/kataras/iris/v12"
)

type handler struct {
	auth     Service
	adapters adapters.Adapters
}

// Register create new user
func (handler handler) Register(ctx iris.Context) {
	request := ctx.Values().Get("body").(*RegisterRequest)

	user, err := handler.auth.Register(request.Email, request.Password)
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

// ResetPasswordRequest
func (handler handler) ResetPasswordRequest(ctx iris.Context) {
	request := ctx.Values().Get("body").(*ResetPasswordRequest)

	err := handler.auth.RequestResetPassword(request.Email)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}

	helper.
		CreateResponse(ctx).
		Ok().
		WithMessage("an email for your request has been sent").
		JSON()

	ctx.Next()
}

// Login login with email
func (handler handler) Login(ctx iris.Context) {
	request := ctx.Values().Get("body").(*LoginRequest)

	user, err := handler.auth.Login(request.Email, request.Password)
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
