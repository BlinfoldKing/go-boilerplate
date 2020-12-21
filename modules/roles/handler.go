package roles

import (
	"fmt"
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	"go-boilerplate/helper"

	"github.com/kataras/iris/v12"
)

type handler struct {
	roles    Service
	adapters adapters.Adapters
}

func (h handler) GetList(ctx iris.Context) {
	request := ctx.Values().Get("pagination").(*entity.Pagination)

	roles, err := h.roles.GetList(*request)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().WithData(roles).JSON()
	ctx.Next()
}

func (h handler) GetByID(ctx iris.Context) {
	id := ctx.Params().GetString("id")

	roles, err := h.roles.GetByID(id)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().WithData(roles).JSON()
	ctx.Next()
}

func (h handler) DeleteByID(ctx iris.Context) {
	id := ctx.Params().GetString("id")

	err := h.roles.DeleteByID(id)
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

	role, err := h.roles.Update(request.ID, entity.RoleChangeSet{
		Slug:        request.Slug,
		Description: request.Description,
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
		WithData(role).
		JSON()
	ctx.Next()
}

func (h handler) CreateRole(ctx iris.Context) {
	request := ctx.Values().Get("body").(*CreateRequest)

	role, err := h.roles.CreateRole(request.Slug, request.Description)

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
		WithData(role).
		JSON()
	ctx.Next()
}
