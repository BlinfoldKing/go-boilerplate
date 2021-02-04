package workorderasset

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/workorder:asset"

// Routes init workorderAsset
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePostgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	workorderAsset := prefix.Party(name)
	workorderAsset.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	workorderAsset.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	workorderAsset.Get("/{id:string}", handler.GetByID)
	workorderAsset.Delete("/{id:string}", handler.DeleteByID)
	workorderAsset.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
