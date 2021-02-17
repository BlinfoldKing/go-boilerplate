package auth

import (
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

	user, err := handler.auth.Register(
		request.Email,
		request.Password,
		request.CompanyContactID,
		request.RoleIDs,
	)

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
		WithMessage("an account activation email for your request has been sent").
		JSON()
}

// VerifyActivationRequest handles activation verification requests
func (handler handler) VerifyActivationRequest(ctx iris.Context) {
	token := ctx.URLParamDefault("token", "")
	email := ctx.URLParamDefault("email", "")

	err := handler.auth.RequestVerifyActivation(token, email)
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
		WithMessage("account is now activated, you can login now").
		JSON()

	ctx.Next()
}

// ActivateAccountRequest handles account activation requests
func (handler handler) ActivateAccountRequest(ctx iris.Context) {
	request := ctx.Values().Get("body").(*ActivateAccountRequest)

	err := handler.auth.RequestActivateAccount(request.Email)
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

// ResetPassword
func (handler handler) ResetPassword(ctx iris.Context) {
	request := ctx.Values().Get("body").(*ResetPasswordSubmit)
	err := handler.auth.ResetPassword(
		request.Token,
		request.Email,
		request.Password,
	)

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
		WithMessage("password changed").
		JSON()

	ctx.Next()
}

// Login login with email
func (handler handler) Login(ctx iris.Context) {
	request := ctx.Values().Get("body").(*LoginRequest)

	user, err := handler.auth.Login(request.Email, request.Password, request.AsRole)
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
