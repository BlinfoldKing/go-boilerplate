package entity

import (
	"time"

	"github.com/satori/uuid"
)

// Brand brand entity
type Brand struct {
	ID            string     `json:"id" xorm:"id"`
	Name          string     `json:"name" xorm:"name"`
	OriginCountry string     `json:"origin_country" xorm:"origin_country"`
	CreatedAt     time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt     time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt     *time.Time `json:"deleted_at" xorm:"deleted"`
}

// BrandGroup user data with role mapped
type BrandGroup struct {
	Brand
	Companies []Company `json:"companies"`
}

// BrandChangeSet change set forbrand
type BrandChangeSet struct {
	Name          string `json:"name" xorm:"name"`
	OriginCountry string `json:"origin_country" xorm:"origin_country"`
}

// NewBrand create newbrand
func NewBrand(name, originCountry string) (brand Brand, err error) {
	brand = Brand{
		ID:            uuid.NewV4().String(),
		Name:          name,
		OriginCountry: originCountry,
	}
	return
}