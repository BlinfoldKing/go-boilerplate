package siteasset

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/site-assets"

// Routes init asset_site
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	siteAsset := prefix.Party(name)
	siteAsset.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	siteAsset.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	siteAsset.Get("/{id:string}", handler.GetByID)
	siteAsset.Delete("/{id:string}", handler.DeleteByID)
	siteAsset.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
