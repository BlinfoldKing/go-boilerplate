package entity

import (
	"time"

	"github.com/satori/uuid"
)

// WorkOrderType work order type
type WorkOrderType = int

// StatusType status type
type StatusType = int

const (
	// Installation :nodoc
	Installation WorkOrderType = iota
	// Maintenance cme
	Maintenance
	// Troubleshooting :nodoc
	Troubleshooting
	// Assesment :nodoc
	Assesment
	// Audit :nodoc
	Audit
)

const (
	// InstallationDelivery :nodoc
	InstallationDelivery StatusType = iota
	// InstallationCheckpoint :nodoc
	InstallationCheckpoint
	// InstallationCheckin :nodoc
	InstallationCheckin
	// InstallationInstalling :nodoc
	InstallationInstalling
	// InstallationVerification :nodoc
	InstallationVerification
	// InstallationComplete :nodoc
	InstallationComplete

	// MaintenanceIssued :nodoc
	MaintenanceIssued
	// MaintenanceCheckpoint :nodoc
	MaintenanceCheckpoint
	// MaintenanceCheckin :nodoc
	MaintenanceCheckin
	// MaintenanceMaintenance :nodoc
	MaintenanceMaintenance
	// MaintenanceVerification :nodoc
	MaintenanceVerification
	// MaintenanceComplete :nodoc
	MaintenanceComplete

	// TroubleshootingIssued :nodoc
	TroubleshootingIssued
	// TroubleshootingCheckpoint :nodoc
	TroubleshootingCheckpoint
	// TroubleshootingCheckin :nodoc
	TroubleshootingCheckin
	// TroubleshootingTroubleshooting :nodoc
	TroubleshootingTroubleshooting
	// TroubleshootingVerification :nodoc
	TroubleshootingVerification
	// TroubleshootingComplete :nodoc
	TroubleshootingComplete

	// AssestmentIssued :nodoc
	AssestmentIssued
	// AssestmentCheckpoint :nodoc
	AssestmentCheckpoint
	// AssestmentCheckin :nodoc
	AssestmentCheckin
	// AssestmentAssestment :nodoc
	AssestmentAssestment
	// AssestmentVerification :nodoc
	AssestmentVerification
	// AssestmentComplete :nodoc
	AssestmentComplete

	// AuditIssued :nodoc
	AuditIssued
	// AuditCheckpoint :nodoc
	AuditCheckpoint
	// AuditCheckin :nodoc
	AuditCheckin
	// AuditAudit :nodoc
	AuditAudit
	// AuditVerification :nodoc
	AuditVerification
	// AuditComplete :nodoc
	AuditComplete
)

// WorkOrder work_order entity
type WorkOrder struct {
	ID          string        `json:"id" xorm:"id"`
	PICID       string        `json:"pic_id" xorm:"pic_id"`
	Name        string        `json:"name" xorm:"name"`
	Type        WorkOrderType `json:"type" xorm:"type"`
	Status      StatusType    `json:"status" xorm:"status"`
	SiteID      *string       `json:"site_id" xorm:"site_id"`
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
	Site     Site       `json:"site"`
}

// WorkOrderChangeSet change set forwork_order
type WorkOrderChangeSet struct {
	PICID       string        `json:"pic_id" xorm:"pic_id"`
	Name        string        `json:"name" xorm:"name"`
	Type        WorkOrderType `json:"type" xorm:"type"`
	Status      StatusType    `json:"status" xorm:"status"`
	SiteID      string        `json:"site_id" xorm:"site_id"`
	Description string        `json:"description" xorm:"description"`
}

// NewWorkOrder create newwork_order
func NewWorkOrder(picid, siteID, name, description string, workOrderType WorkOrderType, status StatusType) (workOrder WorkOrder, err error) {
	workOrder = WorkOrder{
		ID:          uuid.NewV4().String(),
		PICID:       picid,
		SiteID:      siteID,
		Name:        name,
		Type:        workOrderType,
		Description: description,
		Status:      status,
	}
	return
}
