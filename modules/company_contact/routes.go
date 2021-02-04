package companycontact

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/company:contact"

// Routes init companyContact
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePostgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	companyContact := prefix.Party(name)
	companyContact.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	companyContact.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	companyContact.Get("/{id:string}", handler.GetByID)
	companyContact.Delete("/{id:string}", handler.DeleteByID)
	companyContact.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
