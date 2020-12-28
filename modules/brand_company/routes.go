package brand_company

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/brand_company"

// Routes init brand_company
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	brand_company := prefix.Party(name)
	brand_company.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	brand_company.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	brand_company.Get("/{id:string}", handler.GetByID)
	brand_company.Delete("/{id:string}", handler.DeleteByID)
	brand_company.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
