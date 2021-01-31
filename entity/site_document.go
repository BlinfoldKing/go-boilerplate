package entity

import (
	"time"

	"github.com/satori/uuid"
)

// SiteDocument site_document entity
type SiteDocument struct {
	ID            string     `json:"id" xorm:"id"`
	DocumentID    string     `json:"document_id" xorm:"document_id"`
	SiteID        string     `json:"site_id" xorm:"site_id"`
	ApproveStatus int        `json:"approve_status" xorm:"approve_status"`
	Notes         string     `json:"notes" xorm:"notes"`
	CreatedAt     time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt     time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt     *time.Time `json:"deleted_at" xorm:"deleted"`
}

// SiteDocumentChangeSet change set forsite_document
type SiteDocumentChangeSet struct {
	DocumentID    string `json:"document_id" xorm:"document_id"`
	SiteID        string `json:"site_id" xorm:"site_id"`
	ApproveStatus int    `json:"approve_status" xorm:"approve_status"`
	Notes         string `json:"notes" xorm:"notes"`
}

// NewSiteDocument create newsite_document
func NewSiteDocument(
	documentID string,
	siteID string,
	approveStatus int,
	notes string,
) SiteDocument {
	site := SiteDocument{
		ID:            uuid.NewV4().String(),
		DocumentID:    documentID,
		ApproveStatus: approveStatus,
		Notes:         notes,
	}
	return site
}
