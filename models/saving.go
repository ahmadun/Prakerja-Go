package models

import (
	"gorm.io/gorm"
)

type Saving struct {
	gorm.Model
	MemberID string `gorm:"primaryKey;not null;uniqueIndex:idx_member_period" json:"member_id"`
	Period   string `gorm:"not null;size:6;uniqueIndex:idx_member_period" json:"period"`
	Amount   int64  `gorm:"not null" json:"amount"`
	EntryBy  string `gorm:"not null" json:"entry_by"`
	UpdateBy string `json:"update_by"`
}
