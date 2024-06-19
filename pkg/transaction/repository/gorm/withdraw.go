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

	if initialAmount < data.Amount {
		return nil, errors.New("insufficient balance")
	}

	finalAmount = initialAmount - data.Amount

	if err = r.db.Model(&models.UserBalance{}).Where("user_id = ?", data.UserID).UpdateColumn("amount", finalAmount).Error; err != nil {
		tsx.Rollback()
		return nil, err
	}

	if err = r.db.Create(&models.TransactionLog{
		UserID:            data.UserID,
		TransactionType:   "Withdraw-",
		TransactionAmount: data.Amount,
		RemainingAmount:   finalAmount,
	}).Error; err != nil {
		tsx.Rollback()
		return nil, err
	}

	tsx.Commit()
	return &finalAmount, err
}
