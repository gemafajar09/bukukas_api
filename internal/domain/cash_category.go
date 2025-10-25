package domain

import (
	"time"
)

type CashCategory struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:100;not null" json:"name"`
	Type        string    `gorm:"size:10;not null;check:type IN ('in','out','both')" json:"type"`
	Description string    `gorm:"type:text" json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Transactions []CashTransaction `gorm:"foreignKey:CategoryID" json:"transactions,omitempty"`
}
