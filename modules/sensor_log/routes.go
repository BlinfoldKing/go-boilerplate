package sensorlog

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/sensor-log"

// Routes init sensorLog
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	repository := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repository)
	handler := handler{service, adapters}
	sensorLog := prefix.Party(name)
	sensorLog.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	sensorLog.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	sensorLog.Get("/{id:string}", handler.GetByID)
	sensorLog.Delete("/{id:string}", handler.DeleteByID)
	sensorLog.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
