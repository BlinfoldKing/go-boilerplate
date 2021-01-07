package entity

import (
	"time"

	"github.com/satori/uuid"
)

// WorkOrderType work order type
type WorkOrderType = int

const (
	// SiteAcquisitionType site acquisition
	SiteAcquisitionType WorkOrderType = 0
	// CMEType cme
	CMEType WorkOrderType = 1
	// InstallationAndCommissioningType installation and commissioning
	InstallationAndCommissioningType WorkOrderType = 2
	// PreventiveMaintenanceType preventive maintenance type
	PreventiveMaintenanceType WorkOrderType = 3
	// CorrectiveMaintenanceType corrective maintenance type
	CorrectiveMaintenanceType WorkOrderType = 4
	// InternalAuditType internal audit type
	InternalAuditType WorkOrderType = 5
	// ExternalAuditType external external audit type
	ExternalAuditType WorkOrderType = 6
)

// WorkOrder work_order entity
type WorkOrder struct {
	ID          string        `json:"id" xorm:"id"`
	PICID       string        `json:"pic_id" xorm:"pic_id"`
	Name        string        `json:"name" xorm:"name"`
	Type        WorkOrderType `json:"type" xorm:"type"`
	Description string        `json:"description" xorm:"description"`
	CreatedAt   time.Time     `json:"created_at" xorm:"created"`
	UpdatedAt   time.Time     `json:"updated_at" xorm:"updated"`
	DeletedAt   *time.Time    `json:"deleted_at" xorm:"deleted"`
}

// WorkOrderGroup stores work order group with mapped tables
type WorkOrderGroup struct {
	WorkOrder
	User     []User     `json:"user"`
	Asset    []Asset    `json:"asset"`
	Document []Document `json:"document"`
}

// WorkOrderChangeSet change set forwork_order
type WorkOrderChangeSet struct {
	PICID       string        `json:"pic_id" xorm:"pic_id"`
	Name        string        `json:"name" xorm:"name"`
	Type        WorkOrderType `json:"type" xorm:"type"`
	Description string        `json:"description" xorm:"description"`
}

// NewWorkOrder create newwork_order
func NewWorkOrder(picid, name, description string, workOrderType WorkOrderType) (workOrder WorkOrder, err error) {
	workOrder = WorkOrder{
		ID:          uuid.NewV4().String(),
		PICID:       picid,
		Name:        name,
		Type:        workOrderType,
		Description: description,
	}
	return
}
