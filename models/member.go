package models

import (
	"time"
)

type Member struct {
	Id         string    `gorm:"primaryKey" json:"id" form:"id"`
	Name       string    `gorm:"not null" json:"title"`
	BirthDay   time.Time `gorm:"not null" json:"birth_day"`
	PlaceBirth string    `gorm:"not null" json:"place_birth"`
	Gender     string    `gorm:"not null" json:"gender"`
	JoinDate   time.Time `gorm:"not null" json:"join_date"`
	Active     bool      `gorm:"not null;default:true" json:"active"`
	EntryBy    string    `gorm:"not null" json:"entry_by"`
	EntryDate  time.Time `gorm:"not null" json:"entry_date"`
	UpdateBy   string    `json:"update_by"`
	UpdateDate time.Time `json:"update_date"`
	Savings    []Saving  `gorm:"foreignKey:MemberID;references:Id"`
	Loans      []Loan    `gorm:"foreignKey:MemberID;references:Id"`
}
