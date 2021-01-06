package entity

import (
	"time"

	"github.com/satori/uuid"
)

// CompanyType company type
type CompanyType = int

const (
	// VENDORType vendor
	VENDORType CompanyType = 0
	// PRINCIPALType principal
	PRINCIPALType CompanyType = 1
	// PARTNERType partner
	PARTNERType CompanyType = 2
	// INTERNALType internal
	INTERNALType CompanyType = 3
)

// Company company entity
type Company struct {
	ID          string      `json:"id" xorm:"id"`
	Name        string      `json:"name" xorm:"name"`
	Type        CompanyType `json:"type" xorm:"type"`
	Address     string      `json:"address" xorm:"address"`
	PhoneNumber string      `json:"phone_number" xorm:"phone_number"`
	CreatedAt   time.Time   `json:"created_at" xorm:"created"`
	UpdatedAt   time.Time   `json:"updated_at" xorm:"updated"`
	DeletedAt   *time.Time  `json:"deleted_at" xorm:"deleted"`
}

// CompanyGroup company data with mapped tables
type CompanyGroup struct {
	Company
	Documents []Document `json:"documents"`
}

// CompanyChangeSet change set forcompany
type CompanyChangeSet struct {
	Name        string      `json:"name" xorm:"name"`
	Type        CompanyType `json:"type" xorm:"type"`
	Address     string      `json:"address" xorm:"address"`
	PhoneNumber string      `json:"phone_number" xorm:"phone_number"`
}

// NewCompany create newcompany
func NewCompany(
	name string,
	companyType CompanyType,
	address string,
	phoneNumber string,
) (Company, error) {
	company := Company{
		ID:          uuid.NewV4().String(),
		Name:        name,
		Type:        companyType,
		Address:     address,
		PhoneNumber: phoneNumber,
	}
	return company, nil
}
