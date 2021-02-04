package historydocument

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/history:document"

// Routes init historyDocument
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePostgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	historyDocument := prefix.Party(name)
	historyDocument.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	historyDocument.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	historyDocument.Get("/{id:string}", handler.GetByID)
	historyDocument.Delete("/{id:string}", handler.DeleteByID)
	historyDocument.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
