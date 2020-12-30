package productcategory

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/product_category"

// Routes init productCategory
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	productCategory := prefix.Party(name)
	productCategory.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	productCategory.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	productCategory.Get("/{id:string}", handler.GetByID)
	productCategory.Delete("/{id:string}", handler.DeleteByID)
	productCategory.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
