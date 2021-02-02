package warehousecontact

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/warehouse:contact"

// Routes init warehouseContact
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePostgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	warehouseContact := prefix.Party(name)
	warehouseContact.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	warehouseContact.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	warehouseContact.Get("/{id:string}", handler.GetByID)
	warehouseContact.Delete("/{id:string}", handler.DeleteByID)
	warehouseContact.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
