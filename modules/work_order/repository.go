package workorder

import (
	"go-boilerplate/entity"
)

// Repository abstraction for storage
type Repository interface {
	Save(entity.WorkOrder) error
	DeleteByID(id string) error
	FindByID(id string) (entity.WorkOrder, error)
	Update(id string, changeset entity.WorkOrderChangeSet) error
	GetList(pagination entity.Pagination) (WorkOrders []entity.WorkOrder, count int, err error)
	Create(
		wo entity.WorkOrder,
		involvedIDs, documentIDs *[]string,
		assets *[]struct {
			ID  string `json:"id" validate:"required"`
			Qty int    `json:"qty" validate:"required"`
		},
		products *[]struct {
			ID  string `json:"id" validate:"required"`
			Qty int    `json:"qty" validate:"required"`
		},
	) error
	ApproveMutationV2(wo entity.WorkOrderGroup, userid string) error
	ApproveAssestment(wo entity.WorkOrderGroup, userid string) error
	ApproveAudit(wo entity.WorkOrderGroup, userid string) error
}
