package domain

import (
	"time"
)

type CashBalance struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Date           time.Time `gorm:"unique;not null" json:"date"`
	OpeningBalance float64   `gorm:"type:numeric(15,2);default:0" json:"opening_balance"`
	TotalIn        float64   `gorm:"type:numeric(15,2);default:0" json:"total_in"`
	TotalOut       float64   `gorm:"type:numeric(15,2);default:0" json:"total_out"`
	ClosingBalance float64   `gorm:"type:numeric(15,2);default:0" json:"closing_balance"`
	CalculatedAt   time.Time `json:"calculated_at"`
}
