package history

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/history"

// Routes init history
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePostgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	history := prefix.Party(name)
	history.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	history.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	history.Get("/{id:string}", handler.GetByID)
	history.Delete("/{id:string}", handler.DeleteByID)
	history.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
