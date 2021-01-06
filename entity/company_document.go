package entity

import (
	"time"

	"github.com/satori/uuid"
)

// CompanyDocument company_document entity
type CompanyDocument struct {
	ID         string     `json:"id" xorm:"id"`
	CompanyID  string     `json:"company_id" xorm:"company_id"`
	DocumentID string     `json:"document_id" xorm:"document_id"`
	CreatedAt  time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt  time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt  *time.Time `json:"deleted_at" xorm:"deleted"`
}

// CompanyDocumentChangeSet change set forcompany_document
type CompanyDocumentChangeSet struct {
	CompanyID  string `json:"company_id" xorm:"company_id"`
	DocumentID string `json:"document_id" xorm:"document_id"`
}

// NewCompanyDocument create newcompanyDocument
func NewCompanyDocument(companyID, documentID string) (companyDocument CompanyDocument, err error) {
	companyDocument = CompanyDocument{
		ID:         uuid.NewV4().String(),
		CompanyID:  companyID,
		DocumentID: documentID,
	}
	return
}
