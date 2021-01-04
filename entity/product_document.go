package entity

import (
	"time"

	"github.com/satori/uuid"
)

// ProductDocument product_document entity
type ProductDocument struct {
	ID         string     `json:"id" xorm:"id"`
	ProductID  string     `json:"product_id" xorm:"product_id"`
	DocumentID string     `json:"document_id" xorm:"document_id"`
	CreatedAt  time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt  time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt  *time.Time `json:"deleted_at" xorm:"deleted"`
}

// ProductDocumentChangeSet change set forproduct_document
type ProductDocumentChangeSet struct {
	ProductID  string `json:"product_id" xorm:"product_id"`
	DocumentID string `json:"document_id" xorm:"document_id"`
}

// NewProductDocument create newproduct_document
func NewProductDocument(productID, documentID string) (productDocument ProductDocument, err error) {
	productDocument = ProductDocument{
		ID:         uuid.NewV4().String(),
		ProductID:  productID,
		DocumentID: documentID,
	}
	return
}
