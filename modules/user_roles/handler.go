package userroles

import (
	"go-boilerplate/adapters"
	"go-boilerplate/helper"

	"github.com/kataras/iris/v12"
)

type handler struct {
	userroles Service
	adapters  adapters.Adapters
}

func (h handler) CreateRole(ctx iris.Context) {
	request := ctx.Values().Get("body").(*CreateRequest)

	userRole, err := h.userroles.CreateRole(request.UserID, request.RoleID)
	if err != nil {
		helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().WithData(userRole).JSON()
	ctx.Next()
}
