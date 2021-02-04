package workorderdocument

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/workorder:document"

// Routes init workorderDocument
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePostgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	workorderDocument := prefix.Party(name)
	workorderDocument.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	workorderDocument.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	workorderDocument.Get("/{id:string}", handler.GetByID)
	workorderDocument.Delete("/{id:string}", handler.DeleteByID)
	workorderDocument.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
