package site

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/site"

// Routes init site
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	site := prefix.Party(name)
	
	site.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	site.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	site.Get("/{id:string}", handler.GetByID)
	site.Delete("/{id:string}", handler.DeleteByID)
	site.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
