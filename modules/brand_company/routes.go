package brandcompany

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/brand:company"

// Routes init brandCompany
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePostgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	brandCompany := prefix.Party(name)
	brandCompany.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	brandCompany.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	brandCompany.Get("/{id:string}", handler.GetByID)
	brandCompany.Delete("/{id:string}", handler.DeleteByID)
	brandCompany.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
