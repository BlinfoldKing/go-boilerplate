package templateitems

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/template:items"

// Routes init templateItems
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	templateItems := prefix.Party(name)
	templateItems.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	templateItems.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	templateItems.Get("/{id:string}", handler.GetByID)
	templateItems.Delete("/{id:string}", handler.DeleteByID)
	templateItems.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
