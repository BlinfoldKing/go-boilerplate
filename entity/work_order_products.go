package entity

import (
	"github.com/satori/uuid"
	"time"
)

// WorkOrderProductStatus status for workorder asset
type WorkOrderProductStatus = int

const (
	// PRODUCTUNCHECKED 0
	PRODUCTUNCHECKED WorkOrderProductStatus = iota
	// PRODUCTCHECKED 1
	PRODUCTCHECKED
	// PRODUCTREVISION 2
	PRODUCTREVISION
)

// WorkOrderProducts workOrderProducts entity
type WorkOrderProducts struct {
	ID          string                 `json:"id" xorm:"id"`
	Status      int                    `json:"status" xorm:"status"`
	WorkOrderID string                 `json:"work_order_id" xorm:"work_order_id"`
	ProductID   string                 `json:"product_id" xorm:"product_id"`
	Qty         WorkOrderProductStatus `json:"qty" xorm:"qty"`
	CreatedAt   time.Time              `json:"created_at" xorm:"created"`
	UpdatedAt   time.Time              `json:"updated_at" xorm:"updated"`
	DeletedAt   *time.Time             `json:"deleted_at" xorm:"deleted"`
}

// WorkOrderProductsChangeSet change set forworkOrderProducts
type WorkOrderProductsChangeSet struct {
	ID          string                 `json:"id" xorm:"id"`
	Status      int                    `json:"status" xorm:"status"`
	WorkOrderID string                 `json:"work_order_id" xorm:"work_order_id"`
	ProductID   string                 `json:"product_id" xorm:"product_id"`
	Qty         WorkOrderProductStatus `json:"qty" xorm:"qty"`
	CreatedAt   time.Time              `json:"created_at" xorm:"created"`
	UpdatedAt   time.Time              `json:"updated_at" xorm:"updated"`
	DeletedAt   *time.Time             `json:"deleted_at" xorm:"deleted"`
}

// NewWorkOrderProducts create newworkOrderProducts
func NewWorkOrderProducts(status int, workOrderid string, productid string, qty int) (workOrderProducts WorkOrderProducts, err error) {
	workOrderProducts = WorkOrderProducts{
		ID:          uuid.NewV4().String(),
		Status:      status,
		WorkOrderID: workOrderid,
		ProductID:   productid,
		Qty:         qty,
	}
	return
}
