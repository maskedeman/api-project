package gorm

import (
	"api-project/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (r *repository) Deposit(data models.UserBalance) (*uint, error) {

	tsx := r.db.Begin()

	var finalAmount uint

	initialAmount, err := r.GetAmountByUserID(data.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			initialAmount = 0
		} else {
			return nil, err
		}
	}
	fmt.Printf("initialAmount: %v\n", initialAmount)
	finalAmount = data.Amount
	finalAmount += initialAmount
	fmt.Printf("finalAmount: %v\n", finalAmount)
	if initialAmount == 0 {
		fmt.Printf("----------------")
		if err := r.db.Create(&models.UserBalance{
			UserID: data.UserID,
			Amount: finalAmount,
		}).Error; err != nil {
			tsx.Rollback()
			return nil, err
		}
	} else {
		fmt.Printf("++++++++++++++++")
		if err := r.db.Model(&models.UserBalance{}).Where("user_id = ?", data.UserID).UpdateColumn("amount", finalAmount).Error; err != nil {
			tsx.Rollback()
			return nil, err
		}
	}

	if err := r.db.Create(&models.TransactionLog{
		// TransactionID:     data.Timestamp.ID,
		UserID:            data.UserID,
		TransactionType:   "Deposit+",
		TransactionAmount: data.Amount,
		RemainingAmount:   finalAmount,
	}).Error; err != nil {
		tsx.Rollback()
		return nil, err
	}

	tsx.Commit()

	return &finalAmount, nil
}
