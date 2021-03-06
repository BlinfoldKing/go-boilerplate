package modules

import (
	"go-boilerplate/adapters"
	"go-boilerplate/config"
	"go-boilerplate/modules/asset"
	"go-boilerplate/modules/auth"
	"go-boilerplate/modules/company"
	"go-boilerplate/modules/documents"
	"go-boilerplate/modules/ping"
	"go-boilerplate/modules/policy"
	"go-boilerplate/modules/product"
	"go-boilerplate/modules/roles"
	"go-boilerplate/modules/user_roles"
	"go-boilerplate/modules/users"

	"github.com/kataras/iris/v12"
)

// Init init modules
func Init(app *iris.Application, adapters adapters.Adapters) {
	prefix := app.Party(config.PREFIX())

	// init routes
	auth.Routes(prefix, adapters)
	documents.Routes(prefix, adapters)
	users.Routes(prefix, adapters)
	roles.Routes(prefix, adapters)
	policy.Routes(prefix, adapters)
	userroles.Routes(prefix, adapters)
	ping.Routes(prefix, adapters)
	product.Routes(prefix, adapters)
	company.Routes(prefix, adapters)
	asset.Routes(prefix, adapters)

	// init queues
	ping.Queue(adapters)
}
