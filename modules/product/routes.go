package product

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/product"

// Routes init product
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	service := InitProductService(adapters)
	handler := handler{service, adapters}
	product := prefix.Party(name)
	product.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	product.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	product.Get("/{id:string}", handler.GetByID)
	product.Delete("/{id:string}", handler.DeleteByID)
	product.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
