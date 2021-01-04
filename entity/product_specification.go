package entity

import (
	"time"

	"github.com/satori/uuid"
)

// ProductSpecification product_specification entity
type ProductSpecification struct {
	ID        string     `json:"id" xorm:"id"`
	ProductID string     `json:"product_id" xorm:"product_id"`
	Parameter string     `json:"parameter" xorm:"parameter"`
	Value     string     `json:"value" xorm:"value"`
	CreatedAt time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt *time.Time `json:"deleted_at" xorm:"deleted"`
}

// ProductSpecificationChangeSet change set forproduct_specification
type ProductSpecificationChangeSet struct {
	Parameter string `json:"parameter" xorm:"parameter"`
	Value     string `json:"value" xorm:"value"`
}

// NewProductSpecification create newproduct_specification
func NewProductSpecification(productID, parameter, value string) (productSpecification ProductSpecification, err error) {
	productSpecification = ProductSpecification{
		ID:        uuid.NewV4().String(),
		ProductID: productID,
		Parameter: parameter,
		Value:     value,
	}
	return
}
