package auth

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	"go-boilerplate/helper"
	"go-boilerplate/modules/roles"
	"go-boilerplate/modules/users"

	"github.com/fatih/structs"
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

	roles := make([]entity.Role, 0)
	for _, id := range user.Roles {
		role, err := handler.roles.GetByID(id)
		if err != nil {
			helper.
				CreateErrorResponse(ctx, err).
				InternalServer().
				JSON()
			return
		}

		roles = append(roles, role)
	}

	userWithRole := structs.Map(user)
	userWithRole["roles"] = roles

	ctx.Values().Set("user", userWithRole)
	ctx.Next()
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

	roles := make([]entity.Role, 0)
	for _, id := range user.Roles {
		role, err := handler.roles.GetByID(id)
		if err != nil {
			helper.
				CreateErrorResponse(ctx, err).
				InternalServer().
				JSON()
			return
		}

		roles = append(roles, role)
	}

	userWithRole := structs.Map(user)
	userWithRole["roles"] = roles

	ctx.Values().Set("user", userWithRole)
	ctx.Next()
}

func (handler handler) Logout(ctx iris.Context) {
	helper.CreateResponse(ctx).Ok().WithMessage("logout success").JSON()
	ctx.Next()
}
