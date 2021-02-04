package workorder

import (
	"fmt"
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	"go-boilerplate/helper"

	"github.com/kataras/iris/v12"
)

type handler struct {
	workorders Service
	adapters   adapters.Adapters
}

func (h handler) GetList(ctx iris.Context) {
	request := ctx.Values().Get("pagination").(entity.Pagination)
	workOrders, count, err := h.workorders.GetList(request)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreatePaginationResponse(ctx, request, workOrders, count).JSON()
	ctx.Next()
}

func (h handler) GetByID(ctx iris.Context) {
	id := ctx.Params().GetString("id")
	workOrder, err := h.workorders.GetByID(id)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreateResponse(ctx).Ok().WithData(workOrder).JSON()
	ctx.Next()
}

func (h handler) DeleteByID(ctx iris.Context) {
	id := ctx.Params().GetString("id")
	err := h.workorders.DeleteByID(id)
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
	workOrder, err := h.workorders.Update(id, entity.WorkOrderChangeSet{
		PICID:       request.PICID,
		SiteID:      request.SiteID,
		Name:        request.Name,
		Description: request.Description,
		Type:        request.Type,
		Status:      request.Status,
	})
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreateResponse(ctx).Ok().WithData(workOrder).JSON()
	ctx.Next()
}

func (h handler) ApproveMutation(ctx iris.Context) {
	request := ctx.Values().Get("body").(*ApproveRequest)
	id := ctx.Params().GetString("id")
	workOrder, err := h.workorders.ApproveMutation(id, request.SiteID)

	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().WithData(workOrder).JSON()
	ctx.Next()
}

func (h handler) DeclineMutation(ctx iris.Context) {
	id := ctx.Params().GetString("id")
	workOrder, err := h.workorders.DeclineMutation(id)

	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().WithData(workOrder).JSON()
	ctx.Next()
}

func (h handler) ApproveAsset(ctx iris.Context) {
	user := ctx.Values().Get("user").(entity.UserGroup)
	id := ctx.Params().GetString("id")
	workOrder, err := h.workorders.ApproveAsset(id, user.ID)

	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().WithData(workOrder).JSON()
	ctx.Next()
}

func (h handler) DeclineAsset(ctx iris.Context) {
	user := ctx.Values().Get("user").(entity.UserGroup)
	id := ctx.Params().GetString("id")
	workOrder, err := h.workorders.DeclineAsset(id, user.ID)

	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().WithData(workOrder).JSON()
	ctx.Next()
}

func (h handler) Create(ctx iris.Context) {
	request := ctx.Values().Get("body").(*CreateRequest)
	workOrder, err := h.workorders.CreateWorkOrder(
		request.PICID,
		request.SiteID,
		request.Name,
		request.Description,
		request.Type,
		request.InvolvedIDs,
		request.Status,
		request.Assets,
		request.DocumentIDs,
	)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreateResponse(ctx).Ok().WithData(workOrder).JSON()
	ctx.Next()
}
