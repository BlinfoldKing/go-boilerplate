package templatesinvolvedid

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/templates-involved-id"

// Routes init templatesInvolvedID
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	templatesInvolvedID := prefix.Party(name)
	templatesInvolvedID.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	templatesInvolvedID.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	templatesInvolvedID.Get("/{id:string}", handler.GetByID)
	templatesInvolvedID.Delete("/{id:string}", handler.DeleteByID)
	templatesInvolvedID.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
