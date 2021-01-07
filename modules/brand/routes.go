package brand

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/brand"

// Routes init brand
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	service := InitBrandService(adapters)

	handler := handler{service, adapters}
	brand := prefix.Party(name)
	brand.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	brand.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	brand.Get("/{id:string}", handler.GetByID)
	brand.Delete("/{id:string}", handler.DeleteByID)
	brand.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
