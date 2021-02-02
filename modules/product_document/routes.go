package productdocument

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/product:document"

// Routes init productDocument
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePostgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	productDocument := prefix.Party(name)
	productDocument.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	productDocument.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	productDocument.Get("/{id:string}", handler.GetByID)
	productDocument.Delete("/{id:string}", handler.DeleteByID)
	productDocument.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
