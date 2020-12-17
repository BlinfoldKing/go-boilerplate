package roles

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/roles"

// Routes init roles
func Routes(app *iris.Application, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}

	roles := app.Party(name)

	roles.Post("/", middlewares.ValidateBody(&CreateRequest{}),
		handler.CreateRole)
	roles.Get("/", handler.GetList)
	roles.Get("/{id:string}", handler.GetByID)
	roles.Delete("/{id:string}", handler.DeleteByID)
	roles.Put("/", middlewares.ValidateBody(&UpdateRequest{}),
		handler.Update)
}
