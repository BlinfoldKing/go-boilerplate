package entity

import (
	"time"

	"github.com/satori/uuid"
)

// SiteDocument site_document entity
type SiteDocument struct {
	ID         string     `json:"id" xorm:"id"`
	DocumentID string     `json:"document_id" xorm:"document_id"`
	SiteID     string     `json:"site_id" xorm:"site_id"`
	CreatedAt  time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt  time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt  *time.Time `json:"deleted_at" xorm:"deleted"`
}

// SiteDocumentChangeSet change set forsite_document
type SiteDocumentChangeSet struct {
	DocumentID string     `json:"document_id" xorm:"document_id"`
	SiteID     string     `json:"site_id" xorm:"site_id"`
}

// NewSiteDocument create newsite_document
func NewSiteDocument(documentID string, siteID string) (siteDocument SiteDocument, err error) {
	id := uuid.NewV4().String()

	siteDocument = SiteDocument{
		ID: id,
		DocumentID: documentID,
		SiteID: siteID,
	}
	return
}
