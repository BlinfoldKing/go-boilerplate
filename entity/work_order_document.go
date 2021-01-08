package entity

import (
	"time"

	"github.com/satori/uuid"
)

// WorkOrderDocument work_order_document entity
type WorkOrderDocument struct {
	ID          string     `json:"id" xorm:"id"`
	WorkOrderID string     `json:"work_order_id" xorm:"work_order_id"`
	DocumentID  string     `json:"document_id" xorm:"document_id"`
	CreatedAt   time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt   time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt   *time.Time `json:"deleted_at" xorm:"deleted"`
}

// WorkOrderDocumentChangeSet change set forwork_order_document
type WorkOrderDocumentChangeSet struct {
	WorkOrderID string `json:"work_order_id" xorm:"work_order_id"`
	DocumentID  string `json:"document_id" xorm:"document_id"`
}

// NewWorkOrderDocument create newwork_order_document
func NewWorkOrderDocument(workOrderID, documentID string) (workOrderDocument WorkOrderDocument, err error) {
	workOrderDocument = WorkOrderDocument{
		ID:          uuid.NewV4().String(),
		WorkOrderID: workOrderID,
		DocumentID:  documentID,
	}
	return
}
