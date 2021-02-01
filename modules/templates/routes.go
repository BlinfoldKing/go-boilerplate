package templates

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/templates"

// Routes init templates
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	service := InitTemplateService(adapters)
	handler := handler{service, adapters}
	templates := prefix.Party(name)
	templates.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	templates.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	templates.Get("/{id:string}", handler.GetByID)
	templates.Delete("/{id:string}", handler.DeleteByID)
	templates.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
