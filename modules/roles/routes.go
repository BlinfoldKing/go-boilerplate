package roles

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/roles"

// Routes init roles
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}

	roles := prefix.Party(name)

	roles.Post("/", middlewares.ValidateBody(&CreateRequest{}),
		handler.CreateRole)
	roles.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	roles.Get("/{id:string}", handler.GetByID)
	roles.Delete("/{id:string}", handler.DeleteByID)
	roles.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}),
		handler.Update)
}
