package notifications

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/notifications"

// Routes init notification
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	notification := prefix.Party(name)
	notification.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	notification.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	notification.Get("/{id:string}", handler.GetByID)
	notification.Delete("/{id:string}", handler.DeleteByID)
	notification.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
