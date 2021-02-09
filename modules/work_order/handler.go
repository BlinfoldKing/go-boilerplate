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

func (h handler) RequestMutationV2(ctx iris.Context) {
	id := ctx.Params().GetString("id")
	user := ctx.Values().Get("user").(entity.UserGroup)
	request := ctx.Values().Get("body").(*MutationRequest)

	wo, err := h.workorders.RequestMutationV2(id, user.ID, request.NextSiteID)

	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().WithData(wo).JSON()
	ctx.Next()
}

func (h handler) RequestMutation(ctx iris.Context) {
	id := ctx.Params().GetString("id")
	user := ctx.Values().Get("user").(entity.UserGroup)
	err := h.workorders.RequestMutation(id, user.ID)

	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().JSON()
	ctx.Next()
}

func (h handler) RequestAssestment(ctx iris.Context) {
	id := ctx.Params().GetString("id")
	workOrder, err := h.workorders.RequestAssestment(id)

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

func (h handler) RequestAudit(ctx iris.Context) {
	id := ctx.Params().GetString("id")
	workOrder, err := h.workorders.RequestAudit(id)

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

func (h handler) ApproveMutationV2(ctx iris.Context) {
	id := ctx.Params().GetString("id")
	user := ctx.Values().Get("user").(entity.UserGroup)
	workOrder, err := h.workorders.ApproveMutationV2(id, user.ID)

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
	id := ctx.Params().GetString("id")
	user := ctx.Values().Get("user").(entity.UserGroup)
	workOrder, err := h.workorders.ApproveMutation(id, user.ID)

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

func (h handler) VerifyInstallation(ctx iris.Context) {
	id := ctx.Params().GetString("id")
	user := ctx.Values().Get("user").(entity.UserGroup)
	workOrder, err := h.workorders.VerifyInstallation(id, user.ID)

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

func (h handler) DeclineMutationV2(ctx iris.Context) {
	id := ctx.Params().GetString("id")
	workOrder, err := h.workorders.DeclineMutationV2(id)

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

func (h handler) ApproveAudit(ctx iris.Context) {
	user := ctx.Values().Get("user").(entity.UserGroup)
	id := ctx.Params().GetString("id")
	workOrder, err := h.workorders.ApproveAudit(id, user.ID)

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

func (h handler) DeclineAudit(ctx iris.Context) {
	user := ctx.Values().Get("user").(entity.UserGroup)
	id := ctx.Params().GetString("id")
	workOrder, err := h.workorders.DeclineAudit(id, user.ID)

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

func (h handler) ApproveAssestment(ctx iris.Context) {
	user := ctx.Values().Get("user").(entity.UserGroup)
	id := ctx.Params().GetString("id")
	workOrder, err := h.workorders.ApproveAssestment(id, user.ID)

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

func (h handler) DeclineAssestment(ctx iris.Context) {
	user := ctx.Values().Get("user").(entity.UserGroup)
	id := ctx.Params().GetString("id")
	workOrder, err := h.workorders.DeclineAssestment(id, user.ID)

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
		request.NoOrder,
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
