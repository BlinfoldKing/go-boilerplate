package documents

import (
	"go-boilerplate/adapters"
	"go-boilerplate/helper"

	"github.com/kataras/iris/v12"
)

type handler struct {
	documents Service
	adapters  adapters.Adapters
}

func (h handler) Create(ctx iris.Context) {
	request := &UploadRequest{}
	err := ctx.ReadJSON(request)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}

	presignedURL, err := h.adapters.Minio.GeneratePutURL(request.ObjectName, request.BucketName)
	if err != nil {
		helper.
			CreateErrorResponse(ctx, err).
			InternalServer().
			JSON()
		return
	}

	createdDocument, err := h.documents.CreateDocument(request.ObjectName, request.BucketName, presignedURL)
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
