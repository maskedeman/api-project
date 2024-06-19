package gorm

import "api-project/internal/domain/models"

func (r *repository) GetTransactionByID(id uint) (*models.UserBalance, error) {
	var transaction models.UserBalance
	if err := r.db.First(&transaction, id).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *repository) GetTransactionsByUserID(userID uint) ([]models.UserBalance, error) {
	var transactions []models.UserBalance
	if err := r.db.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *repository) GetTransactionsByUserIDAndTransactionType(userID uint, transactionType string) ([]models.UserBalance, error) {
	var transactions []models.UserBalance
	if err := r.db.Where("user_id = ? AND transaction = ?", userID, transactionType).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *repository) GetAmountByUserID(userID uint) (uint, error) {
	var transaction models.UserBalance
	if err := r.db.Where("user_id = ?", userID).Last(&transaction).Error; err != nil {
		return 0, err
	}
	return transaction.Amount, nil
}

func (r *repository) GetTransactionLogsByUserID(userID uint) (*models.TransactionLog, error) {
	var transactionLogs models.TransactionLog
	if err := r.db.Where("user_id = ?", userID).Last(&transactionLogs).Error; err != nil {
		return nil, err
	}
	return &transactionLogs, nil
}
