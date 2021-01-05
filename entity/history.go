package entity

import (
	"time"

	"github.com/satori/uuid"
)

// History history entity
type History struct {
	ID          string     `json:"id" xorm:"id"`
	UserID      string     `json:"user_id" xorm:"user_id"`
	AssetID     string     `json:"asset_id" xorm:"asset_id"`
	Action      string     `json:"action" xorm:"action"`
	Description string     `json:"description" xorm:"description"`
	Cost        float64    `json:"cost" xorm:"cost"`
	CreatedAt   time.Time  `json:"created_at" xorm:"created"`
	UpdatedAt   time.Time  `json:"updated_at" xorm:"updated"`
	DeletedAt   *time.Time `json:"deleted_at" xorm:"deleted"`
}

// HistoryGroup history data with mapped tables
type HistoryGroup struct {
	History
	User      User       `json:"user"`
	Asset     Asset      `json:"asset"`
	Documents []Document `json:"documents"`
}

// HistoryChangeSet change set forhistory
type HistoryChangeSet struct {
	UserID      string  `json:"user_id" xorm:"user_id"`
	AssetID     string  `json:"asset_id" xorm:"asset_id"`
	Action      string  `json:"action" xorm:"action"`
	Description string  `json:"description" xorm:"description"`
	Cost        float64 `json:"cost" xorm:"cost"`
}

// NewHistory create newhistory
func NewHistory(userID, assetID, action, description string, cost float64) (history History, err error) {
	history = History{
		ID:          uuid.NewV4().String(),
		UserID:      userID,
		AssetID:     assetID,
		Action:      action,
		Description: description,
		Cost:        cost,
	}
	return
}
