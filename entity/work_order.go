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
	// InstallationCreated :nodoc
	InstallationCreated StatusType = iota
	// InstallationPreDeliveryCheck :nodoc
	InstallationPreDeliveryCheck
	// InstallationDelivery :nodoc
	InstallationDelivery
	// InstallationDeliveryCheckpoint :nodoc
	InstallationDeliveryCheckpoint
	// InstallationCheckin :nodoc
	InstallationCheckin
	// InstallationRevision :nodoc
	InstallationRevision
	// InstallationInstalling :nodoc
	InstallationInstalling
	// InstallationVerification :nodoc
	InstallationVerification
	// InstallationComplete :nodoc
	InstallationComplete

	// MaintenanceIssued :nodoc
	MaintenanceIssued
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
	// AssestmentCheckin :nodoc
	AssestmentCheckin
	// AssestmentAssestment :nodoc
	AssestmentAssestment
	// AssestmentVerification :nodoc
	AssestmentVerification
	// AssestmentRevision :nodoc
	AssestmentRevision
	// AssestmentComplete :nodoc
	AssestmentComplete

	// AuditIssued :nodoc
	AuditIssued
	// AuditCheckin :nodoc
	AuditCheckin
	// AuditAudit :nodoc
	AuditAudit
	// AuditVerification :nodoc
	AuditVerification
	// AuditComplete :nodoc
	AuditComplete
	// AuditRevision :nodoc
	AuditRevision
)

// WorkOrder work_order entity
type WorkOrder struct {
	ID string `json:"id" xorm:"id"`

	PICID string `json:"pic_id" xorm:"pic_id"`

	Name        string        `json:"name" xorm:"name"`
	Type        WorkOrderType `json:"type" xorm:"type"`
	Status      StatusType    `json:"status" xorm:"status"`
	SiteID      *string       `json:"site_id" xorm:"site_id"`
	NextSiteID  *string       `json:"next_site_id" xorm:"next_site_id"`
	Description string        `json:"description" xorm:"description"`
	NoOrder     string        `json:"no_order" xorm:"no_order"`

	PreviousSiteID *string `json:"previous_site_id" xorm:"previous_site_id"`

	MutationRequestedBy *string    `json:"mutation_requested_by" xorm:"mutation_requested_by"`
	MutationRequestedAt *time.Time `json:"mutation_requested_at" xorm:"mutation_requested_at"`
	MutationApprovedBy  *string    `json:"mutation_approved_by" xorm:"mutation_approved_by"`
	MutationApprovedAt  *time.Time `json:"mutation_approved_at" xorm:"mutation_approved_at"`

	VerifiedBy *string    `json:"verified_by" xorm:"verified_by"`
	VerifiedAt *time.Time `json:"verified_at" xorm:"verified_at"`
	CreatedBy  *string    `json:"created_by" xorm:"created_by"`

	Payload string `json:"payload"`

	CreatedAt time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt *time.Time `json:"deleted_at" xorm:"deleted"`
}

// WorkOrderGroup stores work order group with mapped tables
type WorkOrderGroup struct {
	WorkOrder
	User                    []User       `json:"user"`
	Asset                   []AssetGroup `json:"asset"`
	Document                []Document   `json:"document"`
	Products                []Product    `json:"products"`
	Site                    *Site        `json:"site"`
	MutationApprovedByUser  *User        `json:"mutation_approved_by_user"`
	MutationRequestedByUser *User        `json:"mutation_requested_by_user"`
	VerifiedByUser          *User        `json:"verify_by_user"`
	CreatedByUser           *User        `json:"created_by_user"`
}

// WorkOrderChangeSet change set forwork_order
type WorkOrderChangeSet struct {
	PICID               string        `json:"pic_id" xorm:"pic_id"`
	Name                string        `json:"name" xorm:"name"`
	Type                WorkOrderType `json:"type" xorm:"type"`
	Status              StatusType    `json:"status" xorm:"status"`
	SiteID              *string       `json:"site_id" xorm:"site_id"`
	NextSiteID          *string       `json:"next_site_id" xorm:"next_site_id"`
	Description         string        `json:"description" xorm:"description"`
	NoOrder             string        `json:"no_order" xorm:"no_order"`
	PreviousSiteID      *string       `json:"previous_site_id" xorm:"previous_site_id"`
	MutationRequestedBy *string       `json:"mutation_requested_by" xorm:"mutation_requested_by"`
	MutationRequestedAt *time.Time    `json:"mutation_requested_at" xorm:"mutation_requested_at"`
	MutationApprovedBy  *string       `json:"mutation_approved_by" xorm:"mutation_approved_by"`
	MutationApprovedAt  *time.Time    `json:"mutation_approved_at" xorm:"mutation_approved_at"`
	VerifiedBy          *string       `json:"verified_by" xorm:"verified_by"`
	VerifiedAt          *time.Time    `json:"verified_at" xorm:"verified_at"`
	Payload             string        `json:"payload"`
	CreatedBy           *string       `json:"created_by" xorm:"created_by"`
}

// NewWorkOrder create newwork_order
func NewWorkOrder(noOrder, picid string, siteID *string, name, description string, workOrderType WorkOrderType, status StatusType, payload string, createdBy *string) (workOrder WorkOrder, err error) {
	workOrder = WorkOrder{
		ID:          uuid.NewV4().String(),
		NoOrder:     noOrder,
		PICID:       picid,
		SiteID:      siteID,
		Name:        name,
		Type:        workOrderType,
		Description: description,
		Status:      status,
		Payload:     payload,
		CreatedBy:   createdBy,
	}
	return
}
