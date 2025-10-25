package usecase

import (
	"fmt"
	"go-project/internal/domain"
	"go-project/internal/repository"
	"time"
)

type CashUsecase interface {
	RecordTransaction(transaction domain.CashTransaction) error
	GetReport(start, end time.Time) ([]domain.CashTransaction, error)
	CalculateDailyBalance(date time.Time) (*domain.CashBalance, error)
	GetCategories() ([]domain.CashCategory, error)
}

type cashUsecase struct {
	repo repository.CashRepository
}

func NewCashUsecase(repo repository.CashRepository) CashUsecase {
	return &cashUsecase{repo: repo}
}

func (u *cashUsecase) RecordTransaction(transaction domain.CashTransaction) error {
	if transaction.Type != "in" && transaction.Type != "out" {
		return ErrInvalidTransactionType
	}

	transaction.TransactionDate = time.Now()
	if err := u.repo.CreateTransaction(&transaction); err != nil {
		return err
	}

	date := transaction.TransactionDate.Truncate(24 * time.Hour)
	balance, _ := u.repo.GetBalanceByDate(date)
	if balance == nil {
		balance = &domain.CashBalance{
			Date: date,
		}
	}

	if transaction.Type == "in" {
		balance.TotalIn += transaction.Amount
	} else {
		balance.TotalOut += transaction.Amount
	}

	balance.ClosingBalance = balance.OpeningBalance + balance.TotalIn - balance.TotalOut
	balance.CalculatedAt = time.Now()

	return u.repo.SaveOrUpdateBalance(balance)
}

func (u *cashUsecase) GetReport(start, end time.Time) ([]domain.CashTransaction, error) {
	return u.repo.GetTransactions(start, end)
}

func (u *cashUsecase) CalculateDailyBalance(date time.Time) (*domain.CashBalance, error) {
	balance, _ := u.repo.GetBalanceByDate(date)
	if balance == nil {
		balance = &domain.CashBalance{Date: date}
	}
	balance.ClosingBalance = balance.OpeningBalance + balance.TotalIn - balance.TotalOut
	balance.CalculatedAt = time.Now()
	_ = u.repo.SaveOrUpdateBalance(balance)
	return balance, nil
}

func (u *cashUsecase) GetCategories() ([]domain.CashCategory, error) {
	return u.repo.GetAllCategories()
}

var (
	ErrInvalidTransactionType = fmt.Errorf("jenis transaksi tidak valid: harus 'masuk' atau 'keluar'")
)
