package documents

import (
	"fmt"
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	"go-boilerplate/helper"

	"github.com/kataras/iris/v12"
)

type handler struct {
	documents Service
	adapters  adapters.Adapters
}

func (h handler) GetList(ctx iris.Context) {
	request := ctx.Values().Get("pagination").(entity.Pagination)
	documents, count, err := h.documents.GetList(request)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreatePaginationResponse(ctx, request, documents, count).JSON()
	ctx.Next()
}

func (h handler) Upload(ctx iris.Context) {
	request := &DocumentUploadRequest{}
	err := ctx.ReadJSON(request)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}

	presignedURL, err := h.documents.UploadDocument(request.ObjectName, request.BucketName)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().WithData(map[string]interface{}{"url": presignedURL}).JSON()
	ctx.Next()
}

func (h handler) Create(ctx iris.Context) {
	request := &entity.Document{}
	err := ctx.ReadJSON(request)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}

	createdDocument, err := h.documents.CreateDocument(*request)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreateResponse(ctx).Ok().WithData(createdDocument).JSON()
	ctx.Next()
}

func (h handler) GetByID(ctx iris.Context) {
	id := ctx.Params().GetString("id")

	document, err := h.documents.GetByID(id)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().WithData(document).JSON()
	ctx.Next()
}

func (h handler) DeleteByID(ctx iris.Context) {
	id := ctx.Params().GetString("id")
	err := h.documents.DeleteByID(id)
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
	document, err := h.documents.Update(id, entity.DocumentChangeSet{
		Name:       request.Name,
		FileSize:   request.FileSize,
		FileType:   request.FileType,
		ObjectName: request.ObjectName,
		BucketName: request.BucketName,
		URLLink:    request.URLLink,
	})
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}
	helper.CreateResponse(ctx).Ok().WithData(document).JSON()
	ctx.Next()
}

func (h handler) Download(ctx iris.Context) {
	request := &DocumentDownloadRequest{}
	err := ctx.ReadJSON(request)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}

	presignedURL, err := h.documents.DownloadDocument(request.ObjectName, request.BucketName)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().WithData(map[string]interface{}{"url": presignedURL}).JSON()
	ctx.Next()
}
