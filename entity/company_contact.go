package entity

import (
	"time"

	"github.com/satori/uuid"
)

// CompanyContact company_contact entity
type CompanyContact struct {
	ID        string     `json:"id" xorm:"id"`
	CompanyID string     `json:"company_id" xorm:"company_id"`
	ContactID string     `json:"contact_id" xorm:"contact_id"`
	CreatedAt time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt *time.Time `json:"deleted_at" xorm:"deleted"`
}

// CompanyContactChangeSet change set forcompany_contact
type CompanyContactChangeSet struct {
	CompanyID string `json:"company_id" xorm:"company_id"`
	ContactID string `json:"contact_id" xorm:"contact_id"`
}

// NewCompanyContact create newcompany_contact
func NewCompanyContact(companyID, contactID string) (companyContact CompanyContact, err error) {
	companyContact = CompanyContact{
		ID:        uuid.NewV4().String(),
		CompanyID: companyID,
		ContactID: contactID,
	}
	return
}
