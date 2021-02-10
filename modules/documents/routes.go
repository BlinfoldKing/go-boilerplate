package documents

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/documents"

// Routes init documents
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	service := InitDocumentsService(adapters)
	handler := handler{service, adapters}

	documents := prefix.Party(name)

	documents.Post(":upload", handler.Upload)
	documents.Post(":download", handler.Download)

	documents.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	documents.Get("/{id:string}", handler.GetByID)
	documents.Post("/", handler.Create)
	documents.Delete("/{id:string}", handler.DeleteByID)
	documents.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
