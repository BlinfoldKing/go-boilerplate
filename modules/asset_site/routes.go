package assetsite

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/asset_site"

// Routes init asset_site
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	assetSite := prefix.Party(name)
	assetSite.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	assetSite.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	assetSite.Get("/{id:string}", handler.GetByID)
	assetSite.Delete("/{id:string}", handler.DeleteByID)
	assetSite.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
