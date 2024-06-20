package gorm

import (
	"api-project/internal/domain/models"
	"api-project/internal/domain/presenters"
	"fmt"
)

func (r *repository) Transfer(data presenters.TransferRequest) (*models.UserBalance, error) {

	tsx := r.db.Begin()

	var transaction models.UserBalance
	var fromFinalAmount, toFinalAmount uint

	fromIntialAmount, err := r.GetAmountByUserID(data.From)
	if err != nil {
		return nil, err
	}

	// Check if the user has enough balance to transfer
	if fromIntialAmount < data.Amount {
		return nil, fmt.Errorf("insufficient balance; available balance: %d", fromIntialAmount)
	}

	// Calculate the remaining balance after the transfer in the From user
	fromFinalAmount = fromIntialAmount - data.Amount

	toIntialAmount, err := r.GetAmountByUserID(data.To)
	if err != nil {
		return nil, err
	}

	// Calculate the remaining balance after the transfer in the To user
	toFinalAmount = toIntialAmount + data.Amount

	// Update the balance of the From user
	if err := r.db.Model(&models.UserBalance{}).Where("user_id = ?", data.From).Update("amount", fromFinalAmount).Error; err != nil {
		tsx.Rollback()
		return nil, err
	}

	// Update the balance of the To user
	if err := r.db.Model(&models.UserBalance{}).Where("user_id = ?", data.To).Update("amount", toFinalAmount).Error; err != nil {
		tsx.Rollback()
		return nil, err
	}

	// Create a transaction log for the transfer in the From user
	if err := r.db.Create(&models.TransactionLog{
		UserID:            data.From,
		TransactionType:   "Transfer-Debit(-)",
		TransactionBy:     "Self",
		TransactionAmount: data.Amount,
		RemainingAmount:   fromFinalAmount,
	}).Error; err != nil {
		tsx.Rollback()
		return nil, err
	}

	// Create a transaction log for the transfer in the To user
	if err := r.db.Create(&models.TransactionLog{
		UserID:            data.To,
		TransactionType:   "Transfer-Credit(+)",
		TransactionBy:     fmt.Sprintf("%d", data.From),
		TransactionAmount: data.Amount,
		RemainingAmount:   toFinalAmount,
	}).Error; err != nil {
		tsx.Rollback()
		return nil, err
	}

	tsx.Commit()

	return &transaction, nil
}
