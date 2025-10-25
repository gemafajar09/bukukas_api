package domain

import (
	"time"
)

type CashTransaction struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	TransactionDate time.Time `gorm:"not null" json:"transaction_date"`
	Type            string    `gorm:"size:10;not null;check:type IN ('in','out')" json:"type"`
	CategoryID      uint      `json:"category_id"`
	Description     string    `gorm:"type:text" json:"description"`
	Amount          float64   `gorm:"type:numeric(15,2);not null" json:"amount"`
	PaymentMethod   string    `gorm:"size:20;default:'cash'" json:"payment_method"`
	ReferenceID     *uint     `json:"reference_id,omitempty"`
	CreatedBy       uint      `json:"created_by"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`

	Category *CashCategory `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	User     *User         `gorm:"foreignKey:CreatedBy" json:"user,omitempty"`
}
