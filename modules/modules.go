package modules

import (
	"go-boilerplate/adapters"
	"go-boilerplate/config"
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
	prefix := app.Party(config.PREFIX())

	auth.Routes(prefix, adapters)
	documents.Routes(prefix, adapters)
	users.Routes(prefix, adapters)
	roles.Routes(prefix, adapters)
	policy.Routes(prefix, adapters)
	userroles.Routes(prefix, adapters)
	ping.Routes(prefix, adapters)
}
