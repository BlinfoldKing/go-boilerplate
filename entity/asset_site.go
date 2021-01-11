package entity

import (
	"time"

	"github.com/satori/uuid"
)

// AssetSite asset_site entity
type AssetSite struct {
	ID        string     `json:"id" xorm:"id"`
	AssetID   string     `json:"asset_id" xorm:"asset_id"`
	SiteID    string     `json:"site_id" xorm:"site_id"`
	CreatedAt time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt *time.Time `json:"deleted_at" xorm:"deleted"`
}

// AssetSiteChangeSet change set forasset_site
type AssetSiteChangeSet struct {
	AssetID   string     `json:"asset_id" xorm:"asset_id"`
	SiteID    string     `json:"site_id" xorm:"site_id"`
}

// NewAssetSite create newasset_site
func NewAssetSite(assetID string, siteID string) (assetSite AssetSite, err error) {
	id := uuid.NewV4().String()
	assetSite = AssetSite{
		ID:      id,
		AssetID: assetID,
		SiteID:  siteID,
	}
	return
}
