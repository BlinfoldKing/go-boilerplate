package entity

import (
	"time"

	"github.com/satori/uuid"
)

// BrandCompany brand_company entity
type BrandCompany struct {
	ID        string     `json:"id" xorm:"id"`
	BrandID   string     `json:"brand_id" xorm:"brand_id"`
	CompanyID string     `json:"company_id" xorm:"company_id"`
	CreatedAt time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt *time.Time `json:"deleted_at" xorm:"deleted"`
}

// BrandCompanyChangeSet change set forbrand_company
type BrandCompanyChangeSet struct {
	BrandID   string `json:"brand_id" xorm:"brand_id"`
	CompanyID string `json:"company_id" xorm:"company_id"`
}

// NewBrandCompany create newbrand_company
func NewBrandCompany(brandID, companyID string) (brandCompany BrandCompany, err error) {
	brandCompany = BrandCompany{
		ID:        uuid.NewV4().String(),
		BrandID:   brandID,
		CompanyID: companyID,
	}
	return
}
