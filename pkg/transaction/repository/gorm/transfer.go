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

	if fromIntialAmount < data.Amount {
		return nil, fmt.Errorf("insufficient balance; available balance: %d", fromIntialAmount)
	}

	fromFinalAmount = fromIntialAmount - data.Amount

	toIntialAmount, err := r.GetAmountByUserID(data.To)
	if err != nil {
		return nil, err
	}

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

	if err := r.db.Create(&models.TransactionLog{
		// TransactionID:     data.ID,
		UserID:            data.From,
		TransactionType:   "Transfer-Debit(-)",
		TransactionAmount: data.Amount,
		RemainingAmount:   fromFinalAmount,
	}).Error; err != nil {
		tsx.Rollback()
		return nil, err
	}

	if err := r.db.Create(&models.TransactionLog{
		// TransactionID:     data.ID,
		UserID:            data.To,
		TransactionType:   "Transfer-Credit(+)",
		TransactionAmount: data.Amount,
		RemainingAmount:   toFinalAmount,
	}).Error; err != nil {
		tsx.Rollback()
		return nil, err
	}

	tsx.Commit()

	return &transaction, nil
}
