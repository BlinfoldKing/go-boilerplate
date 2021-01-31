package entity

import (
	"time"

	"github.com/satori/uuid"
)

// TemplateItems template_items entity
type TemplateItems struct {
	ID         string     `json:"id" xorm:"id"`
	TemplateID string     `json:"template_id" xorm:"template_id"`
	ProductID  string     `json:"product_id" xorm:"product_id"`
	Qty        int        `json:"qty" xorm:"qty"`
	CreatedAt  time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt  time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt  *time.Time `json:"deleted_at" xorm:"deleted"`
}

// TemplateItemsChangeSet change set fortemplate_items
type TemplateItemsChangeSet struct {
	TemplateID string `json:"template_id" xorm:"template_id"`
	ProductID  string `json:"product_id" xorm:"product_id"`
	Qty        int    `json:"qty" xorm:"qty"`
}

// NewTemplateItems create newtemplate_items
func NewTemplateItems(templateID, productID string, qty int) (templateItems TemplateItems, err error) {
	templateItems = TemplateItems{
		ID:         uuid.NewV4().String(),
		TemplateID: templateID,
		ProductID:  productID,
		Qty:        qty,
	}
	return
}
