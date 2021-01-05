package users

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/users"

// Routes init users
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	service := InitUserService(adapters)
	handler := handler{service, adapters}

	users := prefix.Party(name)

	users.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	users.Get("/{id:string}", handler.GetByID)
	users.Delete("/{id:string}", handler.DeleteByID)
	users.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}),
		handler.Update)
}
