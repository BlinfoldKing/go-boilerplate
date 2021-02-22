package assetwarehouse

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/asset-warehouse"

// Routes init assetWarehouse
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	service := InitService(adapters)

	handler := handler{service, adapters}
	assetWarehouse := prefix.Party(name)
	assetWarehouse.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	assetWarehouse.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	assetWarehouse.Get("/{id:string}", handler.GetByID)
	assetWarehouse.Delete("/{id:string}", handler.DeleteByID)
	assetWarehouse.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
