package work_order_asset

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/work_order_asset"

// Routes init work_order_asset
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	work_order_asset := prefix.Party(name)
	work_order_asset.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	work_order_asset.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	work_order_asset.Get("/{id:string}", handler.GetByID)
	work_order_asset.Delete("/{id:string}", handler.DeleteByID)
	work_order_asset.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
