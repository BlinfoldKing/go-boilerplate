package involved_user

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/involved_user"

// Routes init involved_user
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	involved_user := prefix.Party(name)
	involved_user.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	involved_user.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	involved_user.Get("/{id:string}", handler.GetByID)
	involved_user.Delete("/{id:string}", handler.DeleteByID)
	involved_user.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
