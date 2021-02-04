package workorder

import (
	"go-boilerplate/adapters"
	"go-boilerplate/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "/work-order"

// Routes init work_order
func Routes(prefix iris.Party, adapters adapters.Adapters) {
	service := InitWorkOrderService(adapters)
	handler := handler{service, adapters}
	workOrder := prefix.Party(name)
	workOrder.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)
	workOrder.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)
	workOrder.Get("/{id:string}", handler.GetByID)
	workOrder.Delete("/{id:string}", handler.DeleteByID)
	workOrder.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)
	workOrder.Post("/{id:string}/mutation:approve",
		middlewares.ValidateBody(&ApproveRequest{}), handler.ApproveMutation)
	workOrder.Post("/{id:string}/mutation:decline", handler.DeclineMutation)
	workOrder.Post("/{id:string}/asset:approve", handler.ApproveAsset)
	workOrder.Post("/{id:string}/asset:decline", handler.DeclineAsset)
}
