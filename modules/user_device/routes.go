package userdevice

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/user_device"

// Routes init user_device
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	userdevice := prefix.Party(name)
	userdevice.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	userdevice.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	userdevice.Get("/{id:string}", handler.GetByID)
	userdevice.Delete("/{id:string}", handler.DeleteByID)
	userdevice.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
