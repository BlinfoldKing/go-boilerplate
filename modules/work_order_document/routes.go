package work_order_document

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/work_order_document"

// Routes init work_order_document
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	work_order_document := prefix.Party(name)
	work_order_document.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	work_order_document.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	work_order_document.Get("/{id:string}", handler.GetByID)
	work_order_document.Delete("/{id:string}", handler.DeleteByID)
	work_order_document.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
