package entity

import (
	"github.com/satori/uuid"
)

// Company company entity
type Company struct {
	ID   string `json:"id" xorm:"id"`
	Name string `json:"name" xorm:"name"`
}

// CompanyChangeSet change set forcompany
type CompanyChangeSet struct {
	Name string `json:"name" xorm:"name"`
}

// NewCompany create newcompany
func NewCompany(name string) (company Company, err error) {
	company = Company{uuid.NewV4().String(), name}
	return
}
