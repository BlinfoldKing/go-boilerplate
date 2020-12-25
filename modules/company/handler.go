package company

import (
	"fmt"
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	"go-boilerplate/helper"

	"github.com/kataras/iris/v12"
)

type handler struct {
	companys Service
	adapters adapters.Adapters
}

func (h handler) GetList(ctx iris.Context) {
	request := ctx.Values().Get("pagination").(entity.Pagination)
	companys, count, err := h.companys.GetList(request)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreatePaginationResponse(ctx, request, companys, count).JSON()
	ctx.Next()
}
func (h handler) GetByID(ctx iris.Context) {
	id := ctx.Params().GetString("id")
	company, err := h.companys.GetByID(id)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreateResponse(ctx).Ok().WithData(company).JSON()
	ctx.Next()
}
func (h handler) DeleteByID(ctx iris.Context) {
	id := ctx.Params().GetString("id")
	err := h.companys.DeleteByID(id)
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
	company, err := h.companys.Update(id, entity.CompanyChangeSet{
		Name: request.Name,
	})
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreateResponse(ctx).Ok().WithData(company).JSON()
	ctx.Next()
}
func (h handler) Create(ctx iris.Context) {
	request := ctx.Values().Get("body").(*CreateRequest)
	company, err := h.companys.CreateCompany(request.Name)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreateResponse(ctx).Ok().WithData(company).JSON()
	ctx.Next()
}
