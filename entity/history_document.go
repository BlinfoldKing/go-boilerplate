package entity

import (
	"time"

	"github.com/satori/uuid"
)

// HistoryDocument history_document entity
type HistoryDocument struct {
	ID         string     `json:"id" xorm:"id"`
	HistoryID  string     `json:"history_id" xorm:"history_id"`
	DocumentID string     `json:"document_id" xorm:"document_id"`
	CreatedAt  time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt  time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt  *time.Time `json:"deleted_at" xorm:"deleted"`
}

// HistoryDocumentChangeSet change set forhistory_document
type HistoryDocumentChangeSet struct {
	HistoryID  string `json:"history_id" xorm:"history_id"`
	DocumentID string `json:"document_id" xorm:"document_id"`
}

// NewHistoryDocument create newhistory_document
func NewHistoryDocument(historyID, documentID string) (historyDocument HistoryDocument, err error) {
	historyDocument = HistoryDocument{
		ID:         uuid.NewV4().String(),
		HistoryID:  historyID,
		DocumentID: documentID,
	}
	return
}
