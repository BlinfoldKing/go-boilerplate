package auth

import (
	"go-boilerplate/adapters"
	"go-boilerplate/helper"
	"go-boilerplate/modules/users"

	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

type handler struct {
	users    users.Service
	adapters adapters.Adapters
}

// Register create new user
func (handler handler) Register(ctx iris.Context) {
	var request RegisterRequest
	err := ctx.ReadJSON(&request)
	if err != nil {
		logrus.Error(err)
		return
	}

	err = handler.adapters.Validator.Struct(&request)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			BadRequest().
			JSON()
		return
	}

	user, err := handler.users.CreateUser(request.Email, request.Password)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}

	ctx.Values().Set("user", user)
	ctx.Next()
}
