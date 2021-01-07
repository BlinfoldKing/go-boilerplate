package company

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/company"

// Routes init company
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	service := InitCompanyService(adapters)
	handler := handler{service, adapters}
	company := prefix.Party(name)
	company.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	company.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	company.Get("/{id:string}", handler.GetByID)
	company.Delete("/{id:string}", handler.DeleteByID)
	company.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
