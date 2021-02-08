package sitecontact

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/site-contacts"

// Routes init site_contact
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	siteContact := prefix.Party(name)
	siteContact.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	siteContact.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	siteContact.Get("/{id:string}", handler.GetByID)
	siteContact.Delete("/{id:string}", handler.DeleteByID)
	siteContact.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
