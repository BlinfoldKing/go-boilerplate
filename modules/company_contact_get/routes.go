package companycontactget

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/company:contact"

// Routes init companyContact
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	service := InitCompanyContactGetService(adapters)
	handler := handler{service, adapters}
	companyContact := prefix.Party(name)
	companyContact.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	companyContact.Get("/{id:string}", handler.GetByID)
}
