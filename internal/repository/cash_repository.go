package repository

import (
	"go-project/internal/domain"
	"time"

	"gorm.io/gorm"
)

type CashRepository interface {
	CreateTransaction(transaction *domain.CashTransaction) error
	GetTransactions(start, end time.Time) ([]domain.CashTransaction, error)
	GetBalanceByDate(date time.Time) (*domain.CashBalance, error)
	SaveOrUpdateBalance(balance *domain.CashBalance) error
	GetAllCategories() ([]domain.CashCategory, error)
}

type cashRepository struct {
	db *gorm.DB
}

func NewCashRepository(db *gorm.DB) CashRepository {
	return &cashRepository{db: db}
}

func (r *cashRepository) CreateTransaction(transaction *domain.CashTransaction) error {
	return r.db.Create(transaction).Error
}

func (r *cashRepository) GetTransactions(start, end time.Time) ([]domain.CashTransaction, error) {
	var transactions []domain.CashTransaction
	err := r.db.Preload("Category").
		Where("transaction_date BETWEEN ? AND ?", start, end).
		Order("transaction_date asc").
		Find(&transactions).Error
	return transactions, err
}

func (r *cashRepository) GetBalanceByDate(date time.Time) (*domain.CashBalance, error) {
	var balance domain.CashBalance
	err := r.db.Where("date = ?", date.Format("2006-01-02")).First(&balance).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &balance, err
}

func (r *cashRepository) SaveOrUpdateBalance(balance *domain.CashBalance) error {
	var existing domain.CashBalance
	err := r.db.Where("date = ?", balance.Date.Format("2006-01-02")).First(&existing).Error
	if err == gorm.ErrRecordNotFound {
		return r.db.Create(balance).Error
	}
	balance.ID = existing.ID
	return r.db.Save(balance).Error
}

func (r *cashRepository) GetAllCategories() ([]domain.CashCategory, error) {
	var cats []domain.CashCategory
	err := r.db.Order("name asc").Find(&cats).Error
	return cats, err
}
