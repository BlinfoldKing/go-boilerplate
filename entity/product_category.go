package entity

import (
	"time"

	"github.com/satori/uuid"
)

// ProductCategory product_category entity
type ProductCategory struct {
	ID        string     `json:"id" xorm:"id"`
	ParentID  string     `json:"parent_id" xorm:"parent_id"`
	Code      string     `json:"code" xorm:"code"`
	Name      string     `json:"name" xorm:"name"`
	CreatedAt time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt *time.Time `json:"deleted_at" xorm:"deleted"`
}

// ProductCategoryChangeSet change set forproduct_category
type ProductCategoryChangeSet struct {
	Code string `json:"code" xorm:"code"`
	Name string `json:"name" xorm:"name"`
}

// NewProductCategory create newproduct_category
func NewProductCategory(parentID, code, name string) (productCategory ProductCategory, err error) {
	productCategory = ProductCategory{
		ID:       uuid.NewV4().String(),
		ParentID: parentID,
		Code:     code,
		Name:     name,
	}
	return
}
