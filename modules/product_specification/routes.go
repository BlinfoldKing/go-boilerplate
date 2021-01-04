package productspecification

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/product_specification"

// Routes init product_specification
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	productSpecification := prefix.Party(name)
	productSpecification.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	productSpecification.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	productSpecification.Get("/{id:string}", handler.GetByID)
	productSpecification.Delete("/{id:string}", handler.DeleteByID)
	productSpecification.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
