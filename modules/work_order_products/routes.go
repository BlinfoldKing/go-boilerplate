package workorderproducts

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/work-order-products"

// Routes init workOrderProducts
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	workOrderProducts := prefix.Party(name)
	workOrderProducts.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	workOrderProducts.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	workOrderProducts.Get("/{id:string}", handler.GetByID)
	workOrderProducts.Delete("/{id:string}", handler.DeleteByID)
	workOrderProducts.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
