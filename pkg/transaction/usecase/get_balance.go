package usecase

import (
	"api-project/internal/domain/presenters"
)

func (uCase *usecase) GetBalance(userID uint) (*presenters.BalanceResponse, map[string]string) {
	var err error
	errMap := make(map[string]string)

	_, err = uCase.authRepo.GetUserByID(userID)
	if err != nil {
		errMap["userID"] = err.Error()
		return nil, errMap
	}

	// Delegate the update of data
	amount, err := uCase.repo.GetAmountByUserID(userID)
	if err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}

	return &presenters.BalanceResponse{
		UserID:         userID,
		CurrentBalance: amount,
	}, errMap
}
