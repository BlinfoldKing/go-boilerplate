package involveduser

import (
	"fmt"
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	"go-boilerplate/helper"

	"github.com/kataras/iris/v12"
)

type handler struct {
	involvedUsers Service
	adapters      adapters.Adapters
}

func (h handler) GetList(ctx iris.Context) {
	request := ctx.Values().Get("pagination").(entity.Pagination)
	involvedUsers, count, err := h.involvedUsers.GetList(request)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreatePaginationResponse(ctx, request, involvedUsers, count).JSON()
	ctx.Next()
}
func (h handler) GetByID(ctx iris.Context) {
	id := ctx.Params().GetString("id")
	involvedUser, err := h.involvedUsers.GetByID(id)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreateResponse(ctx).Ok().WithData(involvedUser).JSON()
	ctx.Next()
}
func (h handler) DeleteByID(ctx iris.Context) {
	id := ctx.Params().GetString("id")
	err := h.involvedUsers.DeleteByID(id)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreateResponse(ctx).Ok().WithMessage(fmt.Sprintf("%s deleted", id)).JSON()
	ctx.Next()
}
func (h handler) Update(ctx iris.Context) {
	request := ctx.Values().Get("body").(*UpdateRequest)
	id := ctx.Params().GetString("id")
	involvedUser, err := h.involvedUsers.Update(id, entity.InvolvedUserChangeSet{
		WorkOrderID: request.WorkOrderID,
		UserID:      request.UserID,
	})
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreateResponse(ctx).Ok().WithData(involvedUser).JSON()
	ctx.Next()
}
func (h handler) Create(ctx iris.Context) {
	request := ctx.Values().Get("body").(*CreateRequest)
	involvedUser, err := h.involvedUsers.CreateInvolvedUser(request.UserID, request.WorkOrderID)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreateResponse(ctx).Ok().WithData(involvedUser).JSON()
	ctx.Next()
}
