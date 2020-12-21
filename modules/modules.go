package modules

import (
	"go-boilerplate/adapters"
	"go-boilerplate/modules/auth"
	"go-boilerplate/modules/documents"
	"go-boilerplate/modules/ping"
	"go-boilerplate/modules/policy"
	"go-boilerplate/modules/roles"
	"go-boilerplate/modules/user_roles"
	"go-boilerplate/modules/users"

	"github.com/kataras/iris/v12"
)

// Init init modules
func Init(app *iris.Application, adapters adapters.Adapters) {
	auth.Routes(app, adapters)
	documents.Routes(app, adapters)
	users.Routes(app, adapters)
	roles.Routes(app, adapters)
	policy.Routes(app, adapters)
	userroles.Routes(app, adapters)
	ping.Routes(app, adapters)
}
