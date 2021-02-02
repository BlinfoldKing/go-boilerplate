package companydocument

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/company:document"

// Routes init companyDocument
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePostgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	companyDocument := prefix.Party(name)
	companyDocument.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	companyDocument.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	companyDocument.Get("/{id:string}", handler.GetByID)
	companyDocument.Delete("/{id:string}", handler.DeleteByID)
	companyDocument.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
