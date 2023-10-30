package models

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model
	MemberID      string        `gorm:"primaryKey;not null;uniqueIndex:idx_member_period" json:"member_id"`
	LoanDate      time.Time     `gorm:"not null;uniqueIndex:idx_member_period" json:"loan_date"`
	Amount        int64         `gorm:"not null" json:"amount"`
	CreditMonthly int64         `gorm:"not null" json:"credit_monthly"`
	CreditStart   time.Time     `gorm:"not null;" json:"credit_start"`
	CreditFinish  time.Time     `gorm:"not null;" json:"credit_finish"`
	EntryBy       string        `gorm:"not null" json:"entry_by"`
	UpdateBy      string        `json:"update_by"`
	LoanPayments  []LoanPayment `gorm:"foreignKey:MemberID,LoanDate;references:MemberID,LoanDate"`
}
