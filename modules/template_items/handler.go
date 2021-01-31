package templateitems

import (
	"fmt"
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	"go-boilerplate/helper"

	"github.com/kataras/iris/v12"
)

type handler struct {
	templateItems Service
	adapters      adapters.Adapters
}

func (h handler) GetList(ctx iris.Context) {
	request := ctx.Values().Get("pagination").(entity.Pagination)
	templateItems, count, err := h.templateItems.GetList(request)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreatePaginationResponse(ctx, request, templateItems, count).JSON()
	ctx.Next()
}
func (h handler) GetByID(ctx iris.Context) {
	id := ctx.Params().GetString("id")
	templateItems, err := h.templateItems.GetByID(id)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreateResponse(ctx).Ok().WithData(templateItems).JSON()
	ctx.Next()
}
func (h handler) DeleteByID(ctx iris.Context) {
	id := ctx.Params().GetString("id")
	err := h.templateItems.DeleteByID(id)
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
	templateItems, err := h.templateItems.Update(id, entity.TemplateItemsChangeSet{
		TemplateID: request.TemplateID,
		ProductID:  request.ProductID,
		Qty:        request.Qty,
	})
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreateResponse(ctx).Ok().WithData(templateItems).JSON()
	ctx.Next()
}
func (h handler) Create(ctx iris.Context) {
	request := ctx.Values().Get("body").(*CreateRequest)
	templateItems, err := h.templateItems.CreateTemplateItems(request.TemplateID, request.ProductID, request.Qty)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreateResponse(ctx).Ok().WithData(templateItems).JSON()
	ctx.Next()
}
