package work_order

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/work_order"

// Routes init work_order
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	service := InitWorkOrderService(adapters)
	handler := handler{service, adapters}
	work_order := prefix.Party(name)
	work_order.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	work_order.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	work_order.Get("/{id:string}", handler.GetByID)
	work_order.Delete("/{id:string}", handler.DeleteByID)
	work_order.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
}
