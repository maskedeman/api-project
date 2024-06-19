package interfaces

import (
	"api-project/internal/domain/models"
	"api-project/internal/domain/presenters"
)

type TransactionUsecase interface {
	Deposit(transaction models.UserBalance) (*uint, map[string]string)
	Withdraw(transaction models.UserBalance) (*uint, map[string]string)
	GetBalance(userID uint) (*presenters.BalanceResponse, map[string]string)
	Transfer(transaction presenters.TransferRequest) (*presenters.TransferResponse, map[string]string)
}

type TransactionRepository interface {
	Deposit(transaction models.UserBalance) (*uint, error)
	Withdraw(transaction models.UserBalance) (*uint, error)
	// GetBalance(userID uint) (*models.UserBalance, error)
	Transfer(transaction presenters.TransferRequest) (*models.UserBalance, error)

	GetTransactionByID(id uint) (*models.UserBalance, error)
	GetTransactionsByUserID(userID uint) ([]models.UserBalance, error)
	GetTransactionsByUserIDAndTransactionType(userID uint, transactionType string) ([]models.UserBalance, error)
	GetAmountByUserID(userID uint) (uint, error)
	GetTransactionLogsByUserID(userID uint) (*models.TransactionLog, error)
}
