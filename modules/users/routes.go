package users

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"
	"go-boilerplate/modules/mail"
	"go-boilerplate/modules/otps"
	"go-boilerplate/modules/roles"
	userroles "go-boilerplate/modules/user_roles"

	"github.com/kataras/iris/v12"
)

const name = "/users"

// Routes init users
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	roleRepository := roles.CreatePosgresRepository(adapters.Postgres)
	roleService := roles.CreateService(roleRepository)

	userRoleRepository := userroles.CreatePosgresRepository(adapters.Postgres)
	userRoleService := userroles.CreateService(userRoleRepository)

	otpsRepository := otps.CreatePostgresRepository(adapters.Postgres)
	otpsService := otps.CreateService(otpsRepository)

	mailService := mail.CreateMailgunService(adapters.Mailgun)

	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository, roleService, userRoleService, otpsService, mailService)
	handler := handler{service, adapters}

	users := prefix.Party(name)

	users.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	users.Get("/{id:string}", handler.GetByID)
	users.Delete("/{id:string}", handler.DeleteByID)
	users.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}),
		handler.Update)
}
