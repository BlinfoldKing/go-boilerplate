package users

import (
	"fmt"
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	"go-boilerplate/helper"

	"github.com/kataras/iris/v12"
)

type handler struct {
	users    Service
	adapters adapters.Adapters
}

func (h handler) GetList(ctx iris.Context) {
	limit := ctx.URLParamIntDefault("limit", 10)
	offset := ctx.URLParamIntDefault("offset", 10)

	users, err := h.users.GetList(limit, offset)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().WithData(users).JSON()
	ctx.Next()
}

func (h handler) GetByID(ctx iris.Context) {
	id := ctx.Params().GetString("id")

	users, err := h.users.GetByID(id)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().WithData(users).JSON()
	ctx.Next()
}

func (h handler) DeleteByID(ctx iris.Context) {
	id := ctx.Params().GetString("id")

	err := h.users.DeleteByID(id)
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
		WithMessage(fmt.Sprintf("%s deleted", id)).
		JSON()
	ctx.Next()
}

func (h handler) Update(ctx iris.Context) {
	request := ctx.Values().Get("body").(*UpdateRequest)

	user, err := h.users.Update(request.ID, entity.UserChangeSet{
		Email: request.Email,
	})
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
		WithData(user).
		JSON()
	ctx.Next()
}
