package sitedocument

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/siteDocument"

// Routes init siteDocument
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	siteDocument := prefix.Party(name)
	siteDocument.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	siteDocument.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	siteDocument.Get("/{id:string}", handler.GetByID)
	siteDocument.Delete("/{id:string}", handler.DeleteByID)
	siteDocument.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
