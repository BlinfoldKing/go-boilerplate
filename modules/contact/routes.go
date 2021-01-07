package contact

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/contact"

// Routes init contact
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	service := InitContactService(adapters)
	handler := handler{service, adapters}
	contact := prefix.Party(name)
	contact.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	contact.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	contact.Get("/{id:string}", handler.GetByID)
	contact.Delete("/{id:string}", handler.DeleteByID)
	contact.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
