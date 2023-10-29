package models

import (
	"time"

	"gorm.io/gorm"
)

type Saving struct {
	gorm.Model
	MemberID   string    `gorm:"primaryKey;not null;uniqueIndex:idx_member_period" json:"member_id"`
	Period     string    `gorm:"not null;size:6;uniqueIndex:idx_member_period" json:"period"`
	Amount     int64     `gorm:"not null" json:"amount"`
	EntryBy    string    `gorm:"not null" json:"entry_by"`
	EntryDate  time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"entry_date"`
	UpdateBy   string    `json:"update_by"`
	UpdateDate time.Time `json:"update_date"`
}
