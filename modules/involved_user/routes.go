package involveduser

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/involved-user"

// Routes init involved_user
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	service := InitInvolvedUserService(adapters)
	handler := handler{service, adapters}
	involvedUser := prefix.Party(name)
	involvedUser.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	involvedUser.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	involvedUser.Get("/{id:string}", handler.GetByID)
	involvedUser.Delete("/{id:string}", handler.DeleteByID)
	involvedUser.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
