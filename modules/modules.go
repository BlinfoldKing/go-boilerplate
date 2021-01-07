package modules

import (
	"go-boilerplate/adapters"
	"go-boilerplate/config"
	"go-boilerplate/modules/asset"
	"go-boilerplate/modules/auth"
	"go-boilerplate/modules/brand"
	"go-boilerplate/modules/company"
	"go-boilerplate/modules/documents"
	"go-boilerplate/modules/history"
	involveduser "go-boilerplate/modules/involved_user"
	"go-boilerplate/modules/mail"
	"go-boilerplate/modules/neo4j"
	"go-boilerplate/modules/notifications"
	"go-boilerplate/modules/ping"
	"go-boilerplate/modules/policy"
	"go-boilerplate/modules/product"
	productcategory "go-boilerplate/modules/product_category"
	productspecification "go-boilerplate/modules/product_specification"
	"go-boilerplate/modules/roles"
	userroles "go-boilerplate/modules/user_roles"
	"go-boilerplate/modules/users"
	workorder "go-boilerplate/modules/work_order"

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
	neo4j.Routes(prefix, adapters)
	notifications.Routes(prefix, adapters)
	brand.Routes(prefix, adapters)
	productcategory.Routes(prefix, adapters)
	productspecification.Routes(prefix, adapters)
	history.Routes(prefix, adapters)
	workorder.Routes(prefix, adapters)
	involveduser.Routes(prefix, adapters)

	// init queues
	ping.Queue(adapters)
	mail.Queue(adapters)
}
