package gorm

import "api-project/internal/domain/models"

func (r *repository) GetTransactionByID(id uint) (*models.UserBalance, error) {
	var transaction models.UserBalance

	// Find the transaction by its ID
	if err := r.db.First(&transaction, id).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *repository) GetAmountByUserID(userID uint) (uint, error) {

	var transaction models.UserBalance

	// Find the last transaction by the user ID
	if err := r.db.Where("user_id = ?", userID).Last(&transaction).Error; err != nil {
		return 0, err
	}
	return transaction.Amount, nil
}

func (r *repository) GetTransactionLogsByUserID(userID uint) (*models.TransactionLog, error) {

	var transactionLogs models.TransactionLog

	// Find the last transaction log by the user ID for keeping the transaction id in transfer response
	if err := r.db.Where("user_id = ?", userID).Last(&transactionLogs).Error; err != nil {
		return nil, err
	}
	return &transactionLogs, nil
}
