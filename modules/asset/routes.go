package asset

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/asset"

// Routes init asset
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	asset := prefix.Party(name)
	asset.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	asset.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	asset.Get("/{id:string}", handler.GetByID)
	asset.Delete("/{id:string}", handler.DeleteByID)
	asset.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
