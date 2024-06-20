package gorm

import (
	"api-project/internal/domain/models"
	"errors"

	"gorm.io/gorm"
)

func (r *repository) Deposit(data models.UserBalance) (*uint, error) {

	tsx := r.db.Begin()

	var finalAmount uint

	initialAmount, err := r.GetAmountByUserID(data.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			// If the user does not have any balance, set the initial amount to 0
			initialAmount = 0
		} else {
			return nil, err
		}
	}

	// Add the deposit amount to the user's current balance
	finalAmount = data.Amount
	finalAmount += initialAmount

	if initialAmount == 0 {

		// If the user does not have any balance, create a new record
		if err := r.db.Create(&models.UserBalance{
			UserID: data.UserID,
			Amount: finalAmount,
		}).Error; err != nil {
			tsx.Rollback()
			return nil, err
		}
	} else {

		// Update the user's balance with the new amount
		if err := r.db.Model(&models.UserBalance{}).Where("user_id = ?", data.UserID).UpdateColumn("amount", finalAmount).Error; err != nil {
			tsx.Rollback()
			return nil, err
		}
	}

	// Create a transaction log for the deposit
	if err := r.db.Create(&models.TransactionLog{
		UserID:            data.UserID,
		TransactionType:   "Deposit+",
		TransactionBy:     "Self",
		TransactionAmount: data.Amount,
		RemainingAmount:   finalAmount,
	}).Error; err != nil {
		tsx.Rollback()
		return nil, err
	}

	tsx.Commit()

	return &finalAmount, nil
}
