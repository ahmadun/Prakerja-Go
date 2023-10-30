package models

import (
	"time"

	"gorm.io/gorm"
)

type LoanPayment struct {
	gorm.Model
	MemberID    string    `gorm:"primaryKey;not null;uniqueIndex:idx_member_period" json:"member_id"`
	LoanDate    time.Time `gorm:"not null;" json:"loan_date"`
	Period      string    `gorm:"not null;size:6;uniqueIndex:idx_member_period" json:"period"`
	PaymentTime int       `gorm:"not null;uniqueIndex:idx_member_period" json:"payment_time"`
	Amount      int64     `gorm:"not null" json:"amount"`
	Remark      string    `gorm:"not null" json:"remark"`
	EntryBy     string    `gorm:"not null" json:"entry_by"`
	UpdateBy    string    `json:"update_by"`
}
