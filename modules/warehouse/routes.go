package warehouse

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/warehouse"

// Routes init warehouse
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	warehouse := prefix.Party(name)
	warehouse.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	warehouse.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	warehouse.Get("/{id:string}", handler.GetByID)
	warehouse.Delete("/{id:string}", handler.DeleteByID)
	warehouse.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
