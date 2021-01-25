package entity

import (
	"time"

	"github.com/satori/uuid"
)

// SiteAsset asset_site entity
type SiteAsset struct {
	ID        string     `json:"id" xorm:"id"`
	AssetID   string     `json:"asset_id" xorm:"asset_id"`
	SiteID    string     `json:"site_id" xorm:"site_id"`
	CreatedAt time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt *time.Time `json:"deleted_at" xorm:"deleted"`
}

// SiteAssetChangeSet change set forasset_site
type SiteAssetChangeSet struct {
	AssetID   string     `json:"asset_id" xorm:"asset_id"`
	SiteID    string     `json:"site_id" xorm:"site_id"`
}

// NewSiteAsset create newasset_site
func NewSiteAsset(assetID string, siteID string) (assetSite SiteAsset, err error) {
	id := uuid.NewV4().String()
	assetSite = SiteAsset{
		ID:      id,
		AssetID: assetID,
		SiteID:  siteID,
	}
	return
}
