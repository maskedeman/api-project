package gorm

import (
	"api-project/internal/domain/models"
	"errors"
)

func (r *repository) Withdraw(data models.UserBalance) (*uint, error) {

	tsx := r.db.Begin()

	var finalAmount uint
	var err error

	initialAmount, err := r.GetAmountByUserID(data.UserID)
	if err != nil {
		return nil, err
	}

	// Check if the user has enough balance to withdraw
	if initialAmount < data.Amount {
		return nil, errors.New("insufficient balance")
	}

	// Calculate the remaining balance after the withdrawal
	finalAmount = initialAmount - data.Amount

	// Update the balance of the user
	if err = r.db.Model(&models.UserBalance{}).Where("user_id = ?", data.UserID).UpdateColumn("amount", finalAmount).Error; err != nil {
		tsx.Rollback()
		return nil, err
	}

	// Create a transaction log for the withdrawal
	if err = r.db.Create(&models.TransactionLog{
		UserID:            data.UserID,
		TransactionType:   "Withdraw-",
		TransactionBy:     "Self",
		TransactionAmount: data.Amount,
		RemainingAmount:   finalAmount,
	}).Error; err != nil {
		tsx.Rollback()
		return nil, err
	}

	tsx.Commit()

	return &finalAmount, err
}
