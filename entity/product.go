package entity

import (
	"github.com/satori/uuid"
)

// Product product entity
type Product struct {
	ID   string `json:"id" xorm:"id"`
	Name string `json:"name" xorm:"name"`
}

// ProductChangeSet change set forproduct
type ProductChangeSet struct {
	Name string `json:"name" xorm:"name"`
}

// NewProduct create newproduct
func NewProduct(name string) (product Product, err error) {
	product = Product{uuid.NewV4().String(), name}
	return
}
