package userroles

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/user_roles"

// Routes init roles
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}

	roles := prefix.Party(name)

	roles.Post("/", middlewares.ValidateBody(&CreateRequest{}),
		handler.CreateRole)
}
