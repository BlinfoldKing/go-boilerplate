package auth

import (
	"go-boilerplate/adapters"
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
		logrus.Error(err)
		ctx.Next()
		return
	}

	user, err := handler.users.CreateUser(request.Email, request.Password)
	if err != nil {
		ctx.JSON(err)
		ctx.Next()
		return
	}

	ctx.JSON(user)

	ctx.Next()
}
