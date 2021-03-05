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

	workOrder.Post("/{id:string}/mutation:request", handler.RequestMutation)
	workOrder.Post("/{id:string}/mutation:approve", handler.ApproveMutation)
	workOrder.Post("/{id:string}/mutation:decline", handler.DeclineMutation)

	workOrder.Post("/{id:string}/mutation:request-v2", middlewares.ValidateBody(&MutationRequest{}), handler.RequestMutationV2)
	workOrder.Post("/{id:string}/mutation:approve-v2", handler.ApproveMutationV2)
	workOrder.Post("/{id:string}/mutation:decline-v2", handler.DeclineMutationV2)

	workOrder.Post("/{id:string}/audit:request", handler.RequestAudit)
	workOrder.Post("/{id:string}/audit:approve", handler.ApproveAudit)
	workOrder.Post("/{id:string}/audit:decline", handler.DeclineAudit)

	workOrder.Post("/{id:string}/assestment:request", handler.RequestAssestment)
	workOrder.Post("/{id:string}/assestment:approve", handler.ApproveAssestment)
	workOrder.Post("/{id:string}/assestment:decline", handler.DeclineAssestment)

	workOrder.Post("/{id:string}/status:verify-maintenance", handler.VerifyMaintenance)
	workOrder.Post("/{id:string}/status:verify-troubleshoot", handler.VerifyTroubleshoot)
	workOrder.Post("/{id:string}/status:verify-install", handler.VerifyInstallation)
	workOrder.Post("/{id:string}/status:verify-assestment", handler.VerifyAssestment)
	workOrder.Post("/{id:string}/status:verify-audit", handler.VerifyAudit)
}
