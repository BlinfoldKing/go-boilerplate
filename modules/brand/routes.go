package brand

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"
	brandcompany "go-boilerplate/modules/brand_company"
	"go-boilerplate/modules/company"

	"github.com/kataras/iris/v12"
)

const name = "/brand"

// Routes init brand
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePostgresRepository(adapters.Postgres)

	brandCompanyRepository := brandcompany.CreatePostgresRepository(adapters.Postgres)
	brandCompanyService := brandcompany.CreateService(brandCompanyRepository)

	companyRepository := company.CreatePostgresRepository(adapters.Postgres)
	companyService := company.CreateService(companyRepository)

	service := CreateService(repository, brandCompanyService, companyService)

	handler := handler{service, adapters}
	brand := prefix.Party(name)
	brand.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	brand.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	brand.Get("/{id:string}", handler.GetByID)
	brand.Delete("/{id:string}", handler.DeleteByID)
	brand.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
