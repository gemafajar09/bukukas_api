package mysql

import (
	"go-project/internal/domain"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.User{},
		&domain.CashCategory{},
		&domain.CashTransaction{},
		&domain.CashBalance{},
	)
}
