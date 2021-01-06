package work_order_asset

import (
	"fmt"
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	"go-boilerplate/helper"

	"github.com/kataras/iris/v12"
)

type handler struct {
	work_order_assets Service
	adapters          adapters.Adapters
}

func (h handler) GetList(ctx iris.Context) {
	request := ctx.Values().Get("pagination").(entity.Pagination)
	work_order_assets, count, err := h.work_order_assets.GetList(request)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreatePaginationResponse(ctx, request, work_order_assets, count).JSON()
	ctx.Next()
}
func (h handler) GetByID(ctx iris.Context) {
	id := ctx.Params().GetString("id")
	work_order_asset, err := h.work_order_assets.GetByID(id)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreateResponse(ctx).Ok().WithData(work_order_asset).JSON()
	ctx.Next()
}
func (h handler) DeleteByID(ctx iris.Context) {
	id := ctx.Params().GetString("id")
	err := h.work_order_assets.DeleteByID(id)
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
	work_order_asset, err := h.work_order_assets.Update(id, entity.WorkOrderAssetChangeSet{
		Name: request.Name,
	})
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreateResponse(ctx).Ok().WithData(work_order_asset).JSON()
	ctx.Next()
}
func (h handler) Create(ctx iris.Context) {
	request := ctx.Values().Get("body").(*CreateRequest)
	work_order_asset, err := h.work_order_assets.CreateWorkOrderAsset(request.Name)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreateResponse(ctx).Ok().WithData(work_order_asset).JSON()
	ctx.Next()
}
