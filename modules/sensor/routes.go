package sensor

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/sensor"

// Routes init sensor
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	sensor := prefix.Party(name)
	sensor.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	sensor.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	sensor.Get("/{id:string}", handler.GetByID)
	sensor.Delete("/{id:string}", handler.DeleteByID)
	sensor.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
