package modules

import (
	"go-boilerplate/adapters"
	"go-boilerplate/config"
	"go-boilerplate/modules/asset"
	"go-boilerplate/modules/auth"
	"go-boilerplate/modules/brand"
	brandcompany "go-boilerplate/modules/brand_company"
	"go-boilerplate/modules/company"
	companycontact "go-boilerplate/modules/company_contact"
	companycontactget "go-boilerplate/modules/company_contact_get"
	companydocument "go-boilerplate/modules/company_document"
	"go-boilerplate/modules/contact"
	"go-boilerplate/modules/documents"
	"go-boilerplate/modules/history"
	historydocument "go-boilerplate/modules/history_document"
	involveduser "go-boilerplate/modules/involved_user"
	"go-boilerplate/modules/mail"
	"go-boilerplate/modules/neo4j"
	"go-boilerplate/modules/notifications"
	"go-boilerplate/modules/ping"
	"go-boilerplate/modules/policy"
	"go-boilerplate/modules/product"
	productcategory "go-boilerplate/modules/product_category"
	productdocument "go-boilerplate/modules/product_document"
	productspecification "go-boilerplate/modules/product_specification"
	"go-boilerplate/modules/roles"
	"go-boilerplate/modules/site"
	siteasset "go-boilerplate/modules/site_asset"
	sitecontact "go-boilerplate/modules/site_contact"
	templateitems "go-boilerplate/modules/template_items"
	"go-boilerplate/modules/templates"
	userdevice "go-boilerplate/modules/user_device"
	userroles "go-boilerplate/modules/user_roles"
	"go-boilerplate/modules/users"
	"go-boilerplate/modules/warehouse"
	warehousecontact "go-boilerplate/modules/warehouse_contact"
	workorder "go-boilerplate/modules/work_order"
	workorderasset "go-boilerplate/modules/work_order_asset"
	workorderdocument "go-boilerplate/modules/work_order_document"

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
	contact.Routes(prefix, adapters)
	workorder.Routes(prefix, adapters)
	involveduser.Routes(prefix, adapters)
	warehouse.Routes(prefix, adapters)
	templates.Routes(prefix, adapters)
	templateitems.Routes(prefix, adapters)
	brandcompany.Routes(prefix, adapters)
	companycontact.Routes(prefix, adapters)
	companydocument.Routes(prefix, adapters)
	historydocument.Routes(prefix, adapters)
	productdocument.Routes(prefix, adapters)
	warehousecontact.Routes(prefix, adapters)
	workorderasset.Routes(prefix, adapters)
	workorderdocument.Routes(prefix, adapters)
	site.Routes(prefix, adapters)
	siteasset.Routes(prefix, adapters)
	sitecontact.Routes(prefix, adapters)
	userdevice.Routes(prefix, adapters)
	companycontactget.Routes(prefix, adapters)

	// init queues
	ping.Queue(adapters)
	mail.Queue(adapters)
	notifications.Queue(adapters)
}
